package response

import "github.com/tocura/go-jwt-authentication/core/model"

type SignUp struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func MapToSignUpResponse(user model.User) *SignUp {
	return &SignUp{
		ID:    user.ID,
		Email: user.Email,
		Role:  string(user.Role),
	}
}
