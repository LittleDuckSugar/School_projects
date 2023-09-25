package repository

import "meetupAPI/models"

func GetOrgaById(id string) models.Orga {
	rows, err := currentDB.Query("SELECT * FROM orga WHERE orga_id = $1", id)

	if err != nil {
		panic(err)
	}

	var uuid string
	var email string
	var password string
	var username string
	var tel string
	var note int

	for rows.Next() {
		err = rows.Scan(&uuid, &email, &password, &username, &tel, &note)

		if err != nil {
			panic(err)
		}
	}

	return models.Orga{Id: id, Email: email, Password: password, Username: username, Tel: tel, Note: note}
}

func GetAllOrga() []models.Orga {
	rows, err := currentDB.Query("SELECT * FROM orga")

	if err != nil {
		panic(err)
	}

	var uuid string
	var email string
	var password string
	var username string
	var tel string
	var note int

	var orgas []models.Orga

	for rows.Next() {
		err = rows.Scan(&uuid, &email, &password, &username, &tel, &note)

		if err != nil {
			panic(err)
		}

		orgas = append(orgas, models.Orga{Id: uuid, Email: email, Password: password, Username: username, Tel: tel, Note: note})
	}

	return orgas
}

// GetOrgaByTel return a orga from db using tel
func GetOrgaByTel(Tel string) []models.Orga {
	rows, err := currentDB.Query(`SELECT * FROM "orga" WHERE tel = $1`, Tel)

	if err != nil {
		panic(err)
	}

	var uuid string
	var email string
	var password string
	var username string
	var tel string
	var note int

	var users []models.Orga

	for rows.Next() {
		err = rows.Scan(&uuid, &email, &password, &username, &tel, &note)

		if err != nil {
			panic(err)
		}

		users = append(users, models.Orga{Id: uuid, Email: email, Password: password, Username: username, Tel: tel, Note: note})
	}

	return users
}

// GetOrgaByEmail return a orga from db using email
func GetOrgaByEmail(mail string) []models.Orga {
	rows, err := currentDB.Query(`SELECT * FROM "orga" WHERE email = $1`, mail)

	if err != nil {
		panic(err)
	}

	var uuid string
	var email string
	var password string
	var username string
	var tel string
	var note int

	var users []models.Orga

	for rows.Next() {
		err = rows.Scan(&uuid, &email, &password, &username, &tel, &note)

		if err != nil {
			panic(err)
		}

		users = append(users, models.Orga{Id: uuid, Email: email, Password: password, Username: username, Tel: tel, Note: note})
	}

	return users
}

// GetOrgaByUsername return a Orga from db using username
func GetOrgaByUsername(name string) []models.Orga {
	rows, err := currentDB.Query(`SELECT * FROM "orga" WHERE username = $1`, name)

	if err != nil {
		panic(err)
	}

	var uuid string
	var email string
	var password string
	var username string
	var tel string
	var note int

	var users []models.Orga

	for rows.Next() {
		err = rows.Scan(&uuid, &email, &password, &username, &tel, &note)

		if err != nil {
			panic(err)
		}

		users = append(users, models.Orga{Id: uuid, Email: email, Password: password, Username: username, Tel: tel, Note: note})
	}

	return users
}

// PostOrga create a new Orga in db
func PostOrga(newUser models.OrgaPost) {
	// dynamic
	insertDynStmt := `insert into "orga"("email", "password", "username", "tel", "note") values($1, $2, $3, $4, $5)`

	_, err := currentDB.Exec(insertDynStmt, newUser.Email, newUser.Password, newUser.Username, newUser.Tel, -1)
	if err != nil {
		panic(err)
	}
}

// DeleteOrgaById delete an existing Orga in the db
func DeleteOrgaById(uuid string) {
	// dynamic
	deleteDynStmt := `delete from "orga" where orga_id = $1`

	_, err := currentDB.Exec(deleteDynStmt, uuid)
	if err != nil {
		panic(err)
	}
}
