package controllers

import (
	"log"
	"meetupAPI/models"
	"meetupAPI/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserById handle /user/:id - GET
func GetUserById(c *gin.Context) {
	user := repository.GetUserById(c.Params.ByName("id"))
	if user.Id == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// PutUserById handle /user/:id - PUT
func PutUserById(c *gin.Context) {
	user := repository.GetUserById(c.Params.ByName("id"))
	if user.Id == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// PutUserById handle /user - POST
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

	if len(repository.GetUserByTel(input.Tel)) != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "This tel is already taken by another user"})
		return
	}

	input.EncryptPassword()

	repository.PostUser(input)

	c.JSON(http.StatusCreated, input)
}

//DeleteUsers
func DeleteUserById(c *gin.Context) {
	// email, _ := c.Get("email")

	// var allowedToEdit bool = false

	// if isAdmin(fmt.Sprint(email)) {
	// 	allowedToEdit = true
	// }

	// if !allowedToEdit {
	// 	if c.Params.ByName("id") != repository.GetUserByEmail(fmt.Sprint(email))[0].Id {
	// 		allowedToEdit = false
	// 	} else {
	// 		allowedToEdit = true
	// 	}
	// }

	// if !allowedToEdit {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"err": "You are not allowed to delete this user"})
	// 	return
	// }

	if repository.GetUserById(c.Params.ByName("id")).Username == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	repository.DeleteUserById(c.Params.ByName("id"))
	c.String(http.StatusOK, "User successfully deleted")
}

// GetUsers handle /users - GET
func GetUsers(c *gin.Context) {
	// email, _ := c.Get("email")

	// if isAdmin(fmt.Sprint(email)) {
	// 	c.JSON(http.StatusOK, repository.GetUsers())
	// } else {
	// 	c.JSON(http.StatusUnauthorized, "You are not logged as admin")
	// }

	c.JSON(http.StatusOK, repository.GetUsers())
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

		// c.SetCookie("token", tokenResponse.Token, 36000, "/", "localhost", false, true)
		// c.SetCookie("id", userUsername.Id, 36000, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{
			"user":  userUsername,
			"token": tokenResponse.Token,
		})
	}
}
