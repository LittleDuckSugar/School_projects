package repository

import (
	"SmartSensorClient/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var URL string

func GetPlace() models.Place {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var place models.Place

	jsonErr := json.Unmarshal(body, &place)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	return place
}
