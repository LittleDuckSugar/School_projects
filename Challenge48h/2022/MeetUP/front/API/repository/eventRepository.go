package repository

import "meetupAPI/models"

func GetAllEvent() []models.Event {
	return []models.Event{{Id: "1", Name: "hello_world", Localisation: "0,0,0", Owner: 1, Desc: "lorem in", Users: []uint8{1, 2, 4, 5}, Note: 8, Type: true, Tag: []string{"hello", "world"}, CategoryId: 1}, {Id: "2", Name: "world", Localisation: "11,0,0", Owner: 2, Desc: "lorem inp", Users: []uint8{1, 6, 4, 5}, Note: 2, Type: false, Tag: []string{"world", "hell"}, CategoryId: 5}}
}

func GetEventById(id string) models.Event {
	return models.Event{Id: id, Name: "hello_world", Localisation: "0,0,0", Owner: 1, Desc: "lorem in", Users: []uint8{1, 2, 4, 5}, Note: 8, Type: true, Tag: []string{"hello", "world"}, CategoryId: 1}
}
