package main

import (
	"groupie/controller"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))) // give acces to every static files (css, js, images)

	http.HandleFunc("/artist/", controller.ArtistHandler)
	http.HandleFunc("/artists", controller.ArtistsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
