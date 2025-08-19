package middleware

import (
	"context"
	"dnd-game/internal/util/helper"
	"dnd-game/internal/util/services"
	"net/http"
	"strings"
	"time"
)

func getTokenFromHeader(r *http.Request) string {
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(authHeader) == 2 {
		return authHeader[1]
	}
	return ""
}

func NewAuthMiddleware(service *services.AuthService, helper *helper.ServerHelper, skipper func(r *http.Request) bool) Middleware {
	return func(next http.Handler) http.Handler {
		handler := func(w http.ResponseWriter, r *http.Request) error {
			if skipper(r) { // skips auth based on some condition
				next.ServeHTTP(w, r)
				return nil
			}

			jwtToken := getTokenFromHeader(r)
			if jwtToken == "" {
				return helper.Unauthorized("Malformed token")

			} else {
				claims, err := service.VerifyToken(jwtToken)
				if err != nil {
					return helper.Unauthorized(err.Error())

				} else if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
					return helper.Unauthorized("Token is expired")

				} else {
					ctx := context.WithValue(r.Context(), service.AuthContextKey, claims)
					next.ServeHTTP(w, r.WithContext(ctx))
				}
			}
			return nil
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := handler(w, r)
			helper.HandleHTTPError(err, w)
		})
	}
}
