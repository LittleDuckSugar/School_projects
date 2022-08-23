package repository

import (
	"database/sql"
	"delyaneAPI/models"
)

// GetUsers return a user from db
func GetUsers() []models.User {
	rows, err := currentDB.Query(`SELECT * FROM "user" ORDER BY lastname ASC`)

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var password string
	var email string
	var firstname sql.NullString
	var lastname sql.NullString
	var image sql.NullString
	var uuid_wishlist sql.NullString
	var uuid_cart sql.NullString

	var users []models.User

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &password, &email, &firstname, &lastname, &image, &uuid_wishlist, &uuid_cart)

		if err != nil {
			panic(err)
		}

		users = append(users, models.User{UUID: uuid, Username: username, Password: password, Email: email, FirstName: firstname.String, LastName: lastname.String, Image: image.String, UUID_wishlist: uuid_wishlist.String, UUID_cart: uuid_cart.String})
	}

	return users
}

// GetUserById return a unique user from db using id
func GetUserById(id string) models.User {
	rows, err := currentDB.Query(`SELECT * FROM "user" WHERE uuid = $1`, id)

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var password string
	var email string
	var firstname sql.NullString
	var lastname sql.NullString
	var image sql.NullString
	var uuid_wishlist sql.NullString
	var uuid_cart sql.NullString

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &password, &email, &firstname, &lastname, &image, &uuid_wishlist, &uuid_cart)

		if err != nil {
			panic(err)
		}
	}

	return models.User{UUID: uuid, Username: username, Password: password, Email: email, FirstName: firstname.String, LastName: lastname.String, Image: image.String, UUID_wishlist: uuid_wishlist.String, UUID_cart: uuid_cart.String}
}

// GetUserByEmail return a user from db using email
func GetUserByEmail(mail string) []models.User {
	rows, err := currentDB.Query(`SELECT * FROM "user" WHERE email = $1`, mail)

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var password string
	var email string
	var firstname sql.NullString
	var lastname sql.NullString
	var image sql.NullString
	var uuid_wishlist sql.NullString
	var uuid_cart sql.NullString

	var users []models.User

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &password, &email, &firstname, &lastname, &image, &uuid_wishlist, &uuid_cart)

		if err != nil {
			panic(err)
		}

		users = append(users, models.User{UUID: uuid, Username: username, Password: password, Email: email, FirstName: firstname.String, LastName: lastname.String, Image: image.String, UUID_wishlist: uuid_wishlist.String, UUID_cart: uuid_cart.String})
	}

	return users
}

// GetUserByUsername return a user from db using username
func GetUserByUsername(name string) []models.User {
	rows, err := currentDB.Query(`SELECT * FROM "user" WHERE username = $1`, name)

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var password string
	var email string
	var firstname sql.NullString
	var lastname sql.NullString
	var image sql.NullString
	var uuid_wishlist sql.NullString
	var uuid_cart sql.NullString

	var users []models.User

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &password, &email, &firstname, &lastname, &image, &uuid_wishlist, &uuid_cart)

		if err != nil {
			panic(err)
		}

		users = append(users, models.User{UUID: uuid, Username: username, Password: password, Email: email, FirstName: firstname.String, LastName: lastname.String, Image: image.String, UUID_wishlist: uuid_wishlist.String, UUID_cart: uuid_cart.String})
	}

	return users
}

// PostUser create a new user in db
func PostUser(newUser models.CreateUser) {
	// dynamic
	insertDynStmt := `insert into "user"("username", "password", "email", "image") values($1, $2, $3, $4)`

	_, err := currentDB.Exec(insertDynStmt, newUser.Username, newUser.Password, newUser.Email, "/images/static/profile.png")
	if err != nil {
		panic(err)
	}
}

// PutUserById update an existing user in db
func PutUserById(uuid string, updatedUser models.PostUser) {
	// dynamic
	updateDynStmt := `update "user" SET username = $2, password = $3, email = $4, firstname = $5, lastname = $6, image = $7 where uuid = $1`

	_, err := currentDB.Exec(updateDynStmt, uuid, updatedUser.Username, updatedUser.Password, updatedUser.Email, updatedUser.FirstName, updatedUser.LastName, updatedUser.Image)
	if err != nil {
		panic(err)
	}
}

// DeleteUserById delete an existing user in db
func DeleteUserById(uuid string) {
	// dynamic
	deleteDynStmt := `delete from "user" where uuid = $1`

	_, err := currentDB.Exec(deleteDynStmt, uuid)
	if err != nil {
		panic(err)
	}
}

func SetUserWishlist(uuidWishlist, uuidUser string) {
	// dynamic
	insertDynStmt := `UPDATE public."user" SET uuid_wishlist=$2::uuid WHERE uuid=$1::uuid::uuid;`

	_, err := currentDB.Exec(insertDynStmt, uuidUser, uuidWishlist)
	if err != nil {
		panic(err)
	}
}

func SetUserCart(uuidCart, uuidUser string) {
	// dynamic
	insertDynStmt := `UPDATE public."user" SET uuid_cart=$2::uuid WHERE uuid=$1::uuid::uuid;`

	_, err := currentDB.Exec(insertDynStmt, uuidUser, uuidCart)
	if err != nil {
		panic(err)
	}
}
