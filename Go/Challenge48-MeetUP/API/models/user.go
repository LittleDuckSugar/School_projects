package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	Tel          string `json:"tel"`
	Age          int    `json:"age"`
	Localisation string `json:"localisation"`
}

// LoginUser is the struct used to log a user
type Login struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type CreateUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	Tel          string `json:"tel"`
	Age          int    `json:"age"`
	Localisation string `json:"localisation"`
}

func (user *CreateUser) EncryptPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
}

func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
