package models

type Pet struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     uint   `json:"age"`
	Type    string `json:"type"`
	Breed   string `json:"breed"`
	OwnerID string `json:"owner_id"`
	LeashID string `json:"leash_id"`
}
