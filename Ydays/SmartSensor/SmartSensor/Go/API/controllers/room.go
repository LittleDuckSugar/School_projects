package controllers

import (
	"SmartSensorServer/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRoom provide /room/roomName with a specific room selected by name (GET)
func GetRoom(c *gin.Context) {
	roomName := c.Params.ByName("roomName")
	room, found := Place.Rooms[roomName]
	if found {
		c.JSON(http.StatusOK, room)
	} else {
		c.String(http.StatusNotFound, "Room '%s' not found", roomName)
	}
}

// PostRoom provide /room to create a new room to the place
func PostRoom(c *gin.Context) {

	// Validate input
	var input models.RoomInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRoom := models.CreateNewRoom(input.Name, input.Area)

	Place.AddRoom(newRoom)

	c.JSON(http.StatusCreated, newRoom)
}

// DeleteRoom provide /room/roomName to delete an existing room from the place
func DeleteRoom(c *gin.Context) {
	roomName := c.Params.ByName("roomName")
	if _, found := Place.Rooms[roomName]; found {
		Place.DeleteRoom(roomName)
		c.JSON(http.StatusOK, gin.H{"message": roomName + " deleted"})
	} else {
		c.String(http.StatusNotFound, "Room '%s' not found", roomName)
	}
}

// PatchRoom provide /room/roomName to patch an existing room from the place
func PatchRoom(c *gin.Context) {

	// Validate input
	var input models.RoomInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomName := c.Params.ByName("roomName")

	if _, found := Place.Rooms[roomName]; found {
		if roomName == input.Name {
			Place.Rooms[roomName].Area = input.Area
			c.JSON(http.StatusOK, gin.H{"status": roomName + " updated"})
		} else {

			if _, found := Place.Rooms[input.Name]; !found {
				roomToUpdate := Place.Rooms[roomName]
				Place.DeleteRoom(roomName)
				roomToUpdate.Name = input.Name
				roomToUpdate.Area = input.Area
				Place.AddRoom(roomToUpdate)
				c.JSON(http.StatusOK, gin.H{"Status": "updated"})
			} else {
				c.JSON(418, gin.H{"Status": input.Name + " already exist. Cannot update " + roomName})
			}
		}
	} else {
		c.String(http.StatusNotFound, "Room '%s' not found", roomName)
	}
}
