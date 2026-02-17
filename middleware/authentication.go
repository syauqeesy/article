package middleware

import (
	"context"
	"fmt"
	"net/http"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/configuration"
	"ahmadsyauqi.dev/article/exception"
	"github.com/golang-jwt/jwt/v5"
)

type subjectContextKey struct{}

func GetSubject(ctx context.Context) (string, bool) {
	value, ok := ctx.Value(subjectContextKey{}).(string)

	return value, ok
}

func Authentication(configuration *configuration.Configuration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			fmt.Println(r.Method, r.URL.Path)
			accessTokenCookie, err := r.Cookie("access_token")
			if err != nil || accessTokenCookie.Value == "" {
				common.HttpErrorHandler(w, exception.Unauthorized, nil)
				return
			}

			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(accessTokenCookie.Value, claims, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, exception.Unauthorized
				}

				return []byte(configuration.Authentication.Secret), nil
			})
			if err != nil || !token.Valid {
				common.HttpErrorHandler(w, exception.Unauthorized, nil)
				return
			}

			claimType, ok := claims["type"].(string)
			if !ok || claimType != "access" {
				common.HttpErrorHandler(w, exception.Unauthorized, nil)
				return
			}

			subject, ok := claims["sub"].(string)
			if !ok || subject == "" {
				common.HttpErrorHandler(w, exception.Unauthorized, nil)
				return
			}

			ctx := context.WithValue(r.Context(), subjectContextKey{}, subject)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
