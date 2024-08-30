package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
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

func createSession(userID string) string {
	sessionID := fmt.Sprintf("sess-%s-%s", userID, generateID())
	sessionStore[sessionID] = userID
	return sessionID
}

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		userID, ok := session.Values["userID"].(string)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		r.Header.Set("userID", userID)
		next.ServeHTTP(w, r)
	})
}
