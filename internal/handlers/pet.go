package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/leakedmemory/prototyping-class-project/internal/models"
	"github.com/leakedmemory/prototyping-class-project/internal/monitors"
	"github.com/leakedmemory/prototyping-class-project/pkg/encoding"
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
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	id := generateID()
	leashID := r.FormValue("leash-id")
	name := r.FormValue("name")
	dateOfBirthStr := r.FormValue("date-of-birth")
	petType := r.FormValue("type")
	breed := r.FormValue("breed")

	dateOfBirth, err := time.Parse("2006-01-02", dateOfBirthStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	_, err = h.database.GetPetByLeashID(leashID)
	if err == nil {
		http.Error(w, "Pet with this leash_id already exists", http.StatusBadRequest)
		return
	}

	env := os.Getenv("ENV")
	if env != "local" && env != "prod" {
		http.Error(w, "Failed to get ENV environment variable", http.StatusInternalServerError)
		return
	}

	qrCode, err := encoding.GenerateQRCode(leashID)
	if err != nil {
		http.Error(w, "Failed to generate QR Code", http.StatusInternalServerError)
		return
	}

	qrCodePath, err := saveQRCode(env, leashID, qrCode)
	if err != nil {
		http.Error(w, "Failed to save QR Code", http.StatusInternalServerError)
		return
	}

	newPet := &models.Pet{
		ID:          id,
		LeashID:     leashID,
		Name:        name,
		DateOfBirth: dateOfBirth,
		Type:        petType,
		Breed:       breed,
		QRCodePath:  qrCodePath,
	}

	_, err = h.database.AddPet(newPet, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	owner, _ := h.database.GetUserByID(userID)
	pm := monitors.NewPetMonitor(name, owner.Phone)

	h.petMonitorsMutex.Lock()
	h.petMonitors[leashID] = pm
	h.petMonitorsMutex.Unlock()

	pm.Monitor()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Pet-Count", fmt.Sprintf("%d", len(owner.Pets)))
	w.Header().Set("HX-Trigger", "petAdded")

	template.PetCard(newPet).Render(r.Context(), w)
}

func saveQRCode(env, leashID string, qrCode []byte) (string, error) {
	var qrCodePath string
	if env == "local" {
		qrCodePath = fmt.Sprintf("tmp/%s.png", leashID)
	} else if env == "prod" {
		qrCodePath = fmt.Sprintf("%s.png", leashID)
	}

	return qrCodePath, os.WriteFile(qrCodePath, qrCode, 0644)
}

func (h *Handler) DeletePetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "HTTP method not allowed", http.StatusMethodNotAllowed)
		return
	}

	session, _ := store.Get(r, "e-leash-session")
	ownerID, ok := session.Values["userID"].(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	petID := r.PathValue("id")
	pet, err := h.database.DeletePet(ownerID, petID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.petMonitorsMutex.Lock()
	{
		if pm, ok := h.petMonitors[pet.LeashID]; ok {
			pm.Stop()
			delete(h.petMonitors, pet.LeashID)
		}
	}
	h.petMonitorsMutex.Unlock()

	if err = os.Remove(pet.QRCodePath); err != nil {
		http.Error(w, "Could not remove associated QR Code file", http.StatusInternalServerError)
		return
	}

	owner, err := h.database.GetUserByID(ownerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("X-Pet-Count", fmt.Sprintf("%d", len(owner.Pets)))
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) PetPingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	leashID := r.URL.Query().Get("leash_id")
	if leashID == "" {
		http.Error(w, "Missing 'leash_id' parameter", http.StatusBadRequest)
		return
	}

	h.petMonitorsMutex.RLock()
	if pm, ok := h.petMonitors[leashID]; ok {
		pm.Ping()
		h.petMonitorsMutex.RUnlock()
	} else {
		http.Error(w, "Monitor not found", http.StatusInternalServerError)
		h.petMonitorsMutex.RUnlock()
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Ping received"})
}

func (h *Handler) PetInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "HTTP method not allowed", http.StatusMethodNotAllowed)
		return
	}

	leashID := r.URL.Query().Get("leash-id")
	if leashID == "" {
		http.Error(w, "Missing 'leash-id' parameter", http.StatusBadRequest)
		return
	}

	owner, err := h.database.GetUserByPetLeashID(leashID)
	if err != nil {
		http.Error(w, "Owner pet not found", http.StatusNotFound)
		return
	}

	pet, err := h.database.GetPetByLeashID(leashID)
	if err != nil {
		http.Error(w, "Pet with leash ID not found", http.StatusNotFound)
		return
	}

	template.PetInfo(owner, pet).Render(r.Context(), w)
}

func (h *Handler) PetGetQRCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "HTTP method not allowed", http.StatusMethodNotAllowed)
		return
	}

	leashID := r.URL.Query().Get("leash-id")
	if leashID == "" {
		http.Error(w, "Missing 'leash-id' parameter", http.StatusBadRequest)
		return
	}

	env := os.Getenv("ENV")
	if env != "local" && env != "prod" {
		http.Error(w, "Failed to get ENV environment variable", http.StatusInternalServerError)
		return
	}

	var qrCodePath string
	if env == "local" {
		qrCodePath = fmt.Sprintf("tmp/%s.png", leashID)
	} else if env == "prod" {
		qrCodePath = fmt.Sprintf("%s.png", leashID)
	}

	imageData, err := os.ReadFile(qrCodePath)
	if err != nil {
		http.Error(w, "Failed to read QR code image", http.StatusInternalServerError)
		return
	}

	base64Img := base64.StdEncoding.EncodeToString(imageData)

	template.PetQRCode(base64Img).Render(r.Context(), w)
}

func (h *Handler) PetConnectionStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "HTTP method not allowed", http.StatusMethodNotAllowed)
		return
	}

	leashID := r.URL.Query().Get("leash-id")
	if leashID == "" {
		http.Error(w, "Missing leash-id parameter", http.StatusBadRequest)
		return
	}

	h.petMonitorsMutex.RLock()
	pm, ok := h.petMonitors[leashID]
	h.petMonitorsMutex.RUnlock()

	if !ok {
		http.Error(w, "Pet monitor not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<div class="status-circle" data-connected="%t"></div>`, pm.IsConnected())
}
