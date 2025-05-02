package models

type Place struct {
	Name  string           `json:"place_name"`
	Rooms map[string]*Room `json:"rooms"`
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
