package main

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"happenedapi/gen/protos/v1/happenedv1connect"
	"happenedapi/pkg/server"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	s3Client := s3.NewFromConfig(cfg)
	api := server.New(s3Client)
	mux := http.NewServeMux()

	path, handler := happenedv1connect.NewHappenedServiceHandler(api)
	mux.Handle(path, handler)

	err = http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)

	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			slog.Info("shutting down server")
			os.Exit(1)
		} else {
			slog.Error("unexpected error", slog.Any("error", err))
		}
	}
}
