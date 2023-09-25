package models

import (
	"golang.org/x/crypto/bcrypt"
)

type UserDB struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	Tel          string `json:"tel"`
	Localisation string `json:"localisation"`
	Age          int    `json:"age"`
}

type User struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	Tel          string `json:"tel"`
	Localisation string `json:"localisation"`
	Age          int    `json:"age"`
}

type PostUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	Tel          string `json:"tel"`
	Localisation string `json:"localisation"`
	Age          int    `json:"age"`
}

func (user *User) EncryptPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
}

