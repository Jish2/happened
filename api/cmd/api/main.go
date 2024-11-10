package main

import (
	"connectrpc.com/connect"
	"context"
	"database/sql"
	"errors"
	"fmt"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"happenedapi/gen/protos/v1/happenedv1connect"
	"happenedapi/pkg/auth"
	"happenedapi/pkg/server"
	"log"
	"log/slog"
	"net/http"
	"os"

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
	logger.Info("config: ", config)
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
	// Create server
	api := server.New(s3Client)
	mux := http.NewServeMux()

	// Create server, register interceptors
	path, handler := happenedv1connect.NewHappenedServiceHandler(
		api,
		connect.WithInterceptors(
			auth.NewAuthInterceptor(),
		),
	)
	mux.Handle(path, handler)

	logger.Info("server listening", slog.Int("port", Port))
	err = http.ListenAndServe(
		fmt.Sprintf("localhost:%d", Port),
		h2c.NewHandler(mux, &http2.Server{}),
	)

	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			slog.Info("shutting down server")
			os.Exit(0)
		} else {
			slog.Error("unexpected error", slog.Any("error", err))
			os.Exit(1)
		}
	}
}
