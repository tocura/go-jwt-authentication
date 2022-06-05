package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tocura/go-jwt-authentication/adapter/http/middleware"
	"github.com/tocura/go-jwt-authentication/pkg/log"
	"github.com/tocura/go-jwt-authentication/pkg/web"
)

type protectedHandler struct{}

func newProtectedHandler(router chi.Router) {
	handler := &protectedHandler{}

	router.Route("/protected", func(r chi.Router) {
		r.Use(middleware.IsAuthorized)
		r.Get("/", handler.GetProtectedContent)
	})
}

func (h *protectedHandler) GetProtectedContent(w http.ResponseWriter, r *http.Request) {
	roleHeader := r.Header.Get("X-Role")
	log.Info(r.Context(), roleHeader)
	if roleHeader == "" || roleHeader == "normal" {
		log.Warn(r.Context(), "unauthorized role")
		responseError(w, web.NewError(http.StatusUnauthorized, "You don't have permission to access this content"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Congrats! You're a premium user! :)")
}
