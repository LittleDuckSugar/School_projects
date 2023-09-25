package models

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type EventDB struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Desc       string    `json:"desc"`
	Note       int       `json:"note"`
	Owner      string    `json:"owner"`
	Type       bool      `json:"type"`
	CategoryId string    `json:"categoryid"`
	Users      []uint8   `json:"users"`
	Tag        []uint8   `json:"tag"`
	Date       time.Time `json:"date"`
	Duration   int64     `json:"duration"`
}

type Event struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Desc       string    `json:"desc"`
	Note       int       `json:"note"`
	Owner      string    `json:"owner"`
	Type       bool      `json:"type"`
	CategoryId string    `json:"categoryid"`
	Users      []string  `json:"users"`
	Tag        []string  `json:"tag"`
	Date       time.Time `json:"date"`
	Duration   int64     `json:"duration"`
}

type EventPost struct {
	Name       string    `json:"name"`
	Desc       string    `json:"desc"`
	Owner      string    `json:"owner,omitempty"`
	Type       bool      `json:"type"`
	CategoryId string    `json:"categoryid"`
	Users      []string  `json:"users"`
	Tag        []string  `json:"tag"`
	Date       time.Time `json:"date"`
	Duration   int64     `json:"duration"`
}

func (event *Event) GetCategoryName() string {
	resp, err := http.Get("http://localhost:80/api/category/" + event.CategoryId)
	if err != nil {
		log.Fatalln(err)
	}
	var Cat Category
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &Cat)
	if err != nil {
		log.Fatalln(err)
	}
	return Cat.Name
}

func (event *Event) GetOwner() Orga {
	resp, err := http.Get("http://localhost:80/api/orga/" + event.Owner)
	if err != nil {
		log.Fatalln(err)
	}
	var Orga Orga
	defer resp.Body.Close()
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &Orga)
	if err != nil {
		log.Fatalln(err)
	}
	return Orga
}

func (event *EventDB) ConvertTagsToDisplay() []string {
	var output string

	for _, value := range event.Tag {
		output += string(value)
	}
	return strings.Split(output[1:len(output)-1], ",")
}

func (event *EventDB) ConvertUsersToDisplay() []string {
	var output string

	for _, value := range event.Users {
		output += string(value)
	}
	return strings.Split(output[1:len(output)-1], ",")
}

func (event *EventDB) ConvertDurationToSring() string {
	var output string

	for _, value := range event.Users {
		output += string(value)
	}

	return output
}

func (event *EventPost) ConvertTagToUint8() []uint8 {
	var output []uint8

	output = append(output, uint8('{'))

	for _, r := range strings.Join(event.Tag, ",") {
		output = append(output, uint8(r))
	}

	output = append(output, uint8('}'))

	return output
}

func (event *EventPost) ConvertUsersToUint8() []uint8 {
	var output []uint8

	output = append(output, uint8('{'))

	for _, r := range strings.Join(event.Users, ",") {
		output = append(output, uint8(r))
	}

	output = append(output, uint8('}'))

	return output
}
