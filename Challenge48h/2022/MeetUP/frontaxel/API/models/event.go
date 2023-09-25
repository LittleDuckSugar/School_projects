package models

import "strings"

type EventDB struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Desc         string  `json:"desc"`
	Note         int     `json:"note"`
	Owner        int     `json:"owner"`
	Users        []uint8 `json:"users"`
	Type         bool    `json:"type"`
	CategoryId   int     `json:"categoryid"`
	Localisation string  `json:"localisation"`
	Tag          []uint8 `json:"Tag"`
}

type Event struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Desc         string   `json:"desc"`
	Note         int      `json:"note"`
	Owner        int      `json:"owner"`
	Users        []uint8  `json:"users"`
	Type         bool     `json:"type"`
	CategoryId   int      `json:"categoryid"`
	Localisation string   `json:"localisation"`
	Tag          []string `json:"Tag"`
}

type EventPost struct {
	Name         string   `json:"name"`
	Desc         string   `json:"desc"`
	Note         int      `json:"note"`
	Owner        int      `json:"owner"`
	Users        []uint8  `json:"users"`
	Type         bool     `json:"type"`
	CategoryId   int      `json:"categoryid"`
	Localisation string   `json:"localisation"`
	Tag          []string `json:"Tag"`
}

func (event *EventDB) ConvertTagsToDisplay() []string {
	var output string

	for _, value := range event.Tag {
		output += string(value)
	}
	return strings.Split(output[1:len(output)-1], ",")
}

func (event *EventPost) ConvertProductsToPost() []uint8 {
	var output []uint8

	output = append(output, uint8('{'))

	for _, r := range strings.Join(event.Tag, ",") {
		output = append(output, uint8(r))
	}

	output = append(output, uint8('}'))

	return output
}
