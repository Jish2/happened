package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"happenedapi/gen/protos/v1/happenedv1connect"
	"happenedapi/pkg/server"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"connectrpc.com/grpcreflect"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for database/sql
)

type Config struct {
	DbHost string `env:"DB_HOST"`
	DbUser string `env:"DB_USER"`
	DbPass string `env:"DB_PASS"`
	DbName string `env:"DB_NAME"`
	DbPort int    `env:"DB_PORT"`
}

var (
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
	stage := flag.String("stage", "production", "-stage development|production")
	flag.Parse()
	var err error

	if os.Getenv("PORT") != "" {
		Port, err = strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			slog.Error("failed to parse PORT", slog.Any("error", err))
			os.Exit(1)
		}
	}

	slog.Info("STAGE 2", slog.String("stage", *stage))

	if *stage == "development" {
		// Load .env
		err = godotenv.Load(".env")
		if err != nil {
			slog.Error("failed to load .env", slog.Any("error", err))
		}
	}

	// Parse env into config
	var config Config
	err = env.Parse(&config)
	if err != nil {
		slog.Error("failed to parse env", slog.Any("error", err))
		os.Exit(1)
	}
	logger := slog.Default()
	logger.Info("config: ", slog.Any("config", config))
	connString := pgConnString(config)

	// Setup Dependencies
	// Postgres
	db, err := sql.Open("pgx", connString)
	if err != nil {
		slog.Error("failed to open postgres", slog.Any("error", err))
		os.Exit(1)
	}
	if err := db.Ping(); err != nil {
		slog.Error("failed to ping postgres", slog.Any("error", err))
		os.Exit(1)
	}

	ctx := context.Background()
	cfg, err := awsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		slog.Error("failed to load aws config", slog.Any("error", err))
		os.Exit(1)
	}

	// Setup S3 bucket
	s3Client := s3.NewFromConfig(cfg)
	api := server.New(s3Client)
	mux := http.NewServeMux()

	// Create server
	path, handler := happenedv1connect.NewHappenedServiceHandler(api)
	mux.Handle(path, handler)
	slog.Info("happenedv1connect.HappenedServiceName", slog.String("name", happenedv1connect.HappenedServiceName))

	reflector := grpcreflect.NewStaticReflector(
		happenedv1connect.HappenedServiceName,
	)
	mux.HandleFunc("GET /healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}))
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	logger.Info("server listening",
		slog.Int("port", Port),
		slog.String("path", path),
	)
	err = http.ListenAndServe(
		fmt.Sprintf(":%d", Port),
		h2c.NewHandler(mux, &http2.Server{}),
	)

	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			slog.Info("shutting down server", slog.Any("error", err))
			os.Exit(0)
		} else {
			slog.Error("unexpected error", slog.Any("error", err))
			os.Exit(1)
		}
	}
}
