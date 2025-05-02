package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"SmartSensorClient/models"
	"SmartSensorClient/repository"
)

//Handle /sensor page
func SensorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/sensor page called")

	sensor := models.Sensor{}

	if r.Method == "POST" {
		etage, _ := strconv.Atoi(r.FormValue("etage"))
		delay, _ := strconv.Atoi(r.FormValue("sensorDelay"))

		var isHum, isTemp, isAllowed bool

		if r.FormValue("isHum") == "true" {
			isHum = true
		}

		if r.FormValue("isTemp") == "true" {
			isTemp = true
		}

		if r.FormValue("isAllowed") == "true" {
			isAllowed = true
		}

		repository.UpdateSensorBySensorName(models.SensorPost{
			SystemSettings: models.SystemSettingsPost{
				SensorName:      r.FormValue("sensorName"),
				RoomName:        r.FormValue("roomName"),
				Position:        r.FormValue("position"),
				Etage:           int8(etage),
				CurrentTimezone: r.FormValue("timezone")},
			SensorSettings: models.SensorSettingsPost{
				SensorDelay: uint(delay),
				IsHum:       isHum,
				IsTemp:      isTemp,
				IsAllowed:   isAllowed},
			InfluxDBConfig: models.InfluxDBConfigPost{
				InfluxDBURL:    r.FormValue("influxdbURL"),
				InfluxDBToken:  r.FormValue("influxdbTOKEN"),
				InfluxDBOrg:    r.FormValue("influxdbORG"),
				InfluxDBBucket: r.FormValue("influxdbBUCKET")}},
			r.URL.Path[8:])
		time.Sleep(30 * time.Millisecond)

		sensor = repository.GetSensorBySensorName(r.FormValue("sensorName"))
	} else {
		sensor = repository.GetSensorBySensorName(r.URL.Path[8:])
	}

	parsedTemplate, _ := template.ParseFiles("templates/sensor.html", "templates/structure/header.html", "templates/structure/footer.html")
	parsedTemplate.Execute(w, sensor)
}

func UpdateSensor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/updateSensor page called")

	sensor := repository.GetSensorBySensorName(r.URL.Path[14:])

	sensor.RoomAvailable = repository.GetRoomsName()

	parsedTemplate, _ := template.ParseFiles("templates/updateSensor.html", "templates/structure/header.html", "templates/structure/footer.html")
	parsedTemplate.Execute(w, sensor)
}
