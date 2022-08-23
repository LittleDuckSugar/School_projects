package repository

import (
	"delyaneAPI/models"
)

// GetAdminByUsername return a user from db using username
func GetAdminByUsername(name string) []models.Admin {
	rows, err := currentDB.Query(`SELECT * FROM "admin" WHERE username = $1`, name)

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var password string
	var email string

	var admins []models.Admin

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &password, &email)

		if err != nil {
			panic(err)
		}

		admins = append(admins, models.Admin{UUID: uuid, Username: username, Password: password, Email: email})
	}

	return admins
}

// GetAdminByEmail return a user from db using username
func GetAdminByEmail(mail string) []models.Admin {
	rows, err := currentDB.Query(`SELECT * FROM "admin" WHERE email = $1`, mail)

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var password string
	var email string

	var admins []models.Admin

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &password, &email)

		if err != nil {
			panic(err)
		}

		admins = append(admins, models.Admin{UUID: uuid, Username: username, Password: password, Email: email})
	}

	return admins
}
