package service

import (
	"context"
	"net/http"

	"github.com/tocura/go-jwt-authentication/core/model"
	"github.com/tocura/go-jwt-authentication/core/port"
	"github.com/tocura/go-jwt-authentication/pkg/log"
	"github.com/tocura/go-jwt-authentication/pkg/web"
)

var (
	errorInvalidCredentials = web.NewError(http.StatusNotFound, "Email and/or password are invalid")
	errorEmailAlreadyInUse  = web.NewError(http.StatusConflict, "Email already in use")
	errorCreateUser         = web.NewError(http.StatusInternalServerError, "Error to create user")
	errorGenerateJWT        = web.NewError(http.StatusInternalServerError, "Error to generate token")
)

type userService struct {
	repo  port.UserRepository
	token port.TokenService
}

func NewUserService(repo port.UserRepository, token port.TokenService) port.UserService {
	return &userService{
		repo:  repo,
		token: token,
	}
}

func (us *userService) Create(ctx context.Context, user model.User) (*model.User, error) {
	usr, _ := us.repo.GetByEmail(ctx, user.Email)
	if usr != nil {
		log.Warn(ctx, "email already exists in database")
		return nil, errorEmailAlreadyInUse
	}

	encryptedPassword, err := us.token.GenerateHashPassword(ctx, user.Password)
	if err != nil {
		log.Error(ctx, "error to encrypt password", err)
		return nil, errorCreateUser
	}

	user.SetEncryptedPassword(string(encryptedPassword))

	newUser, err := us.repo.Create(ctx, user)
	if err != nil {
		return nil, errorCreateUser
	}

	return newUser, nil
}

func (us *userService) Login(ctx context.Context, login model.Login) (string, error) {
	user, err := us.repo.GetByEmail(ctx, login.Email)
	if err != nil {
		return "", errorInvalidCredentials
	}

	if !us.token.IsValidPassword(ctx, user.Password, login.Password) {
		log.Warn(ctx, "invalid login credentials")
		return "", errorInvalidCredentials
	}

	token, err := us.token.GenerateJWT(ctx, user.Email, string(user.Role))
	if err != nil {
		return "", errorGenerateJWT
	}

	return token, nil
}
