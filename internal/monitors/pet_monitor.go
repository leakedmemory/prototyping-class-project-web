package monitors

import (
	"log"
	"sync"
	"time"
)

const (
	defaultPingWindow          = 1000 * time.Millisecond
	defaultMaxLastPingCapacity = 10
	defaultMissThreshold       = 4
)

const (
	received int8 = iota
	missed
)

type PetMonitor struct {
	petName          string
	ownerPhone       string
	lastPingTime     time.Time
	pingWindow       time.Duration
	lastPings        []int8
	missThreshold    int
	mutex            sync.RWMutex
	isBeingMonitored bool
	isConnected      bool
}

func NewPetMonitor(petName, ownerPhone string) *PetMonitor {
	return &PetMonitor{
		petName:          petName,
		ownerPhone:       ownerPhone,
		lastPingTime:     time.Now(),
		pingWindow:       defaultPingWindow,
		lastPings:        make([]int8, 0, defaultMaxLastPingCapacity),
		missThreshold:    defaultMissThreshold,
		isBeingMonitored: false,
		isConnected:      true,
	}
}

func (pm *PetMonitor) Monitor() {
	if pm.isBeingMonitored {
		return
	}

	pm.isBeingMonitored = true
	log.Println("Monitoring pets...")

	ticker := time.NewTicker(pm.pingWindow)
	defer ticker.Stop()

	for range ticker.C {
		currentTime := time.Now()
		intervalFromLastPing := currentTime.Sub(pm.lastPingTime)
		if intervalFromLastPing > pm.pingWindow {
			pm.appendPing(currentTime, missed)
		}

		pm.mutex.RLock()

		if len(pm.lastPings) == defaultMaxLastPingCapacity {
			missedPings := 0
			for _, ping := range pm.lastPings {
				if ping == missed {
					missedPings++
				}
			}

			if missedPings > pm.missThreshold && pm.isConnected {
				// TODO: notify owner
				pm.isConnected = false
				log.Printf(
					"Pet %s may have ran away at %v:%v\n",
					pm.petName, currentTime.Hour(), currentTime.Minute(),
				)
			} else if missedPings <= pm.missThreshold && !pm.isConnected {
				pm.isConnected = true
				log.Printf(
					"Pet %s reconnected at %v:%v\n",
					pm.petName, currentTime.Hour(), currentTime.Minute(),
				)
			}
		}

		pm.mutex.RUnlock()
	}
}

func (pm *PetMonitor) Ping() {
	pm.appendPing(time.Now(), received)
}

func (pm *PetMonitor) appendPing(t time.Time, status int8) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	if len(pm.lastPings) == defaultMaxLastPingCapacity {
		pm.lastPings = pm.lastPings[1:]
	}
	pm.lastPings = append(pm.lastPings, status)
	pm.lastPingTime = t
}
