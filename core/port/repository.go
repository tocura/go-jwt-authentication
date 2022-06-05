package port

import (
	"context"

	"github.com/tocura/go-jwt-authentication/core/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}
