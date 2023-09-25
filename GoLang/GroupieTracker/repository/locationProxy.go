package repository

import (
	"encoding/json"
	"groupie/models"
	"io"
	"net/http"
)

func GetLocationById(id string) models.Locations {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var location models.Locations

	json.Unmarshal(body, &location)

	return location
}
