package server

import (
	"database/sql"
	"fmt"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"happenedapi/pkg/images"

	"net/http"
	"strings"
)

type HumaMiddleware func(ctx huma.Context, next func(huma.Context))

func ClerkAuthMiddleware(api huma.API) HumaMiddleware {
	return func(ctx huma.Context, next func(huma.Context)) {
		sessionToken := strings.TrimPrefix(ctx.Header("Authorization"), "Bearer ")

		claims, err := jwt.Verify(ctx.Context(), &jwt.VerifyParams{
			Token: sessionToken,
		})
		if err != nil {
			huma.WriteErr(api, ctx,
				http.StatusUnauthorized,
				"unauthorized",
				fmt.Errorf("invalid token"),
			)
			return
		}
		ctx = huma.WithValue(ctx, "claims", claims)
		next(ctx)
	}
}

func RegisterAPI(api huma.API, db *sql.DB, imageService *images.Service) {
	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer)

	config := huma.DefaultConfig("My API", "1.0.0")
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"BearerAuth": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}

	registerRoutes(api, db, imageService)

}

func registerRoutes(
	api huma.API,
	db *sql.DB,
	imageService *images.Service,
) {

	_ = db
	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Path:        "/greeting/public/{name}",
		Method:      http.MethodGet,
		Summary:     "Get a greeting",
		Description: "Get a greeting for a person by name.",
		Tags:        []string{"Greetings"},
	}, GreetHandler())

	huma.Register(api, huma.Operation{
		OperationID: "protected-greet",
		Method:      http.MethodGet,
		Path:        "/greeting/protected/{name}",
		Middlewares: huma.Middlewares{ClerkAuthMiddleware(api)},
		Tags:        []string{"Greetings"},
		Summary:     "Get a protected greeting",
		Description: "Protected version of greet",
		Security: []map[string][]string{
			{
				"BearerAuth": {},
			},
		},
	}, protectedGreetHandler())

	huma.Get(api, "/create-upload-url", CreateUploadURLHandler(imageService))

}
