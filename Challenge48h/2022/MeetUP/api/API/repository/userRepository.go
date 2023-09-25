package repository

import (
	"meetupAPI/models"
)

// GetUserById return a single user from db by id
func GetUserById(id string) models.User {
	rows, err := currentDB.Query("SELECT * FROM users WHERE users_id = $1", id)

	if err != nil {
		panic(err)
	}

	var uuid string
	var email string
	var password string
	var username string
	var tel string
	var age int
	var location string

	for rows.Next() {
		err = rows.Scan(&uuid, &email, &password, &username, &tel, &age, &location)

		if err != nil {
			panic(err)
		}
	}

	return models.User{Id: id, Email: email, Password: password, Username: username, Tel: tel, Age: age, Localisation: location}
}

// GetUsers return all users saved in db
func GetUsers() []models.User {
	rows, err := currentDB.Query("SELECT * FROM users")

	if err != nil {
		panic(err)
	}

	var uuid string
	var email string
	var password string
	var username string
	var tel string
	var age int
	var location string

	var users []models.User

	for rows.Next() {
		err = rows.Scan(&uuid, &email, &password, &username, &tel, &age, &location)

		if err != nil {
			panic(err)
		}
		users = append(users, models.User{Id: uuid, Email: email, Password: password, Username: username, Tel: tel, Age: age, Localisation: location})
	}

	return users
}

// GetUserByTel return a user from db using tel
func GetUserByTel(Tel string) []models.User {
	rows, err := currentDB.Query(`SELECT * FROM "users" WHERE tel = $1`, Tel)

	if err != nil {
		panic(err)
	}

	var uuid string
	var email string
	var password string
	var username string
	var tel string
	var age int
	var location string

	var users []models.User

	for rows.Next() {
		err = rows.Scan(&uuid, &email, &password, &username, &tel, &age, &location)

		if err != nil {
			panic(err)
		}

		users = append(users, models.User{Id: uuid, Email: email, Password: password, Username: username, Tel: tel, Age: age, Localisation: location})
	}

	return users
}

// GetUserByEmail return a user from db using email
func GetUserByEmail(mail string) []models.User {
	rows, err := currentDB.Query(`SELECT * FROM "users" WHERE email = $1`, mail)

	if err != nil {
		panic(err)
	}

	var uuid string
	var email string
	var password string
	var username string
	var tel string
	var age int
	var location string

	var users []models.User

	for rows.Next() {
		err = rows.Scan(&uuid, &email, &password, &username, &tel, &age, &location)

		if err != nil {
			panic(err)
		}

		users = append(users, models.User{Id: uuid, Email: email, Password: password, Username: username, Tel: tel, Age: age, Localisation: location})
	}

	return users
}

// GetUserByUsername return a user from db using username
func GetUserByUsername(name string) []models.User {
	rows, err := currentDB.Query(`SELECT * FROM "users" WHERE username = $1`, name)

	if err != nil {
		panic(err)
	}

	var uuid string
	var email string
	var password string
	var username string
	var tel string
	var age int
	var location string

	var users []models.User

	for rows.Next() {
		err = rows.Scan(&uuid, &email, &password, &username, &tel, &age, &location)

		if err != nil {
			panic(err)
		}

		users = append(users, models.User{Id: uuid, Email: email, Password: password, Username: username, Tel: tel, Age: age, Localisation: location})
	}

	return users
}

// PostUser create a new user in db
func PostUser(newUser models.CreateUser) {
	// dynamic
	insertDynStmt := `insert into "users"("email", "password", "username", "tel", "age", "location") values($1, $2, $3, $4, $5, $6)`

	_, err := currentDB.Exec(insertDynStmt, newUser.Email, newUser.Password, newUser.Username, newUser.Tel, newUser.Age, newUser.Localisation)
	if err != nil {
		panic(err)
	}
}

// DeleteUserById delete an existing user in the db
func DeleteUserById(uuid string) {
	// dynamic
	deleteDynStmt := `delete from "users" where users_id = $1`

	_, err := currentDB.Exec(deleteDynStmt, uuid)
	if err != nil {
		panic(err)
	}
}
