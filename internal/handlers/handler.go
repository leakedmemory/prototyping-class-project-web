package handlers

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"

	"github.com/leakedmemory/prototyping-class-project-web/internal/db"
	"github.com/leakedmemory/prototyping-class-project-web/web/template"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSIONS_COOKIE_KEY")))

const SEVEN_DAYS_IN_SECONDS = 60 * 60 * 24 * 7

func init() {
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   SEVEN_DAYS_IN_SECONDS,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}
}

type Handler struct {
	database      *db.DB
	leashMonitors map[string]*LeashMonitor
	leashMutex    sync.RWMutex
}

func NewHandler(database *db.DB) *Handler {
	h := &Handler{
		database:      database,
		leashMonitors: make(map[string]*LeashMonitor),
	}
	h.initLeashMonitors()
	return h
}

func (h *Handler) initLeashMonitors() {
	users, err := h.database.GetAllUsers()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return
	}

	for _, user := range users {
		for _, pet := range user.Pets {
			if pet.LeashID != "" {
				monitor := NewLeashMonitor(pet.LeashID, pet.Name)
				h.leashMonitors[pet.LeashID] = monitor
				go monitor.Monitor()
			}
		}
	}
}

func (h *Handler) RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Endpoint not found", http.StatusNotFound)
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "e-leash-session")
	userID, ok := session.Values["userID"].(string)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	user, err := h.database.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userData := template.UserData{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Pets:  user.Pets,
	}

	template.Home(userData).Render(r.Context(), w)
}
