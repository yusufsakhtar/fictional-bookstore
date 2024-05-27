package models

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	ID        string `json:"id"`
	Cart      *Cart  `json:"cart"`
}
