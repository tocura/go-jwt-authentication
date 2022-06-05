package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/tocura/go-jwt-authentication/pkg/log"
)

const RequestIDHeader = "X-Request-Id"

func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			requestID = uuid.NewString()
		}

		ctx := log.NewLogContext(r.Context(), map[string]interface{}{
			"X-Request-Id": requestID,
		})

		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
