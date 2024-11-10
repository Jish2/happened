package main

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	greetv1 "happenedapi/gen/greet/v1"
	"happenedapi/gen/greet/v1/greetv1connect"
	"log"
	"log/slog"
	"net/http"
	"os"
)

// Ensure interface satisfaction
var _ greetv1connect.GreetServiceHandler = (*GreetServer)(nil)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers", req.Header())

	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()

	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
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
