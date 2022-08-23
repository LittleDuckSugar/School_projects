package public

import (
	"delyaneAPI/models"
	"delyaneAPI/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNewsletters handle /newsletters (GET)
func GetNewsletters(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetNewsletters())
}

// PostNewsletter handle /newsletter for a new entry (POST)
func PostNewsletter(c *gin.Context) {
	// Validate input
	var input models.PostNewsletter
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, newsletter := range repository.GetNewsletters() {
		if input.Email == newsletter.Email {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}
	}

	repository.PostNewsletter(input)

	c.JSON(http.StatusCreated, input)
}
