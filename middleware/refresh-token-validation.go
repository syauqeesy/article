package middleware

import (
	"context"
	"net/http"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/configuration"
	"ahmadsyauqi.dev/article/exception"
	"github.com/golang-jwt/jwt/v5"
)

type refreshTokenContextKey struct{}

func GetRefreshToken(ctx context.Context) (string, bool) {
	value, ok := ctx.Value(refreshTokenContextKey{}).(string)

	return value, ok
}

func RefreshTokenValidation(configuration *configuration.Configuration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims := jwt.MapClaims{}
			refreshTokenCookie, err := r.Cookie("refresh_token")
			if err != nil || refreshTokenCookie.Value == "" {
				common.HttpErrorHandler(w, exception.Unauthorized, nil)
				return
			}

			token, err := jwt.ParseWithClaims(refreshTokenCookie.Value, claims, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, exception.Unauthorized
				}

				return []byte(configuration.Authentication.RefreshSecret), nil
			})
			if err != nil || !token.Valid {
				common.HttpErrorHandler(w, exception.Unauthorized, nil)
				return
			}

			ctx := context.WithValue(r.Context(), refreshTokenContextKey{}, refreshTokenCookie.Value)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
