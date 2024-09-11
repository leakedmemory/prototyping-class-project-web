package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var pingData struct {
		LeashID string `json:"leash_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&pingData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	h.leashMutex.RLock()
	monitor, exists := h.leashMonitors[pingData.LeashID]
	h.leashMutex.RUnlock()

	if !exists {
		http.Error(w, "Unknown leash ID", http.StatusNotFound)
		return
	}

	monitor.RecordPing()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Ping received"})
}

func (h *Handler) TestPingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello from Go server!"))
	if err != nil {
		panic("pqp")
	}
}
