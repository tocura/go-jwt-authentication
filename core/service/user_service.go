package service

import (
	"context"
	"net/http"

	"github.com/tocura/go-jwt-authentication/core/model"
	"github.com/tocura/go-jwt-authentication/core/port"
	"github.com/tocura/go-jwt-authentication/pkg/log"
	"github.com/tocura/go-jwt-authentication/pkg/web"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) port.UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) Create(ctx context.Context, user model.User) (*model.User, error) {
	usr, _ := us.repo.GetByEmail(ctx, user.Email)
	if usr != nil {
		log.Warn(ctx, "email already exists in database")
		return nil, web.NewError(http.StatusConflict, "Email already in use")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Error(ctx, "error to encrypt password", err)
		return nil, web.NewError(http.StatusInternalServerError, "Error to create user")
	}

	user.SetEncryptedPassword(string(encryptedPassword))

	newUser, err := us.repo.Create(ctx, user)
	if err != nil {
		return nil, web.NewError(http.StatusInternalServerError, "Error to create user")
	}

	return newUser, nil
}

func (us *userService) Login(ctx context.Context, login model.Login) (*model.User, error) {
	user, err := us.repo.Login(ctx, login)
	if err != nil {
		return nil, web.NewError(http.StatusNotFound, "Email and/or password are invalid")
	}

	return user, nil
}
