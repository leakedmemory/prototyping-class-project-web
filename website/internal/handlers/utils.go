package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

var sessionStore = make(map[string]string)

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
