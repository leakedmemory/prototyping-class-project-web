package models

import "time"

type Pet struct {
	ID          string    `json:"id"`
	LeashID     string    `json:"leash_id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Type        string    `json:"type"`
	Breed       string    `json:"breed"`
	QRCodePath  string    `json:"qrcode_path"`
}
