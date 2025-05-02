package controllers

import (
	"SmartSensorClient/repository"
	"fmt"
	"html/template"
	"net/http"
)

//Handle / page
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/ page called")

	place := repository.GetPlace()

	parsedTemplate, _ := template.ParseFiles("templates/index.html", "templates/structure/header.html", "templates/structure/footer.html")
	parsedTemplate.Execute(w, place)
}
