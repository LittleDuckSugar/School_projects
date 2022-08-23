package repository

import "delyaneAPI/models"

// GetNewsletters get all newsletters from db
func GetNewsletters() []models.Newsletter {
	rows, err := currentDB.Query("SELECT * FROM newsletter")

	if err != nil {
		panic(err)
	}

	var uuid string
	var email string

	var newsletters []models.Newsletter

	for rows.Next() {
		err = rows.Scan(&uuid, &email)

		if err != nil {
			panic(err)
		}

		newsletters = append(newsletters, models.Newsletter{UUID: uuid, Email: email})
	}

	return newsletters
}

// PostNewsletter create a new newsletter in db
func PostNewsletter(newnewsletter models.PostNewsletter) {
	// dynamic
	insertDynStmt := `insert into "newsletter"("email") values($1)`

	_, err := currentDB.Exec(insertDynStmt, newnewsletter.Email)
	if err != nil {
		panic(err)
	}
}
