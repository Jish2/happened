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
	"github.com/joho/godotenv"

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
		time.Sleep(time.Millisecond * 250)
		cmd := exec.Command("sh", "-c", fmt.Sprintf("curl http://localhost:%d/openapi.yaml > openapi.yaml", Port))

		if err := cmd.Run(); err != nil {
			logger.Error("error generating openapi spec", slog.Any("error", err))
		}
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
