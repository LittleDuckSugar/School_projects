package repository

import "meetupAPI/models"

func GetAdminById(id string) models.Admin {
	rows, err := currentDB.Query("SELECT * FROM admin WHERE admin_id = $1", id)

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var email string
	var password string

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &email, &password)

		if err != nil {
			panic(err)
		}
	}

	return models.Admin{Id: uuid, Email: email, Password: password, Username: username}
}

func GetAdmins() []models.Admin {
	rows, err := currentDB.Query("SELECT * FROM admin")

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var email string
	var password string

	var admins []models.Admin

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &email, &password)

		if err != nil {
			panic(err)
		}

		admins = append(admins, models.Admin{Id: uuid, Username: username, Email: email, Password: password})
	}

	return admins
}

func GetAdminByEmail(mail string) []models.Admin {
	rows, err := currentDB.Query("SELECT * FROM admin")

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var email string
	var password string

	var admins []models.Admin

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &email, &password)

		if err != nil {
			panic(err)
		}

		admins = append(admins, models.Admin{Id: uuid, Username: username, Email: email, Password: password})
	}

	return admins
}

func GetAdminByUsername(name string) []models.Admin {
	rows, err := currentDB.Query("SELECT * FROM admin WHERE username = $1", name)

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var email string
	var password string

	var admins []models.Admin

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &email, &password)

		if err != nil {
			panic(err)
		}

		admins = append(admins, models.Admin{Id: uuid, Email: email, Password: password, Username: username})
	}

	return admins
}
