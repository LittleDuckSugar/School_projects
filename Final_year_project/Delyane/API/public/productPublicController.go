package public

import (
	"delyaneAPI/models"
	"delyaneAPI/repository"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProductById handle /product/id (GET) - PUBLIC
func GetProductById(c *gin.Context) {
	// Does the product with this ID exist
	if len(c.Params.ByName("id")) == 36 {
		if repository.GetProductById(c.Params.ByName("id")).UUID == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, repository.GetProductById(c.Params.ByName("id")))
}

// GetProducts handle /products for all products (GET) - PUBLIC
func GetProducts(c *gin.Context) {
	if category := c.DefaultQuery("category", "none"); category == "none" {
		c.JSON(http.StatusOK, repository.GetProducts())
	} else if category == "null" {
		c.JSON(http.StatusOK, repository.GetProductsByCategory(category))
	} else {
		if len(category) == 36 {
			if repository.GetCategoryById(category).UUID == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "This category doesn't exist"})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "This category doesn't exist"})
			return
		}
		c.JSON(http.StatusOK, repository.GetProductsByCategory(category))
	}
}

// PostProduct handle /product for creating a new product (POST) - PRIVATE
func PostProduct(c *gin.Context) {
	var input models.PostProduct

	input.Title = c.PostForm("title")

	input.Description = c.PostForm("description")
	price, _ := strconv.Atoi(c.PostForm("price"))
	input.Price = uint(price)
	input.UUID_category = c.PostForm("uuid_category")
	input.Technical = c.PostForm("technical")
	input.Dimension = c.PostForm("dimension")
	input.Authentification = c.PostForm("authentification")
	input.Support = c.PostForm("support")

	if repository.GetCategoryById(input.UUID_category).UUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This category doesn't exist"})
		return
	}

	// single file
	image, err := c.FormFile("image")
	if err != nil {
		panic(err)
	}

	userID := c.PostForm("uuid_user")
	if userID != "" {
		if repository.GetUserById(userID).UUID == "" {
			c.JSON(http.StatusBadRequest, "This user does not exist")
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, "A user must be specified")
		return
	}

	repository.PostProduct(input, userID)

	for _, product := range repository.GetProductByTitle(input.Title) {
		if product.Image == input.Image {
			fmt.Println("Product found")

			imageName := generateImageName(image, product.UUID)

			input.Image = "/images/products/" + imageName
			c.SaveUploadedFile(image, "./images/products/"+imageName)

			repository.PutProductById(product.UUID, input)

			break
		}
	}

	c.JSON(http.StatusCreated, input)
}

// PutProductById handle /product/id for editing an existing product (PUT) - PRIVATE
func PutProductById(c *gin.Context) {
	// Does the product with this ID exist
	if len(c.Params.ByName("id")) == 36 {
		if repository.GetProductById(c.Params.ByName("id")).UUID == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input models.PostProduct

	input.Title = c.PostForm("title")

	input.Description = c.PostForm("description")
	price, _ := strconv.Atoi(c.PostForm("price"))
	input.Price = uint(price)
	input.UUID_category = c.PostForm("uuid_category")
	input.Technical = c.PostForm("technical")
	input.Dimension = c.PostForm("dimension")
	input.Authentification = c.PostForm("authentification")
	input.Support = c.PostForm("support")

	if repository.GetCategoryById(input.UUID_category).UUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This category doesn't exist"})
		return
	}

	// single file
	image, err := c.FormFile("image")
	if err != nil {
		panic(err)
	}

	// Deleting old image
	os.Remove("." + repository.GetProductById(c.Params.ByName("id")).Image)

	imageName := generateImageName(image, c.Params.ByName("id"))

	// Upload the file to specific dst.
	c.SaveUploadedFile(image, "./images/products/"+imageName)

	input.Image = "/images/products/" + imageName

	repository.PutProductById(c.Params.ByName("id"), input)

	c.JSON(http.StatusOK, input)
}

// DeleteProductById handle /product/id for deleting an existing product (DELETE) - PRIVATE
func DeleteProductById(c *gin.Context) {
	// Does the product with this ID exist
	if len(c.Params.ByName("id")) == 36 {
		if repository.GetProductById(c.Params.ByName("id")).UUID == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	os.Remove("." + repository.GetProductById(c.Params.ByName("id")).Image)

	repository.DeleteProductById(c.Params.ByName("id"))

	c.String(http.StatusOK, "Product successfully deleted")
}
