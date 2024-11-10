package main

import (
	"errors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"happenedapi/gen/protos/v1/happenedv1connect"
	"happenedapi/server"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	api := server.New()
	mux := http.NewServeMux()

	path, handler := happenedv1connect.NewHappenedServiceHandler(api)
	mux.Handle(path, handler)

	err := http.ListenAndServe(
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
