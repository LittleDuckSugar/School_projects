package repository

import (
	"delyaneAPI/models"
)

// GetWishlistById get a wishlist from a specific user in db
func GetWishlistById(id string) models.WishlistDB {
	rows, err := currentDB.Query(`SELECT * FROM "wishlist" WHERE uuid = $1`, id)

	if err != nil {
		panic(err)
	}

	var uuid string
	var products []uint8

	for rows.Next() {
		err = rows.Scan(&uuid, &products)

		if err != nil {
			panic(err)
		}
	}

	return models.WishlistDB{UUID: uuid, Products: products}
}

func GetWishlistByTime(timeLocation []uint8) models.WishlistDB {
	rows, err := currentDB.Query(`SELECT * FROM "wishlist" WHERE products = $1`, timeLocation)

	if err != nil {
		panic(err)
	}

	var uuid string
	var products []uint8

	for rows.Next() {
		err = rows.Scan(&uuid, &products)

		if err != nil {
			panic(err)
		}
	}

	return models.WishlistDB{UUID: uuid, Products: products}
}

// PostWishlist allows to add and delete wishlist item from db for a specific user
func PostWishlist(uuid string, timelocation []uint8) {
	// dynamic
	insertDynStmt := `insert into "wishlist"("uuid","products") values($1, $2)`

	_, err := currentDB.Exec(insertDynStmt, uuid, timelocation)
	if err != nil {
		panic(err)
	}
}

// PutWishlistById allows to add and delete wishlist item from db for a specific user
func PutWishlistById(updatedWishlist models.WishlistDB) {
	// dynamic
	updateDynStmt := `update "wishlist" SET products = $2 where uuid = $1`

	_, err := currentDB.Exec(updateDynStmt, updatedWishlist.UUID, updatedWishlist.Products)
	if err != nil {
		panic(err)
	}
}

// DeleteWishlistById delete a user from db
func DeleteWishlistById(uuid string) {
	// dynamic
	deleteDynStmt := `delete from "wishlist" where uuid = $1`

	_, err := currentDB.Exec(deleteDynStmt, uuid)
	if err != nil {
		panic(err)
	}
}

func ClearUserWishlist(uuid string) {
	// dynamic
	updateDynStmt := `update "wishlist" SET products = $2 where uuid = $1`

	_, err := currentDB.Exec(updateDynStmt, uuid, []uint8{123, 125})
	if err != nil {
		panic(err)
	}
}
