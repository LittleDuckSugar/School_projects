package controller

import (
	"groupie/repository"
	"net/http"
	"text/template"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/artist.html", "templates/structure/header.html", "templates/structure/footer.html")
	parsedTemplate.Execute(w, repository.GetArtistById(r.URL.RequestURI()[8:]))
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/artists.html", "templates/structure/header.html", "templates/structure/footer.html")
	parsedTemplate.Execute(w, repository.GetArtists())
}
