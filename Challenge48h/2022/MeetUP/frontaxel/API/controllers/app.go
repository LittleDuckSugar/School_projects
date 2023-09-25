package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Base(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "OK"})
}
