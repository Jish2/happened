package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"happenedapi/pkg/server"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"time"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/caarlos0/env/v11"
	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for database/sql
	"github.com/joho/godotenv"
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

func pgConnString(config Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPass,
		config.DbName)
}

func main() {
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
	api := server.New(db)
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", Port),
		Handler: api,
	}

	logger.Info("server listening", slog.Int("port", Port))
	// Generate openapi.yaml after the server starts
	go func() {
		// Wait for server start
		time.Sleep(time.Millisecond * 250)

		for {
			response, err := http.Get(fmt.Sprintf("http://localhost:%d/ping", Port))
			if response.StatusCode == http.StatusOK && err == nil {
				break
			}
		}

		logger.Info("Generating openapi.yaml...")
		cmd := exec.Command("sh", "-c", fmt.Sprintf("curl http://localhost:%d/openapi.yaml > openapi.yaml", Port))
		// Generate the OpenAPI spec
		if err := cmd.Run(); err != nil {
			logger.Error("✖ Error generating openapi.yaml", slog.Any("error", err))
			return
		}
		logger.Info("✔ Successfully generated openapi.yaml ")

		logger.Info("Generating Typescript client SDK...")
		// Generate the client SDK with Orval
		cmd = exec.Command("sh", "-c", "make -C ../ gen")

		if err := cmd.Run(); err != nil {
			logger.Error("✖ Error generating Typescript client SDK", slog.Any("error", err))
			return
		}
		logger.Info("✔ Successfully generated client SDK")

	}()
	if err = srv.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			slog.Info("shutting down server")
			os.Exit(0)
		} else {
			slog.Error("unexpected error", slog.Any("error", err))
			os.Exit(1)
		}
	}
}
