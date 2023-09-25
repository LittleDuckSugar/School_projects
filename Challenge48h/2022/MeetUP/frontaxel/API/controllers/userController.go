package controllers

import (
	"meetupAPI/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserById handle /user/:id - GET
func GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetUserById(c.Params.ByName("id")))
}

// PutUserById handle /user/:id - PUT
func PutUserById(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetUserById(c.Params.ByName("id")))
}

// GetUsersd handle /users - GET
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetUsers())
}
