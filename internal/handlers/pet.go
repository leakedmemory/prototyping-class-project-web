package handlers

import (
	"net/http"
	"time"

	"github.com/leakedmemory/prototyping-class-project/internal/models"
	"github.com/leakedmemory/prototyping-class-project/internal/monitors"
	"github.com/leakedmemory/prototyping-class-project/web/template"
)

func (h *Handler) AddPetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "HTTP method not allowed", http.StatusMethodNotAllowed)
		return
	}

	session, _ := store.Get(r, "e-leash-session")
	userID, ok := session.Values["userID"].(string)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	id := generateID()
	leashID := r.FormValue("leashID")
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
		LeashID:     leashID,
		Name:        name,
		DateOfBirth: dateOfBirth,
		Type:        petType,
		Breed:       breed,
	}

	_, err = h.database.AddPet(newPet, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, _ := h.database.GetUserByID(userID)
	pm := monitors.NewPetMonitor(name, user.Phone)

	h.petMonitorsMutex.Lock()
	h.petMonitors[leashID] = pm
	h.petMonitorsMutex.Unlock()

	template.PetList(user.Pets).Render(r.Context(), w)
}

func (h *Handler) DeletePetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "HTTP method not allowed", http.StatusMethodNotAllowed)
		return
	}

	session, _ := store.Get(r, "e-leash-session")
	ownerID, ok := session.Values["userID"].(string)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	petID := r.PathValue("id")
	pet, err := h.database.DeletePet(ownerID, petID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.petMonitorsMutex.Lock()
	delete(h.petMonitors, pet.LeashID)
	h.petMonitorsMutex.Unlock()

	owner, err := h.database.GetUserByID(ownerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	template.PetList(owner.Pets).Render(r.Context(), w)
}

func (h *Handler) PetInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "HTTP method not allowed", http.StatusMethodNotAllowed)
		return
	}

	leashID := r.URL.Query().Get("leash_id")
	if leashID == "" {
		http.Error(w, "Missing 'leash_id' parameter", http.StatusBadRequest)
		return
	}

	user, err := h.database.GetUserByPetLeashID(leashID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var desiredPet *models.Pet
	for _, pet := range user.Pets {
		if pet.LeashID == leashID {
			desiredPet = &pet
			break
		}
	}

	template.PetInfo(user, desiredPet).Render(r.Context(), w)
}
