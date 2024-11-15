package server

import (
	"context"
	"database/sql"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HumaHandler[I, O any] func(ctx context.Context, input *I) (*O, error)

func New(db *sql.DB) chi.Router {
	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer)

	api := humachi.New(r, huma.DefaultConfig("My API", "1.0.0"))

	registerRoutes(api, db)

	return r
}

func registerRoutes(
	api huma.API,
	db *sql.DB,
) {

	_ = db
	huma.Get(api, "/greeting/{name}", greetHandler())
}
