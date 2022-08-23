package models

import "strings"

type WishlistDB struct {
	UUID     string  `json:"uuid"`
	Products []uint8 `json:"products"`
}

type WishlistAPI struct {
	UUID     string   `json:"UUID"`
	Products []string `json:"products"`
}

type PostWishlist struct {
	Products []string `json:"products"`
}

type WishlistProduct struct {
	UUID     string    `json:"uuid"`
	Products []Product `json:"products"`
}

func (wishlist *WishlistDB) ConvertProductsToDisplay() []string {
	var output string

	for _, value := range wishlist.Products {
		output += string(value)
	}

	return strings.Split(output[1:len(output)-1], ",")
}

func (wishlist *PostWishlist) ConvertProductsToPost() []uint8 {
	var output []uint8

	output = append(output, uint8('{'))

	for _, r := range strings.Join(wishlist.Products, ",") {
		output = append(output, uint8(r))
	}

	output = append(output, uint8('}'))

	return output
}
