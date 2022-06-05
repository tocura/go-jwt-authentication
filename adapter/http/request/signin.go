package request

import "github.com/tocura/go-jwt-authentication/core/model"

type SignIn struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s *SignIn) MapToLoginDomain() *model.Login {
	return &model.Login{
		Email:    s.Email,
		Password: s.Password,
	}
}
