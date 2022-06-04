package main

import (
	"context"
	"log"

	"github.com/tocura/go-jwt-authentication/adapter/repository/mysql"
	"github.com/tocura/go-jwt-authentication/pkg/env"
)

func main() {
	cfg, err := env.New()
	if err != nil {
		log.Fatal(context.TODO(), "fail to load application configs")
	}

	_, err = mysql.NewConnection(cfg.Database)
	if err != nil {
		log.Fatal(context.TODO(), "fail to connect to mysql database")
	}
}
