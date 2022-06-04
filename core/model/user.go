package model

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func (u *User) SetEncryptedPassword(encryptedPassword string) {
	u.Password = encryptedPassword
}
