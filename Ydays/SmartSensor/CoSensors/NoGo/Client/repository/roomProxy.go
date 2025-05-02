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

func GetRoomByName(roomName string) models.Room {
	resp, err := http.Get(URL + "/room/" + roomName)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var room models.Room

	jsonErr := json.Unmarshal(body, &room)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	return room
}

func GetRoomsName() []string {
	roomAvailable := []string{}
	for roomName := range GetPlace().Rooms {
		roomAvailable = append(roomAvailable, roomName)
	}
	return roomAvailable
}

func DeleteRoomByName(roomName string) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(http.MethodDelete, URL+"/room/"+roomName, nil)
	if err != nil {
		fmt.Println(err)
	}

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

func UpdateRoomByName(roomUpdated models.Room, roomName string) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	jsonReq, err := json.Marshal(roomUpdated)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(http.MethodPatch, URL+"/room/"+roomName, bytes.NewBuffer(jsonReq))
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

func CreateNewRoom(newRoom models.Room) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	jsonReq, err := json.Marshal(newRoom)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(http.MethodPost, URL+"/room", bytes.NewBuffer(jsonReq))
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
