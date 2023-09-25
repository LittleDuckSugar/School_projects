package controllers

import (
	"meetupAPI/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootEvent(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetAllEvent())
}