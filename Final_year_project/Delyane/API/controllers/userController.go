package controllers

import (
	"delyaneAPI/models"
	"delyaneAPI/repository"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetUserById handle /user/id (GET) - PUBLIC
func GetUserById(c *gin.Context) {
	if !isUserExistById(c.Params.ByName("id")) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found with this ID"})
		return
	}

	c.JSON(http.StatusOK, repository.GetUserById(c.Params.ByName("id")))
}

// GetUsers handle /users (GET) - PRIVATE
func GetUsers(c *gin.Context) {
	email, _ := c.Get("email")

	if isAdmin(fmt.Sprint(email)) {
		c.JSON(http.StatusOK, repository.GetUsers())
	} else {
		c.JSON(http.StatusUnauthorized, "You are not logged as admin")
	}
}

// PostUser handle /user for creating a new user (POST) - PUBLIC
func PostUser(c *gin.Context) {
	// Validate input
	var input models.CreateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(repository.GetUserByEmail(input.Email)) != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this mail already exist"})
		return
	}

	if len(repository.GetUserByUsername(input.Username)) != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "This username is already taken by another user"})
		return
	}

	input.EncryptPassword()

	repository.PostUser(input)

	c.JSON(http.StatusCreated, input)
}

// PutUserById handle /user/id for editing an existing user - PRIVATE
func PutUserById(c *gin.Context) {
	if !isUserExistById(c.Params.ByName("id")) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found with this ID"})
		return
	}

	email, _ := c.Get("email")

	var input models.PostUser

	input.Email = c.PostForm("email")

	var allowedToEdit bool = false

	if isAdmin(fmt.Sprint(email)) {
		allowedToEdit = true
	}

	if !allowedToEdit {
		if c.Params.ByName("id") != repository.GetUserByEmail(fmt.Sprint(email))[0].UUID {
			allowedToEdit = false
		} else {
			allowedToEdit = true
		}
	}

	if !allowedToEdit {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "You are not allowed to edit this user"})
		return
	}

	input.Username = c.PostForm("username")

	if len(repository.GetUserByEmail(input.Email)) > 1 {
		c.JSON(http.StatusConflict, gin.H{"error": "This email is already taken by another user"})
		return
	}

	if len(repository.GetUserByUsername(input.Username)) > 1 {
		c.JSON(http.StatusConflict, gin.H{"error": "This username is already taken by another user"})
		return
	}

	input.FirstName = c.PostForm("firstname")
	input.LastName = c.PostForm("lastname")
	input.Password = c.PostForm("password")
	input.EncryptPassword()

	var hasImage bool = true

	// single file
	image, err := c.FormFile("image")
	if err != nil {
		if err.Error() != "http: no such file" {
			panic(err)
		} else {
			hasImage = false
		}
	}

	if hasImage {
		// Deleting old image if available
		if imagePath := repository.GetUserById(c.Params.ByName("id")).Image; imagePath != "/images/static/profile.png" {
			os.Remove("." + imagePath)
		}

		// Saving new image
		imageName := generateImageName(image, c.Params.ByName("id"))

		// Upload the file to specific dst.
		c.SaveUploadedFile(image, "./images/users/"+imageName)

		input.Image = "/images/users/" + imageName
	} else {
		input.Image = repository.GetUserById(c.Params.ByName("id")).Image
	}

	repository.PutUserById(c.Params.ByName("id"), input)

	c.JSON(http.StatusOK, input)
}

// DeleteUserById handle /user/id for deleting an existing user - PRIVATE
func DeleteUserById(c *gin.Context) {
	if !isUserExistById(c.Params.ByName("id")) {
		c.JSON(http.StatusNotFound, gin.H{"err": "User not found with this ID"})
		return
	}

	email, _ := c.Get("email")

	var allowedToEdit bool = false

	if isAdmin(fmt.Sprint(email)) {
		allowedToEdit = true
	}

	if !allowedToEdit {
		if c.Params.ByName("id") != repository.GetUserByEmail(fmt.Sprint(email))[0].UUID {
			allowedToEdit = false
		} else {
			allowedToEdit = true
		}
	}

	if !allowedToEdit {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "You are not allowed to edit this user"})
		return
	}

	if imagePath := repository.GetUserById(c.Params.ByName("id")).Image; imagePath != "/images/static/profile.png" {
		os.Remove("." + imagePath)
	}

	for _, product := range repository.GetProductByUserId(c.Params.ByName("id")) {
		os.Remove("." + product.Image)
	}

	repository.DeleteUserById(c.Params.ByName("id"))

	c.String(http.StatusOK, "User successfully deleted")
}

// LoginUser handle /user/login for login in a user - PUBLIC
func LoginUser(c *gin.Context) {
	// Validate input
	var input models.Login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var userUsername models.User

	if len(repository.GetUserByUsername(input.Identifier)) == 1 {
		userUsername = repository.GetUserByUsername(input.Identifier)[0]
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
		return
	}

	// USER FOUND in DB

	if userUsername.CheckPassword(input.Password) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
	} else {

		jwtWrapper := models.JwtWrapper{
			SecretKey:       "verysecretkey",
			Issuer:          "AuthService",
			ExpirationHours: 24,
		}

		signedToken, err := jwtWrapper.GenerateToken(userUsername.Email)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"msg": "error signing token",
			})
			c.Abort()
			return
		}

		tokenResponse := models.LoginResponse{
			Token: signedToken,
		}

		c.JSON(http.StatusOK, gin.H{
			"user":  userUsername,
			"token": tokenResponse.Token,
		})
	}
}

// isUserExistById return true if the user exist in the db
func isUserExistById(uuid string) bool {
	if repository.GetUserById(uuid).UUID == "" {
		return false
	} else {
		return true
	}
}
