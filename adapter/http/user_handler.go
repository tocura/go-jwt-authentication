package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tocura/go-jwt-authentication/adapter/http/request"
	"github.com/tocura/go-jwt-authentication/adapter/http/response"
	"github.com/tocura/go-jwt-authentication/core/port"
	"github.com/tocura/go-jwt-authentication/pkg/log"
	"github.com/tocura/go-jwt-authentication/pkg/web"
)

type userHandler struct {
	userService port.UserService
}

func newUserHandler(router chi.Router, userService port.UserService) {
	handler := &userHandler{
		userService: userService,
	}

	router.Post("/signup", handler.SignUp)
	router.Post("/signin", handler.SignIn)
}

func (h *userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var signup request.SignUp
	if err := json.NewDecoder(r.Body).Decode(&signup); err != nil {
		log.Error(r.Context(), "invalid request body", err)
		responseError(w, web.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	signupDomain := signup.MapToUserDomain()
	user, err := h.userService.Create(r.Context(), *signupDomain)
	if err != nil {
		responseError(w, err)
		return
	}

	res := response.MapToSignUpResponse(*user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *userHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var signin request.SignIn
	if err := json.NewDecoder(r.Body).Decode(&signin); err != nil {
		log.Error(r.Context(), "invalid request body", err)
		responseError(w, web.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	signinDomain := signin.MapToLoginDomain()
	token, err := h.userService.Login(r.Context(), *signinDomain)
	if err != nil {
		responseError(w, err)
		return
	}

	res := response.MapToTokenResponse(token)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
