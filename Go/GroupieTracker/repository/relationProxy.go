package repository

import (
	"encoding/json"
	"groupie/models"
	"io"
	"net/http"
)

func GetRelationById(id string) models.Relations {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var relation models.Relations

	json.Unmarshal(body, &relation)

	return relation
}
