package handlers

import "website/internal/db"

type Handler struct {
	database *db.DB
}

func NewHandler(database *db.DB) *Handler {
	return &Handler{database}
}
