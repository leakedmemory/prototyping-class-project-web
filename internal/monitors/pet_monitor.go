package monitors

import (
	// "fmt"
	"log"
	"sync"
	"time"
	// "github.com/leakedmemory/prototyping-class-project/pkg/notifications"
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
	done             chan struct{}
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
		done:             make(chan struct{}),
	}
}

func (pm *PetMonitor) Monitor() {
	if pm.isBeingMonitored {
		return
	}

	pm.isBeingMonitored = true
	log.Printf("Starting monitoring for %v...\n", pm.petName)

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
				intervalFromLastPing := currentTime.Sub(pm.lastPingTime)
				if intervalFromLastPing > pm.pingWindow {
					pm.appendPing(currentTime, missed)
				}

				pm.mutex.RLock()
				{
					if len(pm.lastPings) == defaultMaxLastPingCapacity {
						missedPings := 0
						for _, ping := range pm.lastPings {
							if ping == missed {
								missedPings++
							}
						}

						if pm.disconnected(missedPings) {
							pm.isConnected = false
							pm.notifyDisconnect(currentTime)
						} else if pm.reconnected(missedPings) {
							pm.isConnected = true
							pm.notifyReconnect(currentTime)
						}
					}
				}
				pm.mutex.RUnlock()
			}
		}
	}()
}

func (pm *PetMonitor) Stop() {
	if pm.isBeingMonitored {
		close(pm.done)
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

func (pm *PetMonitor) disconnected(missedPings int) bool {
	return missedPings > pm.missThreshold && pm.isConnected
}

func (pm *PetMonitor) notifyDisconnect(t time.Time) {
	log.Printf(
		"Pet %s may have ran away at %02d:%02d\n",
		pm.petName, t.Hour(), t.Minute(),
	)

	// message := fmt.Sprintf(
	// 	"ALERTA: Seu pet %s pode ter fugido às %v:%v",
	// 	pm.petName, t.Hour(), t.Minute(),
	// )
	//
	// err := notifications.SendSMS(pm.ownerPhone, message)
	// if err != nil {
	// 	log.Printf("Failed to send SMS: %v", err)
	// }
}

func (pm *PetMonitor) reconnected(missedPings int) bool {
	return missedPings <= pm.missThreshold && !pm.isConnected
}

func (pm *PetMonitor) notifyReconnect(t time.Time) {
	log.Printf(
		"Pet %s reconnected at %02d:%02d\n",
		pm.petName, t.Hour(), t.Minute(),
	)

	// message := fmt.Sprintf(
	// 	"BOAS NOTÍCIAS: Seu pet %s se reconectou às %v:%v",
	// 	pm.petName, t.Hour(), t.Minute(),
	// )
	//
	// err := notifications.SendSMS(pm.ownerPhone, message)
	// if err != nil {
	// 	log.Printf("Failed to send SMS: %v", err)
	// }
}
