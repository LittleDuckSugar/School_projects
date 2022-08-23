package models

import "strings"

type CartDB struct {
	UUID     string  `json:"uuid"`
	Products []uint8 `json:"products"`
}

func (cart *CartDB) ConvertProductsToDisplay() []string {
	var output string

	for _, value := range cart.Products {
		output += string(value)
	}

	return strings.Split(output[1:len(output)-1], ",")
}
