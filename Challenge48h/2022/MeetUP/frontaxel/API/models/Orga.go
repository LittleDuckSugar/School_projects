package models

import "golang.org/x/crypto/bcrypt"

type Orga struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Tel      string `json:"tel"`
	Note     int    `json:"note"`
}

type OrgaPost struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Tel      string `json:"tel"`
	Note     int    `json:"note"`
}

func (orga *Orga) EncryptPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(orga.Password), 14)
	orga.Password = string(bytes)
}
