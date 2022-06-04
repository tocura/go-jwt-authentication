package port

import (
	"context"

	"github.com/tocura/go-jwt-authentication/core/model"
)

type UserService interface {
	Create(ctx context.Context, user model.User) (*model.User, error)
	Login(ctx context.Context, login model.Login) (*model.User, error)
}
