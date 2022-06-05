package request

import (
	"github.com/tocura/go-jwt-authentication/core/enum"
	"github.com/tocura/go-jwt-authentication/core/model"
)

type SignUp struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Role     string `json:"role" validate:"required,oneof=normal premium"`
}

func (s *SignUp) MapToUserDomain() *model.User {
	return &model.User{
		Email:    s.Email,
		Password: s.Password,
		Role:     enum.MapToEnumRole(s.Role),
	}
}
