package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

type contextKey string

const (
	JwtKey          contextKey = "jwt"
	JwtExpiresInKey contextKey = "jwtExpiresIn"
)

func JwtMiddleware(jwt *jwtauth.JWTAuth, jwtExpiresIn int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), JwtKey, jwt)
			ctx = context.WithValue(ctx, JwtExpiresInKey, jwtExpiresIn)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
