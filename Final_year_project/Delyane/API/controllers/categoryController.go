package controllers

import (
	"delyaneAPI/models"
	"delyaneAPI/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCategoryById handle /category/id (GET) - PIBLIC
func GetCategoryById(c *gin.Context) {
	if repository.GetCategoryByName(c.Params.ByName("id")).UUID != "" {
		c.JSON(http.StatusNotFound, gin.H{"err": "Category with this ID doesn't exist"})
		return
	}

	c.JSON(http.StatusOK, repository.GetCategoryById(c.Params.ByName("id")))
}

// GetCategories handle /categories (GET) - PUBLIC
func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetCategories())
}

// PostCategory handle /category for a new entry (POST) - PRIVATE
func PostCategory(c *gin.Context) {
	// Validate input
	var input models.PostCategory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if repository.GetCategoryByName(input.Name).UUID != "" {
		c.JSON(http.StatusConflict, "This category already exist")
		return
	}

	repository.PostCategory(input)

	c.JSON(http.StatusCreated, input)
}

// PutCategoryById handle /category/id for editing informations (PUT) - PRIVATE
func PutCategoryById(c *gin.Context) {
	// Validate input
	var input models.PostCategory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if repository.GetCategoryByName(input.Name).UUID != "" {
		c.JSON(http.StatusNotFound, "Category with this ID doesn't exist")
		return
	}

	repository.PutCategoryById(c.Params.ByName("id"), input)

	c.JSON(http.StatusOK, input)
}

// DeleteCategoryById handle /category/id for deleting an existing category (DELETE) - PRIVATE
func DeleteCategoryById(c *gin.Context) {
	if repository.GetCategoryById(c.Params.ByName("id")).UUID == "" {
		c.String(http.StatusNotFound, "Category with this ID doesn't exist")
		return
	}

	repository.DeleteCategoryById(c.Params.ByName("id"))

	c.String(http.StatusOK, "Category successfully deleted")
}
