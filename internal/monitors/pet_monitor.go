package monitors

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/leakedmemory/prototyping-class-project/pkg/notifications"
)

const (
	defaultPingWindow          = 2 * time.Second
	defaultMaxLastPingCapacity = 10
	defaultMissThreshold       = 3
)

const (
	received uint = iota
	missed
)

var brazilLocation *time.Location

func init() {
	var err error
	brazilLocation, err = time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Printf("Failed to load Brazil location: %v\n", err)
		// fallback to a fixed offset if loading the location fails
		brazilLocation = time.FixedZone("UTC-03", -3*60*60)
	}
}

type PetMonitor struct {
	petName          string
	ownerPhone       string
	lastPingTime     atomic.Int64
	pingWindow       time.Duration
	lastPings        []uint
	missThreshold    int
	pingsMutex       sync.RWMutex
	isBeingMonitored bool
	isConnected      bool
	done             chan struct{}
}

func NewPetMonitor(petName, ownerPhone string) *PetMonitor {
	return &PetMonitor{
		petName:          petName,
		ownerPhone:       ownerPhone,
		pingWindow:       defaultPingWindow,
		lastPings:        make([]uint, 0, defaultMaxLastPingCapacity),
		missThreshold:    defaultMissThreshold,
		isBeingMonitored: false,
		isConnected:      true,
		done:             make(chan struct{}, 1),
	}
}

func (pm *PetMonitor) Monitor() {
	if pm.isBeingMonitored {
		return
	}

	pm.isBeingMonitored = true
	log.Printf("Starting monitoring for %v...\n", pm.petName)
	pm.lastPingTime.Store(time.Now().UnixNano())

	go func() {
		ticker := time.NewTicker(pm.pingWindow)
		defer ticker.Stop()

		for {
			select {
			case <-pm.done:
				log.Printf("Stopping monitoring for %v...\n", pm.petName)
				pm.isBeingMonitored = false
				return
			case <-ticker.C:
				currentTime := time.Now()
				lastPing := time.Unix(0, pm.lastPingTime.Load())

				if currentTime.Sub(lastPing) > pm.pingWindow {
					pm.appendPing(missed)
				} else {
					pm.appendPing(received)
				}

				pm.checkConnectionStatus()
			}
		}
	}()
}

func (pm *PetMonitor) Stop() {
	if pm.isBeingMonitored {
		pm.done <- struct{}{}
	}
}

func (pm *PetMonitor) Ping() {
	pm.lastPingTime.Store(time.Now().UnixNano())
}

func (pm *PetMonitor) appendPing(status uint) {
	pm.pingsMutex.Lock()
	defer pm.pingsMutex.Unlock()

	if len(pm.lastPings) == defaultMaxLastPingCapacity {
		pm.lastPings = pm.lastPings[1:]
	}
	pm.lastPings = append(pm.lastPings, status)
}

func (pm *PetMonitor) checkConnectionStatus() {
	pm.pingsMutex.RLock()
	defer pm.pingsMutex.RUnlock()

	if len(pm.lastPings) < defaultMaxLastPingCapacity {
		return
	}

	consecutiveMissedPings := 0
	for i := len(pm.lastPings) - 1; i >= 0; i-- {
		if pm.lastPings[i] == received {
			break
		}
		consecutiveMissedPings++
	}

	if consecutiveMissedPings > pm.missThreshold && pm.isConnected {
		t := time.Now()
		log.Printf(
			"Pet %s may have ran away at %02d:%02d\n",
			pm.petName, t.Hour(), t.Minute(),
		)

		pm.isConnected = false
		pm.notifyDisconnect(&t)
	} else if consecutiveMissedPings <= pm.missThreshold && !pm.isConnected {
		t := time.Now()
		log.Printf(
			"Pet %s reconnected at %02d:%02d\n",
			pm.petName, t.Hour(), t.Minute(),
		)

		pm.isConnected = true
		pm.notifyReconnect(&t)
	}
}

func (pm *PetMonitor) notifyDisconnect(t *time.Time) {
	message := fmt.Sprintf(
		"ALERTA: Seu pet %s pode ter fugido às %02d:%02d.",
		pm.petName, t.In(brazilLocation).Hour(), t.In(brazilLocation).Minute(),
	)

	err := notifications.SendSMS(pm.ownerPhone, message)
	if err != nil {
		log.Printf("Failed to send SMS: %v\n", err)
	}
}

func (pm *PetMonitor) notifyReconnect(t *time.Time) {
	message := fmt.Sprintf(
		"BOAS NOTÍCIAS: Seu pet %s se reconectou às %02d:%02d.",
		pm.petName, t.In(brazilLocation).Hour(), t.In(brazilLocation).Minute(),
	)

	err := notifications.SendSMS(pm.ownerPhone, message)
	if err != nil {
		log.Printf("Failed to send SMS: %v\n", err)
	}
}

func (pm *PetMonitor) IsConnected() bool {
	return pm.isConnected
}
