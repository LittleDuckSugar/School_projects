package controllers

import (
	"log"
	"meetupAPI/models"
	"meetupAPI/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAdminById(c *gin.Context) {
	// email, _ := c.Get("email")

	// var allowedToEdit bool = false

	// if isAdmin(fmt.Sprint(email)) {
	// 	allowedToEdit = true
	// }

	// if !allowedToEdit {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"err": "You are not allowed to see this admin"})
	// 	return
	// }
	admin := repository.GetAdminById(c.Params.ByName("id"))
	if admin.Id == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	c.JSON(http.StatusOK, admin)
}

func GetAdmins(c *gin.Context) {
	// email, _ := c.Get("email")

	// var allowedToEdit bool = false

	// if isAdmin(fmt.Sprint(email)) {
	// 	allowedToEdit = true
	// }

	// if !allowedToEdit {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"err": "You are not allowed to get admins"})
	// 	return
	// }

	c.JSON(http.StatusOK, repository.GetAdmins())
}

// LoginAdmin handle /user/login for login in a user - PUBLIC
func LoginAdmin(c *gin.Context) {
	// Validate input
	var input models.Login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var adminUsername models.Admin

	if len(repository.GetAdminByUsername(input.Identifier)) == 1 {
		adminUsername = repository.GetAdminByUsername(input.Identifier)[0]
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
		return
	}

	// USER FOUND in DB

	if adminUsername.CheckPassword(input.Password) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
	} else {

		jwtWrapper := models.JwtWrapper{
			SecretKey:       "verysecretkey",
			Issuer:          "AuthService",
			ExpirationHours: 24,
		}

		signedToken, err := jwtWrapper.GenerateToken(adminUsername.Email)
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
			"user":  adminUsername,
			"token": tokenResponse.Token,
		})
	}
}

// isAdmin return true if an email is linked to an admin account
func isAdmin(email string) bool {
	if len(repository.GetAdminByEmail(email)) == 1 {
		return true
	}
	return false
}
