package entity

import "time"

type User struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"type:string;not null" json:"name"`
	Email     string    `gorm:"type:string;not null" json:"email"`
	Password  string    `gorm:"type:string;not null" json:"password"`
	CreatedAt time.Time `gorm:"type:time;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:time" json:"updated_at"`
}
