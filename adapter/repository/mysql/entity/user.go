package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/tocura/go-jwt-authentication/core/enum"
	"github.com/tocura/go-jwt-authentication/core/model"
)

type User struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"type:string;not null" json:"name"`
	Email     string    `gorm:"type:string;not null" json:"email"`
	Password  string    `gorm:"type:string;not null" json:"password"`
	Role      string    `gorm:"type:string;not null" json:"role"`
	CreatedAt time.Time `gorm:"type:time;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:time" json:"updated_at"`
}

func (u *User) MapToUserModel() *model.User {
	return &model.User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Role:     enum.EnumRole(u.Role),
	}
}

func MapToUserEntity(user model.User) *User {
	return &User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     string(user.Role),
	}
}

func (u *User) SetCreatedAt() {
	u.CreatedAt = time.Now()
}

func (u *User) SetUpdatedAt() {
	u.UpdatedAt = time.Now()
}

func (u *User) SetID() {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
}
