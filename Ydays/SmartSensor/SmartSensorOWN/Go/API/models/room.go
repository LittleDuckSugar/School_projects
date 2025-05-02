package models

type Room struct {
	Name    string             `json:"name"`
	Sensors map[string]*Sensor `json:"sensors"`
	Area    uint               `json:"area"`
}

type RoomInput struct {
	Name string `json:"name" binding:"required"`
	Area uint   `json:"area" binding:"required"`
}

/*
// DisplayRoom print the current room
func (room Room) String() string {
	// If the room doesn't have any sensors
	if len(room.Sensors) == 0 || room.Sensors == nil {
		return fmt.Sprintf("%s is empty of sensor for %v m²", room.RoomName, room.Area)
	}

	// The room has at least 1 sensor in
	back := fmt.Sprintf("%s : ", room.RoomName)
	number := 1
	for _, sensor := range room.Sensors {
		back += fmt.Sprintf("Sensor %v : %v\n", number, sensor)
		number++
	}
	back += fmt.Sprintf("for %v m²", room.Area)
	return back
}
*/

// AddSensor add a new sensor to the current room
func (room *Room) AddSensor(newSensor *Sensor) {
	if room.Sensors == nil {
		room.Sensors = make(map[string]*Sensor)
	}

	if _, found := room.Sensors[newSensor.SystemSettings.SensorName]; !found {
		room.Sensors[newSensor.SystemSettings.SensorName] = newSensor
	}
}

// DeleteSensor remove a sensor selected by name from the current room
func (room *Room) DeleteSensor(sensorName string) {
	delete(room.Sensors, sensorName)
}

// CreateNewRoom create a new room and add it to the current place
func CreateNewRoom(roomName string, roomArea uint) *Room {
	return &Room{Name: roomName, Area: roomArea}
}
