package model

import "github.com/tocura/go-jwt-authentication/core/enum"

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Role     enum.EnumRole
}

func (u *User) SetEncryptedPassword(encryptedPassword string) {
	u.Password = encryptedPassword
}
