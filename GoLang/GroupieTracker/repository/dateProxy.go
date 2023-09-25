package repository

import (
	"encoding/json"
	"groupie/models"
	"io"
	"net/http"
)

func GetDatesById(id string) models.Dates {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + id)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var dates models.Dates

	json.Unmarshal(body, &dates)

	return dates
}
