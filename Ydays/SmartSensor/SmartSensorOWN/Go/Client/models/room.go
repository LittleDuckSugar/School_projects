package models

type Room struct {
	Name        string             `json:"name"`
	Sensors     map[string]*Sensor `json:"sensors"`
	AverageTemp float64
	AverageHum  float64
	Area        uint `json:"area"`
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
