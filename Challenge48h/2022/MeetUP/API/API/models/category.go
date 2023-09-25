package models

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PostCategory struct {
	Name string `json:"name"`
}
