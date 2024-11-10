package main

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	happenedv1 "happenedapi/gen/protos/v1"
	"happenedapi/gen/protos/v1/happenedv1connect"
	"log"
	"log/slog"
	"net/http"
	"os"
)

// Ensure interface satisfaction
var _ happenedv1connect.HappenedServiceHandler = (*HappenedServer)(nil)

type HappenedServer struct{}

func (s *HappenedServer) CreateEvent(ctx context.Context, c *connect.Request[happenedv1.CreateEventRequest]) (*connect.Response[happenedv1.CreateEventResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *HappenedServer) Greet(
	ctx context.Context,
	req *connect.Request[happenedv1.GreetRequest]) (*connect.Response[happenedv1.GreetResponse], error) {
	log.Println("Request headers", req.Header())

	res := connect.NewResponse(&happenedv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	greeter := &HappenedServer{}
	mux := http.NewServeMux()

	path, handler := happenedv1connect.NewHappenedServiceHandler(greeter)
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
