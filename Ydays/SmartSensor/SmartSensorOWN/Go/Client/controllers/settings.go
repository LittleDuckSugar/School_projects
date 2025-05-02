package controllers

import (
	"SmartSensorClient/models"
	"fmt"
	"html/template"
	"net/http"
)

//Handle / page
func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/settings page called")

	settings := models.InfluxDBConfig{InfluxDBURL: "http://192.168.42.43:8086", InfluxDBToken: "aToken", InfluxDBOrg: "Ynov", InfluxDBBucket: "Ydays"}

	parsedTemplate, _ := template.ParseFiles("templates/settings.html", "templates/structure/header.html", "templates/structure/footer.html")
	parsedTemplate.Execute(w, settings)
}
