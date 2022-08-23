package models

// Category is the struct used to return an existing category
type Category struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

// PostCategory is the struct used to create and edit an existing category
type PostCategory struct {
	Name string `json:"name"`
}
