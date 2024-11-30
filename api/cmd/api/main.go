package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"happenedapi/pkg/images"
	"happenedapi/pkg/server"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/caarlos0/env/v11"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for database/sql
)

type Config struct {
	DbHost         string `env:"DB_HOST"`
	DbUser         string `env:"DB_USER"`
	DbPass         string `env:"DB_PASS"`
	DbName         string `env:"DB_NAME"`
	DbPort         string `env:"DB_PORT"`
	ClerkSecretKey string `env:"CLERK_SECRET_KEY"`
	Port           int    `env:"PORT" envDefault:"8080"`
}

type Options struct {
	Debug bool   `doc:"Enable debug logging"`
	Stage string `doc:"environment" short:"s" default:"production"`
}

type Stage = string

const (
	Development Stage = "development"
	Production  Stage = "production"
)

func main() {

	// Create empty server for generating openapi.yaml with the CLI
	api := server.New(nil, nil)
	logger := slog.Default()

	// Start up and stop hooks
	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		var srv http.Server
		hooks.OnStart(func() {
			ctx := context.Background()

			if opts.Stage == Production {
				slog.Info("launching server in production mode")
			} else {
				slog.Info("defaulting to server development mode")
			}
			var err error
			// Parse env into config
			var config Config
			err = env.Parse(&config)
			if err != nil {
				slog.Error("parsing env to config", slog.Any("error", err))
				os.Exit(1)
			}

			logger.Info("config: ", slog.Any("config", config))
			logger.Info("huma options: ", slog.Any("options", opts))

			dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName)

			if opts.Stage == Development {
				// Load .env
				err := godotenv.Load(".env")
				if err != nil {
					slog.Error("loading env", slog.Any("error", err))
					os.Exit(1)
				}
			}

			logger.Info("setting clerk secret key from environment config")
			clerk.SetKey(config.ClerkSecretKey)

			// Setup Dependencies
			// Postgres
			db, err := sql.Open("pgx", dbURL)
			if err != nil {
				slog.Error("opening database", slog.Any("error", err))
				os.Exit(1)
			}
			logger.Info("pinging db")

			dbctx, cancel := context.WithTimeout(ctx, time.Second*5)
			defer cancel()
			if err := db.PingContext(dbctx); err != nil {
				slog.Error("pinging db", slog.Any("error", err))
				os.Exit(1)
			}
			logger.Info("successfully pinged db")

			cfg, err := awsConfig.LoadDefaultConfig(ctx)
			if err != nil {
				slog.Error("loading default aws config", slog.Any("error", err))
				os.Exit(1)
			}

			// Setup S3 bucket
			s3Client := s3.NewFromConfig(cfg)
			s3PresignClient := s3.NewPresignClient(s3Client)
			imageService := images.NewService(s3PresignClient)

			// Create server
			api = server.New(db, imageService)
			srv = http.Server{
				Addr:    fmt.Sprintf(":%d", config.Port),
				Handler: api.Adapter(),
			}

			logger.Info("server listening", slog.Int("port", config.Port))
			if err = srv.ListenAndServe(); err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					logger.Info("shutting down server")
					os.Exit(0)
				} else {
					slog.Error("unexpected error", slog.Any("error", err))
					os.Exit(1)
				}
			}
		})

		hooks.OnStop(func() {
			// Gracefully shutdown your server here
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := srv.Shutdown(ctx)
			if err != nil {
				log.Fatalln(err)
			}
		})
	})

	// Additional command for generating the OpenAPI specification
	cli.Root().AddCommand(&cobra.Command{
		Use:   "openapi",
		Short: "Print the OpenAPI spec",
		Run: func(cmd *cobra.Command, args []string) {
			b, err := api.OpenAPI().YAML()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(string(b))
		},
	})

	cli.Run()
}
