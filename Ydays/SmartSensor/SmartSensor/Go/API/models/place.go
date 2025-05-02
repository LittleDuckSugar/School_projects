package models

type Place struct {
	Name  string           `json:"place_name"`
	Rooms map[string]*Room `json:"rooms"`
}

// addRoom add an existing room to the current place
func (place *Place) AddRoom(newRoom *Room) {
	if place.Rooms == nil {
		place.Rooms = make(map[string]*Room)
	}

	if _, found := place.Rooms[newRoom.Name]; !found {
		place.Rooms[newRoom.Name] = newRoom
	}
}

// DeleteRoom remove a room selected by name from the current place
func (place *Place) DeleteRoom(roomName string) {
	delete(place.Rooms, roomName)
}

func (place *Place) FindSensor(sensorName string) (*Room, bool) {
	for _, room := range place.Rooms {
		for _, sensor := range room.Sensors {
			if sensor.SystemSettings.SensorName == sensorName {
				return room, true
			}
		}
	}
	return &Room{}, false
}

/*
// Implements Stringer to print the current place
func (place Place) String() string {
	if place.Rooms == nil {
		return "Place was not setup"
	} else if len(place.Rooms) == 0 {
		return "Place is empty, no room here"
	}
	number := 1
	back := "Showing all rooms:\n"
	for _, roomContent := range place.Rooms {
		back += fmt.Sprintf("Room %v : %v\n", number, roomContent)
		number++
	}
	return back
}
*/
