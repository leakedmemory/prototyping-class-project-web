package handlers

import (
	"encoding/json"
	"net/http"
)

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
