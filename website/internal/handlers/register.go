package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"website/internal/models"
)

func generateID() string {
	timestamp := time.Now().UnixNano()

	randomBytes := make([]byte, 4)
	if _, err := rand.Read(randomBytes); err != nil {
		panic(err)
	}

	randomHex := hex.EncodeToString(randomBytes)
	id := fmt.Sprintf("%d-%s", timestamp, randomHex)
	return id
}

func (h *Handler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	id := generateID()
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	password := r.FormValue("password")

	newUser := &models.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: password,
		PetCount: 0,
	}

	err := h.database.AddUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("HX-Redirect", "/user/"+id)
	w.WriteHeader(http.StatusOK)
}
