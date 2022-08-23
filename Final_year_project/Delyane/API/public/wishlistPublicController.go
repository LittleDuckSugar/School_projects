package public

import (
	"delyaneAPI/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetWishlistById
func GetWishlistById(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetWishlistById(c.Params.ByName("id")))
}
