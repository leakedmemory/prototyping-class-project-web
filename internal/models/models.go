package models

import "time"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Pets     []Pet  `json:"pets"`
}

type Pet struct {
	ID          string    `json:"id"`
	LeashID     string    `json:"leash_id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Type        string    `json:"type"`
	Breed       string    `json:"breed"`
}

type PingData struct {
	BoardID string `json:"board_id"`
	Status  string `json:"status"`
}

// 10 seconds / 200ms = 50 slots
const PingWindowSize uint = 50

type BoardStatus struct {
	LastPing    time.Time
	Status      string
	PingWindow  [PingWindowSize]bool
	WindowIndex uint
	MissedPings uint
}
