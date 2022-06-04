package entity

import "github.com/tocura/go-jwt-authentication/core/model"

type Login struct {
	Email    string
	Password string
}

func MapToLoginEntity(login model.Login) *Login {
	return &Login{
		Email:    login.Email,
		Password: login.Password,
	}
}
