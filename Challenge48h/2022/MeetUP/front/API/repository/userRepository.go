package repository

import (
	"meetupAPI/models"
)

func GetUserById(id string) models.User {
	return models.User{Id: "1", Email: "test@mail.fr", Password: "password", Username: "username", Tel: "06"}
}

func GetUsers() []models.User {
	return []models.User{{Id: "1", Email: "test@mail.fr", Password: "password", Username: "username", Tel: "06"}, {Id: "1", Email: "test@mail.fr", Password: "password", Username: "username", Tel: "06"}}
}
