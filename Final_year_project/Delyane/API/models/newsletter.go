package models

// Newsletter is the struct used to return an existing newsletter
type Newsletter struct {
	UUID  string `json:"uuid"`
	Email string `json:"email"`
}

// PostNewsletter is the struct used to create and edit an existing newsletter
type PostNewsletter struct {
	Email string `json:"email"`
}
