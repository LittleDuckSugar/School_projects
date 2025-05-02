package controllers

import (
	"SmartSensorServer/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Import repo from main
var Place models.Place

// GetRoot provide / with the entier place (GET)
func GetRoot(c *gin.Context) {
	c.JSON(http.StatusOK, Place)
}
