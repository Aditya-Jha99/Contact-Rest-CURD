package entity

type Contact struct {
	ID            int    `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	ContactNumber int    `json:"contactNumber"`
}
