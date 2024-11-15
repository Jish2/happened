package server

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

type HumaHandler[I, O any] func(ctx context.Context, input *I) (*O, error)

type Transformer[T any] func(T) T

type GreetingInput struct{
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}

func greetHandler(db *sql.DB) HumaHandler[GreetingInput, GreetingOutput] {
	return func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	}
}


func NewHTTPServer(db *sql.DB) chi.Router {
	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

	huma.Get(api, "/greeting/{name}", greetHandler(db))

	return router
}