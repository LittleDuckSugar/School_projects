package controllers

import (
	"meetupAPI/models"
	"meetupAPI/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootEvent(c *gin.Context) {
	eventsDB := repository.GetAllEvent()

	var event []models.Event

	for _, eventDB := range eventsDB {
		event = append(event, models.Event{Id: eventDB.Id, Name: eventDB.Name, Desc: eventDB.Desc, Note: eventDB.Note, Owner: eventDB.Owner, Type: eventDB.Type, CategoryId: eventDB.CategoryId, Users: eventDB.ConvertUsersToDisplay(), Tag: eventDB.ConvertTagsToDisplay(), Date: eventDB.Date, Duration: eventDB.Duration})
	}

	c.JSON(http.StatusOK, event)
}

func GetEventById(c *gin.Context) {
	event := repository.GetEventById(c.Params.ByName("id"))
	if event.Id == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, models.Event{Id: event.Id, Name: event.Name, Desc: event.Desc, Note: event.Note, Owner: event.Owner, Type: event.Type, CategoryId: event.CategoryId, Users: event.ConvertUsersToDisplay(), Tag: event.ConvertTagsToDisplay(), Date: event.Date, Duration: event.Duration})
}

func PostEvent(c *gin.Context) {
	// email, _ := c.Get("email")

	// var allowedToEdit bool = false

	// if isAdmin(fmt.Sprint(email)) || isOrga(fmt.Sprint(email)) {
	// 	allowedToEdit = true
	// }

	// if !allowedToEdit {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"err": "You are not allowed to create a new event"})
	// 	return
	// }

	var input models.EventPost
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if input.Owner == "" {
	// 	input.Owner = repository.GetOrgaByEmail(fmt.Sprint(email))[0].Id
	// }

	repository.PostEvent(input)

	c.JSON(http.StatusCreated, input)
}
