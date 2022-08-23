package repository

import "delyaneAPI/models"

// GetCategoryById get a single category by id from db
func GetCategoryById(id string) models.Category {
	rows, err := currentDB.Query("SELECT * FROM category WHERE uuid = $1", id)

	if err != nil {
		panic(err)
	}

	var uuid string
	var name string

	for rows.Next() {
		err = rows.Scan(&uuid, &name)

		if err != nil {
			panic(err)
		}
	}

	return models.Category{UUID: uuid, Name: name}
}

// GetCategoryByName get a single category by name from db
func GetCategoryByName(title string) models.Category {
	rows, err := currentDB.Query("SELECT * FROM category WHERE name = $1", title)

	if err != nil {
		panic(err)
	}

	var uuid string
	var name string

	for rows.Next() {
		err = rows.Scan(&uuid, &name)

		if err != nil {
			panic(err)
		}
	}

	return models.Category{UUID: uuid, Name: name}
}

// GetCategories get all categories from db
func GetCategories() []models.Category {
	rows, err := currentDB.Query("SELECT * FROM category")

	if err != nil {
		panic(err)
	}

	var uuid string
	var name string

	var categories []models.Category

	for rows.Next() {
		err = rows.Scan(&uuid, &name)

		if err != nil {
			panic(err)
		}

		categories = append(categories, models.Category{UUID: uuid, Name: name})
	}

	return categories
}

// PostCategory create a new category in db
func PostCategory(newCategory models.PostCategory) {
	// dynamic
	insertDynStmt := `insert into "category"("name") values($1)`

	_, err := currentDB.Exec(insertDynStmt, newCategory.Name)
	if err != nil {
		panic(err)
	}
}

// PutCategoryById update an existing category in the db
func PutCategoryById(uuid string, updatedCategory models.PostCategory) {
	// dynamic
	updateDynStmt := `update "category" SET name = $2 where uuid = $1`

	_, err := currentDB.Exec(updateDynStmt, uuid, updatedCategory.Name)
	if err != nil {
		panic(err)
	}
}

// DeleteCategoryById delete an existing category in the db
func DeleteCategoryById(uuid string) {
	// dynamic
	deleteDynStmt := `delete from "category" where uuid = $1`

	_, err := currentDB.Exec(deleteDynStmt, uuid)
	if err != nil {
		panic(err)
	}
}
