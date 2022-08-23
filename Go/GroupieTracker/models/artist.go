package models

type Artist struct {
	ID              uint
	Image           string
	Name            string
	Members         []string
	CreationDate    uint
	FirstAlbum      string
	LocationURL     string `json:"locations"`
	ConcertDatesURL string `json:"concertDates"`
	RelationsURL    string `json:"relations"`
	Locations       Locations
	ConcertDates    Dates
	Relations       Relations
}
