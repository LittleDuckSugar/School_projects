package repository

import (
	"SmartSensorClient/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetSensorBySensorName(sensorName string) models.Sensor {
	resp, err := http.Get(URL + "/sensor/" + sensorName)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var sensor models.Sensor

	jsonErr := json.Unmarshal(body, &sensor)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	return sensor
}

func UpdateSensorBySensorName(sensorUpdated models.SensorPost, sensorName string) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	jsonReq, err := json.Marshal(sensorUpdated)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(http.MethodPatch, URL+"/sensor/"+sensorName, bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var reponse interface{}

	json.Unmarshal(body, &reponse)

	fmt.Println(reponse)
}
