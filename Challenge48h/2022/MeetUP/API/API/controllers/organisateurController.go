package controllers

import (
	"fmt"
	"log"
	"meetupAPI/models"
	"meetupAPI/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrgaById(c *gin.Context) {
	orga := repository.GetOrgaById(c.Params.ByName("id"))
	if orga.Id == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organisateur not found"})
		return
	}

	c.JSON(http.StatusOK, orga)
}

func GetOrgas(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetAllOrga())
}

func PostOrga(c *gin.Context) {
	var input models.OrgaPost

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print(input.Email, " ", input.Tel, " ", input.Username)
	if len(repository.GetOrgaByEmail(input.Email)) != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Orga with this mail already exist"})
		return
	}

	if len(repository.GetOrgaByUsername(input.Username)) != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "This username is already taken by another Orga"})
		return
	}

	if len(repository.GetOrgaByTel(input.Tel)) != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "This tel is already taken by another Orga"})
		return
	}

	input.EncryptPassword()

	repository.PostOrga(input)
	c.JSON(http.StatusCreated, input)
}

func DeleteOrgaById(c *gin.Context) {
	// email, _ := c.Get("email")

	// var allowedToEdit bool = false

	// if isAdmin(fmt.Sprint(email)) {
	// 	allowedToEdit = true
	// }

	// if !allowedToEdit {
	// 	if c.Params.ByName("id") != repository.GetOrgaByEmail(fmt.Sprint(email))[0].Id {
	// 		allowedToEdit = false
	// 	} else {
	// 		allowedToEdit = true
	// 	}
	// }

	// if !allowedToEdit {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"err": "You are not allowed to delte this user"})
	// 	return
	// }

	if repository.GetOrgaById(c.Params.ByName("id")).Username == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organisateur not found"})
		return
	}

	repository.DeleteOrgaById(c.Params.ByName("id"))
	c.String(http.StatusOK, "Orga successfully deleted")
}

// LoginOrga handle /user/login for login in a user - PUBLIC
func LoginOrga(c *gin.Context) {
	// Validate input
	var input models.Login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var orgaUsername models.Orga

	if len(repository.GetOrgaByUsername(input.Identifier)) == 1 {
		orgaUsername = repository.GetOrgaByUsername(input.Identifier)[0]
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
		return
	}

	// USER FOUND in DB

	if orgaUsername.CheckPassword(input.Password) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
	} else {

		jwtWrapper := models.JwtWrapper{
			SecretKey:       "verysecretkey",
			Issuer:          "AuthService",
			ExpirationHours: 24,
		}

		signedToken, err := jwtWrapper.GenerateToken(orgaUsername.Email)
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
			"user":  orgaUsername,
			"token": tokenResponse.Token,
		})
	}
}

// isAdmin return true if an email is linked to an admin account
func isOrga(email string) bool {
	if len(repository.GetOrgaByEmail(email)) == 1 {
		return true
	}
	return false
}
