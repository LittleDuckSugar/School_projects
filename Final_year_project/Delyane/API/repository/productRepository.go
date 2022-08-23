package repository

import (
	"database/sql"
	"delyaneAPI/models"
)

// GetProductById return a unique product with id from db
func GetProductById(id string) models.Product {
	rows, err := currentDB.Query("SELECT * FROM product WHERE uuid = $1", id)

	if err != nil {
		panic(err)
	}

	var uuid string
	var title string
	var description string
	var price uint
	var image string
	var uuid_category sql.NullString
	var uuid_user string
	var technical sql.NullString
	var dimension sql.NullString
	var authentification sql.NullString
	var support sql.NullString

	for rows.Next() {
		err = rows.Scan(&uuid, &title, &description, &price, &image, &uuid_category, &uuid_user, &technical, &dimension, &authentification, &support)

		if err != nil {
			panic(err)
		}
	}

	return models.Product{UUID: uuid, Title: title, Description: description, Price: price, Image: image, UUID_category: uuid_category.String, UUID_user: uuid_user, Technical: technical.String, Dimension: dimension.String, Authentification: authentification.String, Support: support.String}
}

// GetProductByTitle return a unique product with title from db
func GetProductByTitle(name string) []models.Product {
	rows, err := currentDB.Query("SELECT * FROM product WHERE title = $1", name)

	if err != nil {
		panic(err)
	}

	var uuid string
	var title string
	var description string
	var price uint
	var image string
	var uuid_category sql.NullString
	var uuid_user string
	var technical sql.NullString
	var dimension sql.NullString
	var authentification sql.NullString
	var support sql.NullString

	var products []models.Product

	for rows.Next() {
		err = rows.Scan(&uuid, &title, &description, &price, &image, &uuid_category, &uuid_user, &technical, &dimension, &authentification, &support)

		if err != nil {
			panic(err)
		}

		products = append(products, models.Product{UUID: uuid, Title: title, Description: description, Price: price, Image: image, UUID_category: uuid_category.String, UUID_user: uuid_user, Technical: technical.String, Dimension: dimension.String, Authentification: authentification.String, Support: support.String})
	}

	return products
}

// GetProducts return all products from db
func GetProducts() []models.Product {
	rows, err := currentDB.Query("SELECT * FROM product")

	if err != nil {
		panic(err)
	}

	var uuid string
	var title string
	var description string
	var price uint
	var image string
	var uuid_category sql.NullString
	var uuid_user string
	var technical sql.NullString
	var dimension sql.NullString
	var authentification sql.NullString
	var support sql.NullString

	var products []models.Product

	for rows.Next() {
		err = rows.Scan(&uuid, &title, &description, &price, &image, &uuid_category, &uuid_user, &technical, &dimension, &authentification, &support)

		if err != nil {
			panic(err)
		}

		products = append(products, models.Product{UUID: uuid, Title: title, Description: description, Price: price, Image: image, UUID_category: uuid_category.String, UUID_user: uuid_user, Technical: technical.String, Dimension: dimension.String, Authentification: authentification.String, Support: support.String})
	}

	return products
}

// GetProductsGetProductsByCategory return all products from db linked to a category
func GetProductsByCategory(category string) []models.Product {
	var rows *sql.Rows
	var err error

	if category == "null" {
		rows, err = currentDB.Query("SELECT * FROM product WHERE uuid_category IS NULL")
	} else {
		rows, err = currentDB.Query("SELECT * FROM product WHERE uuid_category = $1", category)
	}

	if err != nil {
		panic(err)
	}

	var uuid string
	var title string
	var description string
	var price uint
	var image string
	var uuid_category sql.NullString
	var uuid_user string
	var technical sql.NullString
	var dimension sql.NullString
	var authentification sql.NullString
	var support sql.NullString

	var products []models.Product

	for rows.Next() {
		err = rows.Scan(&uuid, &title, &description, &price, &image, &uuid_category, &uuid_user, &technical, &dimension, &authentification, &support)

		if err != nil {
			panic(err)
		}

		products = append(products, models.Product{UUID: uuid, Title: title, Description: description, Price: price, Image: image, UUID_category: uuid_category.String, UUID_user: uuid_user, Technical: technical.String, Dimension: dimension.String, Authentification: authentification.String, Support: support.String})
	}

	return products
}

// PostProduct create a new product in the db
func PostProduct(newProduct models.PostProduct, userUUID string) {
	// dynamic
	insertDynStmt := `insert into "product"("title", "description", "price", "image", "uuid_category", "uuid_user", "technical", "dimension", "authentification", "support") values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := currentDB.Exec(insertDynStmt, newProduct.Title, newProduct.Description, newProduct.Price, newProduct.Image, newProduct.UUID_category, userUUID, newProduct.Technical, newProduct.Dimension, newProduct.Authentification, newProduct.Support)
	if err != nil {
		panic(err)
	}
}

// PutProductById update an existing product in the db
func PutProductById(uuid string, updatedProduct models.PostProduct) {
	// dynamic
	updateDynStmt := `update "product" SET title = $2,  description = $3, price = $4,  image = $5, uuid_category = $6, technical = $7, dimension = $8,  authentification = $9, support = $10 where uuid = $1`

	_, err := currentDB.Exec(updateDynStmt, uuid, updatedProduct.Title, updatedProduct.Description, updatedProduct.Price, updatedProduct.Image, updatedProduct.UUID_category, updatedProduct.Technical, updatedProduct.Dimension, updatedProduct.Authentification, updatedProduct.Support)
	if err != nil {
		panic(err)
	}
}

// DeleteProductById delete an existing product in the db using id
func DeleteProductById(uuid string) {
	// dynamic
	deleteDynStmt := `delete from "product" where uuid = $1`

	_, err := currentDB.Exec(deleteDynStmt, uuid)
	if err != nil {
		panic(err)
	}
}

// GetProductByUserId return all products linked to a user
func GetProductByUserId(id string) []models.Product {
	rows, err := currentDB.Query("SELECT * FROM product WHERE uuid_user = $1", id)

	if err != nil {
		panic(err)
	}

	var uuid string
	var title string
	var description string
	var price uint
	var image string
	var uuid_category sql.NullString
	var uuid_user string
	var technical sql.NullString
	var dimension sql.NullString
	var authentification sql.NullString
	var support sql.NullString

	var products []models.Product

	for rows.Next() {
		err = rows.Scan(&uuid, &title, &description, &price, &image, &uuid_category, &uuid_user, &technical, &dimension, &authentification, &support)

		if err != nil {
			panic(err)
		}

		products = append(products, models.Product{UUID: uuid, Title: title, Description: description, Price: price, Image: image, UUID_category: uuid_category.String, UUID_user: uuid_user, Technical: technical.String, Dimension: dimension.String, Authentification: authentification.String, Support: support.String})
	}

	return products
}
