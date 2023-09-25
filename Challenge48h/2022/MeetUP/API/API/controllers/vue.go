package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"meetupAPI/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	resp, err := http.Get("http://localhost:8080/api/")
	if err != nil {
		ShowErrorPage(c, 404)
	} else {
		var events []models.Event
		defer resp.Body.Close()
		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		err = json.Unmarshal(body, &events)
		if err != nil {
			log.Fatalln(err)
		}
		c.HTML(http.StatusOK, "index.html", gin.H{"events": events})
	}
}

type user struct {
	Username string
	Email    string
	Tel      string
	Age      int
	Location string
	Password string
}

type orga struct {
	Username string
	Email    string
	Tel      string
	Password string
}

func ShowRegisterUserPage(c *gin.Context) {
	var err error
	var resp *http.Response
	var result map[string]interface{}
	res := user{c.PostForm("username"), c.PostForm("email"), c.PostForm("tel"), 0, c.PostForm("location"), c.PostForm("password")}
	res.Age, err = strconv.Atoi(c.PostForm("age"))
	if err == nil {
		data, _ := json.Marshal(res)
		resp, err = http.Post("http://localhost:8080/api/user", "application/json", bytes.NewBuffer(data))
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal(body, &result)
	}
	c.HTML(http.StatusOK, "registerUser.html", gin.H{"err": result["error"]})
}

func ShowRegisterOrgaPage(c *gin.Context) {
	var resp *http.Response
	var result map[string]interface{}
	if "" != c.PostForm("username") && "" != c.PostForm("email") && "" != c.PostForm("tel") && "" != c.PostForm("password") {
		res := orga{c.PostForm("username"), c.PostForm("email"), c.PostForm("tel"), c.PostForm("password")}
		data, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}
		resp, err = http.Post("http://localhost:8080/api/orga", "application/json", bytes.NewBuffer(data))
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal(body, &result)

	}
	c.HTML(http.StatusOK, "registerOrga.html", gin.H{"err": result["error"]})
}

func ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func ShowEventPage(c *gin.Context) {
	var result models.Event
	resp, err := http.Get("http://localhost:8080/api/event/" + c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(body, &result)
	if result.Name == "" {
		ShowErrorPage(c, 400)
	} else {
		c.HTML(http.StatusOK, "event.html", gin.H{"event": result})
	}
}

func ShowLoginPage(c *gin.Context) {
	var result map[string]interface{}
	if "" != c.PostForm("username") && "" != c.PostForm("password") && "" != c.PostForm("type") {
		res := models.Login{c.PostForm("username"), c.PostForm("password")}
		var url string
		if c.PostForm("type") == "user" {
			url = "http://localhost:8080/api/user/login"
		} else {
			url = "http://localhost:8080/api/user/register"
		}
		data, err := json.Marshal(res)
		if err != nil {
			log.Fatalln(err)
		}
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal(body, &result)
		if result["token"] != nil {
			c.Writer.Header().Add("Autorization", "Bearer "+result["token"].(string))
		}
		fmt.Print(result)
	}
	c.HTML(http.StatusOK, "login.html", gin.H{"err": result["error"]})
}

func ShowProfilPage(c *gin.Context) {
	// id, err := c.Cookie("id")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(id)

	// if id == "" {
	// 	c.Redirect(http.StatusSeeOther, "/login")
	// 	return
	// }

	const id string = "eb5faa30-374f-4990-bb34-b1dc9842e0b9"
	resp, err := http.Get("http://localhost:8080/api/user/" + id)
	if err != nil {
		log.Fatalln(err)
	}
	var user models.User
	defer resp.Body.Close()
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatalln(err)
	}
	c.HTML(http.StatusOK, "profil.html", gin.H{"user": user})
}

func ShowAllUserPage(c *gin.Context) {
	resp, err := http.Get("http://localhost:8080/api/users")
	if err != nil {
		log.Fatalln(err)
	}
	var users []models.User
	defer resp.Body.Close()
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &users)
	if err != nil {
		ShowErrorPage(c, 101)
	} else {
		c.HTML(http.StatusOK, "users.html", gin.H{"users": users})
	}
}

func ShowAllOrgasPage(c *gin.Context) {
	resp, err := http.Get("http://localhost:8080/api/orgas")
	if err != nil {
		log.Fatalln(err)
	}
	var orgas []models.Orga
	defer resp.Body.Close()
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &orgas)
	if err != nil {
		ShowErrorPage(c, 101)
	} else {
		c.HTML(http.StatusOK, "orgas.html", gin.H{"orgas": orgas})
	}
}

func ShowAllAdminPage(c *gin.Context) {
	resp, err := http.Get("http://localhost:8080/api/admins")
	if err != nil {
		log.Fatalln(err)
	}
	var Admins []models.Admin
	defer resp.Body.Close()
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &Admins)
	if err != nil {
		ShowErrorPage(c, 101)
	} else {
		c.HTML(http.StatusOK, "admins.html", gin.H{"admins": Admins})
	}
}

func ShowOrgaPage(c *gin.Context) {
	const id = "fd392d27-a973-401d-a99d-1539a6145c9e"
	resp, err := http.Get("http://localhost:8080/api/orga/" + id)
	if err != nil {
		log.Fatalln(err)
	}
	var Orga models.Orga
	defer resp.Body.Close()
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &Orga)
	if err != nil {
		log.Fatalln(err)
	}
	c.HTML(http.StatusOK, "organizer_profile.html", gin.H{"orga": Orga})
}

func ShowCreateEventPage(c *gin.Context) {
	var result map[string]interface{}
	fmt.Println(c.PostForm("type"))
	if c.PostForm("name") != "" {
		data, _ := json.Marshal(models.EventPost{Name: c.PostForm("name"), Desc: c.PostForm("desc"), Owner: "a59e088a-b160-4f7b-b712-66c707edf790", Type: true, CategoryId: "a59e088a-b160-4f7b-b712-66c707edf790", Users: []string{}, Tag: []string{}, Date: time.Now(), Duration: 50})
		resp, err := http.Post("http://localhost:8080/api/event", "application/json", bytes.NewBuffer(data))
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal(body, &result)
		print(result)
	}
	c.HTML(http.StatusOK, "create_event.html", gin.H{"err": result["error"]})
}

// func Show

func ShowErrorPage(c *gin.Context, err int) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{"err": err})
}
