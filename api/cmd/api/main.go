package main

import (
	"bytes"
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
	"os/exec"
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

func createOpenAPIAndClientSDK() {
	// Wait for server start
	time.Sleep(time.Millisecond * 100)

	const MaxAttempts = 5
	attempts := 0
	for attempts < MaxAttempts {
		response, err := http.Get(fmt.Sprintf("http://localhost:%d/ping", Port))
		if err == nil && response.StatusCode == http.StatusOK {
			break
		}
		attempts++
		time.Sleep(time.Millisecond * 50)
	}

	slog.Info("Generating openapi.yaml...")
	cmd := exec.Command("sh", "-c", fmt.Sprintf("curl http://localhost:%d/openapi.yaml > openapi.yaml", Port))
	// Generate the OpenAPI spec
	if err := cmd.Run(); err != nil {
		slog.Error("✖ Error generating openapi.yaml", slog.Any("error", err))
		return
	}
	slog.Info("✔ Successfully generated openapi.yaml")

	// Clean before acknowledging to confirm generation from Orval.
	cmd = exec.Command("sh", "-c", "make -C ../ clean")
	if err := cmd.Run(); err != nil {
		slog.Error("error cleaning gen")
		return
	}

	slog.Info("Generating Typescript client SDK...")
	// Generate the client SDK with Orval
	cmd = exec.Command("sh", "-c", "make -C ../ gen")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	fmt.Println(cmd.Stdout)
	if err != nil {
		slog.Error("✖ Error generating Typescript client SDK", slog.Any("error", err))
		return
	}

	for _, err := os.Stat("../client/gen"); os.IsNotExist(err); {
		log.Print("hello")
	}
	slog.Info("✔ Successfully generated client SDK")

}

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		// Generate openapi.yaml after the server starts
		// go createOpenAPIAndClientSDK()

		slog.Info("options", slog.Any("opts", opts))

		hooks.OnStart(func() {
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

			ctx := context.Background()
			cfg, err := awsConfig.LoadDefaultConfig(ctx)
			if err != nil {
				log.Fatalln(err)
			}

			// Setup S3 bucket
			s3Client := s3.NewFromConfig(cfg)
			_ = s3Client

			// Create server
			api = server.New(db)
			srv := http.Server{
				Addr:    fmt.Sprintf(":%d", Port),
				Handler: api.Adapter(),
			}

			logger.Info("server listening", slog.Int("port", Port))

			slog.Info("hello")
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
	})

	cli.Root().AddCommand(&cobra.Command{
		Use:   "openapi",
		Short: "Print the OpenAPI spec",
		Run: func(cmd *cobra.Command, args []string) {
			b, err := api.OpenAPI().YAML()
			if err != nil {
				panic(err)
			}

			fmt.Println(string(b))
		},
	})
	cli.Run()
}
