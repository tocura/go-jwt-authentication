package main

import (
	"context"
	"log"

	"github.com/tocura/go-jwt-authentication/pkg/env"
)

func main() {
	_, err := env.New()
	if err != nil {
		log.Fatal(context.TODO(), "fail to load application configs")
	}
}
