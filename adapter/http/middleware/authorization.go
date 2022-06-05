package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/tocura/go-jwt-authentication/pkg/web"
)

var (
	errorUnauthorized = web.NewError(http.StatusUnauthorized, "Invalid token")
)

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errorUnauthorized
				}

				key := os.Getenv("SECRET")
				return []byte(key), nil
			})

			if err := web.ToPlanetError(err); err != nil {
				b, _ := err.JSON()
				w.WriteHeader(err.Status)
				w.Write(b)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if claims["role"] == "premium" {
					r.Header.Set("X-Role", "premium")
					next.ServeHTTP(w, r)
					return

				} else if claims["role"] == "normal" {
					r.Header.Set("X-Role", "normal")
					next.ServeHTTP(w, r)
					return
				}

			} else {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(errorUnauthorized)
				return
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errorUnauthorized)
			return
		}
	})
}
