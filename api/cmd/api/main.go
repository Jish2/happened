package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/caarlos0/env/v11"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"happenedapi/pkg/server"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for database/sql
)

type Config struct {
	DbHost string `env:"DB_HOST"`
	DbUser string `env:"DB_USER"`
	DbPass string `env:"DB_PASS"`
	DbName string `env:"DB_NAME"`
	DbPort int    `env:"DB_PORT"`
}

const (
	Port = 8080
)

type Options struct {
	Debug bool   `doc:"Enable debug logging"`
	Host  string `doc:"Hostname to listen on."`
	Port  int    `doc:"Port to listen on." short:"p" default:"8888"`
}

func pgConnString(config Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPass,
		config.DbName)
}

var api huma.API

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		// Create empty server for generating openapi.yaml
		ctx := context.Background()
		api = server.New(nil)
		var srv http.Server

		hooks.OnStart(func() {
			// Generate openapi.yaml after the server starts

			// Load .env
			err := godotenv.Load(".env")
			if err != nil {
				return
			}

			// Parse env into config
			var config Config
			err = env.Parse(&config)
			if err != nil {
				return
			}

			logger := slog.Default()
			logger.Info("config: ", slog.Any("config", config))
			connString := pgConnString(config)

			// Setup Dependencies
			// Postgres
			db, err := sql.Open("pgx", connString)
			if err != nil {
				panic(err)
			}
			if err := db.Ping(); err != nil {
				panic(err)
			}

			cfg, err := awsConfig.LoadDefaultConfig(ctx)
			if err != nil {
				log.Fatalln(err)
			}

			// Setup S3 bucket
			s3Client := s3.NewFromConfig(cfg)
			_ = s3Client

			// Create server
			api = server.New(db)
			srv = http.Server{
				Addr:    fmt.Sprintf(":%d", Port),
				Handler: api.Adapter(),
			}
			logger.Info("server listening", slog.Int("port", Port))
			if err = srv.ListenAndServe(); err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					slog.Info("shutting down server")
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
