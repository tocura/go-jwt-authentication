package service

import (
	"context"
	"os"
	"time"

	"github.com/tocura/go-jwt-authentication/core/port"
	"github.com/tocura/go-jwt-authentication/pkg/log"
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
)

type tokenService struct{}

func NewTokenService() port.TokenService {
	return &tokenService{}
}

func (ts *tokenService) GenerateHashPassword(ctx context.Context, password string) (string, error) {
	log.Info(ctx, "hashing password")

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Error(ctx, "error to hash user password", err)
		return "", err
	}

	return string(hashPassword), nil
}

func (ts *tokenService) IsValidPassword(ctx context.Context, hashPassword, password string) bool {
	log.Info(ctx, "checking if user password is valid")

	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func (ts *tokenService) GenerateJWT(ctx context.Context, email, role string) (string, error) {
	log.Info(ctx, "generating jwt auth token")

	key := os.Getenv("SECRET")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Error(ctx, "error to generate jwt token", err)
		return "", err
	}

	return tokenString, nil
}
