package repository

import (
	"encoding/json"
	"groupie/models"
	"io"
	"net/http"
	"strconv"
)

func GetArtistById(id string) models.Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var artist models.Artist

	json.Unmarshal(body, &artist)

	artist.Relations = GetRelationById(strconv.Itoa(int(artist.ID)))
	artist.Locations = GetLocationById(strconv.Itoa(int(artist.ID)))
	artist.ConcertDates = GetDatesById(strconv.Itoa(int(artist.ID)))

	return artist
}

func GetArtists() []models.Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var artists []models.Artist

	json.Unmarshal(body, &artists)

	for _, artist := range artists {
		artist.Locations = GetLocationById(strconv.Itoa(int(artist.ID)))
		artist.Relations = GetRelationById(strconv.Itoa(int(artist.ID)))
		artist.ConcertDates = GetDatesById(strconv.Itoa(int(artist.ID)))
	}

	return artists
}
