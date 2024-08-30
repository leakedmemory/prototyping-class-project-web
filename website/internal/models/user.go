package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Pets     []Pet  `json:"pets"`
}

type Pet struct {
	ID      string `json:"id"`
	LeashID string `json:"leash_id"`
	Name    string `json:"name"`
	Age     uint   `json:"age"`
	Type    string `json:"type"`
	Breed   string `json:"breed"`
}
