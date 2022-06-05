package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tocura/go-jwt-authentication/core/port"
	"github.com/tocura/go-jwt-authentication/pkg/web"
)

func NewHandler(router *chi.Mux, userService port.UserService) {
	router.Route("/api/v1", func(r chi.Router) {
		newUserHandler(r, userService)
	})
}

func responseError(w http.ResponseWriter, err error) {
	if err := web.ToPlanetError(err); err != nil {
		b, _ := err.JSON()
		w.WriteHeader(err.Status)
		w.Write(b)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode("unknown error")
}
