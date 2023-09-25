package repository

import (
	"meetupAPI/models"
	"time"
)

func GetAllEvent() []models.EventDB {
	rows, err := currentDB.Query("SELECT * FROM event")

	if err != nil {
		panic(err)
	}

	var uuid string
	var nom string
	var description string
	var note int
	var owner string
	var eventType bool
	var category_id string
	var users []uint8
	var tags []uint8
	var date time.Time
	var duration int64

	var events []models.EventDB

	for rows.Next() {
		err = rows.Scan(&uuid, &nom, &description, &note, &owner, &eventType, &category_id, &users, &tags, &date, &duration)

		if err != nil {
			panic(err)
		}

		events = append(events, models.EventDB{Id: uuid, Name: nom, Desc: description, Note: note, Owner: owner, Type: eventType, CategoryId: category_id, Users: users, Tag: tags, Date: date, Duration: duration})
	}

	return events
}

func GetEventById(id string) models.EventDB {
	rows, err := currentDB.Query("SELECT * FROM event WHERE event_id = $1", id)

	if err != nil {
		panic(err)
	}

	var uuid string
	var nom string
	var description string
	var note int
	var owner string
	var eventType bool
	var category_id string
	var users []uint8
	var tags []uint8
	var date time.Time
	var duration int64

	for rows.Next() {
		err = rows.Scan(&uuid, &nom, &description, &note, &owner, &eventType, &category_id, &users, &tags, &date, &duration)

		if err != nil {
			panic(err)
		}
	}

	return models.EventDB{Id: uuid, Name: nom, Desc: description, Note: note, Owner: owner, Type: eventType, CategoryId: category_id, Users: users, Tag: tags, Date: date, Duration: duration}
}

func PostEvent(event models.EventPost) {
	// dynamic
	insertDynStmt := `insert into "event"("nom", "description", "note", "owner", "type","category_id","users","tags","date","duration") values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := currentDB.Exec(insertDynStmt, event.Name, event.Desc, -1, event.Owner, event.Type, event.CategoryId, event.ConvertUsersToUint8(), event.ConvertTagToUint8(), event.Date, event.Duration)
	if err != nil {
		panic(err)
	}
}
