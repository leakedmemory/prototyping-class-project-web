package handlers

import (
	"log"
	"sync"
	"time"

	"github.com/leakedmemory/prototyping-class-project-web/internal/models"
)

const (
	pingInterval  = 200 * time.Millisecond
	windowSize    = 50 // 10 seconds / 200ms
	missThreshold = 15 // 30% of 50
)

type LeashMonitor struct {
	LeashID     string
	PetName     string
	LastPing    time.Time
	PingWindow  [windowSize]bool
	WindowIndex int
	MissedPings int
	mutex       sync.Mutex
}

func NewLeashMonitor(leashID, petName string) *LeashMonitor {
	return &LeashMonitor{
		LeashID:  leashID,
		PetName:  petName,
		LastPing: time.Now(),
	}
}

func (lm *LeashMonitor) RecordPing() {
	lm.mutex.Lock()
	defer lm.mutex.Unlock()

	now := time.Now()
	expectedIndex := int(now.Sub(lm.LastPing)/pingInterval) % windowSize

	for i := lm.WindowIndex; i != expectedIndex; i = (i + 1) % windowSize {
		if !lm.PingWindow[i] {
			lm.MissedPings++
		}
		lm.PingWindow[i] = false
	}

	lm.PingWindow[expectedIndex] = true
	if lm.PingWindow[lm.WindowIndex] {
		lm.MissedPings--
	}

	lm.LastPing = now
	lm.WindowIndex = (expectedIndex + 1) % windowSize
}

func (lm *LeashMonitor) Monitor() {
	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()

	for range ticker.C {
		lm.mutex.Lock()
		now := time.Now()
		if now.Sub(lm.LastPing) > pingInterval {
			expectedIndex := int(now.Sub(lm.LastPing)/pingInterval) % windowSize
			for i := lm.WindowIndex; i != expectedIndex; i = (i + 1) % windowSize {
				if !lm.PingWindow[i] {
					lm.MissedPings++
				}
				lm.PingWindow[i] = false
			}
			lm.WindowIndex = expectedIndex
		}

		if lm.MissedPings > missThreshold {
			log.Printf(
				"Pet %s with leash %s may have run away! Missing %d pings in the last 10 seconds.",
				lm.PetName, lm.LeashID, lm.MissedPings,
			)
		}
		lm.mutex.Unlock()
	}
}

func (h *Handler) StartMonitoringLeash(pet *models.Pet) {
	if pet.LeashID == "" {
		return
	}

	h.leashMutex.Lock()
	defer h.leashMutex.Unlock()

	if _, exists := h.leashMonitors[pet.LeashID]; !exists {
		monitor := NewLeashMonitor(pet.LeashID, pet.Name)
		h.leashMonitors[pet.LeashID] = monitor
		go monitor.Monitor()
	}
}
