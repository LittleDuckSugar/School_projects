package repository

import "meetupAPI/models"

func GetCategoryByName(name string) models.Category {
	return models.Category{Id: "1", Name: "Sport"}
}

func GetCategories() []models.Category {
	return []models.Category{{Id: "1", Name: "Sport"}, {Id: "2", Name: "Cuisine"}}
}

func GetCategoryById(id string) models.Category {
	return models.Category{Id: id, Name: "Sport"}
}

func PostCategory(newCategory models.PostCategory) {

}

func PutCategoryById(id string, updatedCategory models.PostCategory) {

}

func DeleteCategoryById(id string) {

}
