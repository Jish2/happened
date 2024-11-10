package auth

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"strings"
)

var clerkClaimsKey = "clerk-claims-key"

func WithTokenClaims(ctx context.Context, claims *clerk.SessionClaims) context.Context {
	return context.WithValue(ctx, clerkClaimsKey, claims)
}

func GetTokenClaims(ctx context.Context) (*clerk.SessionClaims, error) {
	claims, ok := ctx.Value(clerkClaimsKey).(*clerk.SessionClaims)
	if !ok {
		return nil, errors.New("claims not set")
	}

	return claims, nil
}

func NewAuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			sessionToken := strings.TrimPrefix(req.Header().Get("Authorization"), "Bearer ")
			claims, err := jwt.Verify(ctx, &jwt.VerifyParams{
				Token: sessionToken,
			})

			// Token invalid or not provided
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}
			ctx = WithTokenClaims(ctx, claims)
			return next(ctx, req)
		}
	}

	return interceptor
}
