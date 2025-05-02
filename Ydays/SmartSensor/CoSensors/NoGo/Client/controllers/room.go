package controllers

import (
	"SmartSensorClient/models"
	"SmartSensorClient/repository"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

//Handle /room page
func RoomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/room page called")

	room := models.Room{}

	if r.Method == "POST" {
		newArea, _ := strconv.Atoi(r.FormValue("roomArea"))
		roomName := r.FormValue("roomName")

		if r.FormValue("create") == "true" {
			repository.CreateNewRoom(models.Room{Name: roomName, Area: uint(newArea)})
		} else if r.FormValue("create") == "false" {
			repository.UpdateRoomByName(models.Room{Name: roomName, Area: uint(newArea)}, r.URL.Path[6:])
		}

		room = repository.GetRoomByName(r.FormValue("roomName"))
	} else {
		room = repository.GetRoomByName(r.URL.Path[6:])
	}

	// Save average temp and hum of the room
	avgTemp := 0.0
	avgHum := 0.0

	nbr := uint8(0)

	for _, sensor := range room.Sensors {
		avgTemp += sensor.SensorSettings.CurrentTemp
		avgHum += sensor.SensorSettings.CurrentHum
		nbr++
	}

	if nbr == 0 {
		room.AverageTemp = -254
		room.AverageHum = -254
	} else {
		room.AverageTemp = math.Round((avgTemp/float64(nbr))*100) / 100
		room.AverageHum = math.Round((avgHum/float64(nbr))*100) / 100
	}

	parsedTemplate, _ := template.ParseFiles("templates/room.html", "templates/structure/header.html", "templates/structure/footer.html")
	parsedTemplate.Execute(w, room)
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/deleteRoom page called")

	repository.DeleteRoomByName(r.URL.Path[12:])

	http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/updateRoom page called")

	room := repository.GetRoomByName(r.URL.Path[12:])

	parsedTemplate, _ := template.ParseFiles("templates/updateRoom.html", "templates/structure/header.html", "templates/structure/footer.html")
	parsedTemplate.Execute(w, room)
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/createRoom page called")

	parsedTemplate, _ := template.ParseFiles("templates/createRoom.html", "templates/structure/header.html", "templates/structure/footer.html")
	parsedTemplate.Execute(w, nil)
}
