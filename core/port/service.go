package port

import (
	"context"

	"github.com/tocura/go-jwt-authentication/core/model"
)

type UserService interface {
	Create(ctx context.Context, user model.User) (*model.User, error)
	Login(ctx context.Context, login model.Login) (string, error)
}

type TokenService interface {
	GenerateHashPassword(ctx context.Context, password string) (string, error)
	IsValidPassword(ctx context.Context, hashPassword, password string) bool
	GenerateJWT(ctx context.Context, email, role string) (string, error)
}
