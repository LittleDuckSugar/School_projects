package models

import "golang.org/x/crypto/bcrypt"

// User is the struct used to return an existing user
type User struct {
	UUID          string `json:"uuid"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Image         string `json:"image"`
	UUID_wishlist string `json:"uuid_wishlist"`
	UUID_cart     string `json:"uuid_cart"`
}

// PostUser is the struct used to create a new user
type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// PostUser is the struct used to update an existing user
type PostUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Image     string `json:"image"`
}

// LoginUser is the struct used to log a user
type Login struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

func (user *PostUser) EncryptPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
}

func (user *CreateUser) EncryptPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
}

func (user *Login) EncryptPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
}

func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
