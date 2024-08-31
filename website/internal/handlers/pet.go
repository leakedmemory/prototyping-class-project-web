package handlers

import (
	"net/http"
	"time"

	"website/internal/models"
	"website/web/template"
)

func (h *Handler) AddPetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method now allowed", http.StatusMethodNotAllowed)
		return
	}

	session, _ := store.Get(r, "e-leash-session")
	userID, ok := session.Values["userID"].(string)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	id := generateID()
	name := r.FormValue("name")
	dateOfBirthStr := r.FormValue("date-of-birth")
	petType := r.FormValue("type")
	breed := r.FormValue("breed")

	dateOfBirth, err := time.Parse("2006-01-02", dateOfBirthStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	newPet := &models.Pet{
		ID:          id,
		Name:        name,
		DateOfBirth: dateOfBirth,
		Type:        petType,
		Breed:       breed,
	}

	err = h.database.AddPet(newPet, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user, _ := h.database.GetUserByID(userID)

	template.PetList(template.UserData{
		Pets: user.Pets,
	}).Render(r.Context(), w)
}
