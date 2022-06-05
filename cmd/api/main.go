package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	controller "github.com/tocura/go-jwt-authentication/adapter/http"
	pmiddleware "github.com/tocura/go-jwt-authentication/adapter/http/middleware"
	"github.com/tocura/go-jwt-authentication/adapter/http/server"
	"github.com/tocura/go-jwt-authentication/adapter/repository/mysql"
	"github.com/tocura/go-jwt-authentication/core/service"
	"github.com/tocura/go-jwt-authentication/pkg/env"
	"github.com/tocura/go-jwt-authentication/pkg/log"
)

func main() {
	cfg, err := env.New()
	if err != nil {
		log.Fatal(context.TODO(), "fail to load application configs", err)
	}

	router := chi.NewRouter()
	router.Use(pmiddleware.Logger(log.ZapLogger, true))
	router.Use(middleware.Heartbeat("/health"))
	router.Use(middleware.Recoverer)
	router.Use(pmiddleware.RequestID)
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	db, err := mysql.NewConnection(cfg.Database)
	if err != nil {
		log.Fatal(context.TODO(), "fail to connect to mysql database", err)
	}

	// Repositories.
	userRepo := mysql.NewUserRepository(db)

	// Services.
	tokenService := service.NewTokenService()
	userService := service.NewUserService(userRepo, tokenService)

	// Handler.
	controller.NewHandler(router, userService)

	server.New().
		Address(cfg.App.HTTPAddress).
		Routes(router).
		Run()
}
