package handlers

import (
	"net/http"
	"sync"

	"github.com/gorilla/sessions"

	"github.com/leakedmemory/prototyping-class-project/internal/db"
	"github.com/leakedmemory/prototyping-class-project/internal/monitors"
	"github.com/leakedmemory/prototyping-class-project/web/template"
)

var store = sessions.NewCookieStore([]byte("tYMhbaajY6tTeXNUyRktpnuf2Wq73d31EkHTJAKryRg="))

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
	database         *db.DB
	petMonitors      map[string]*monitors.PetMonitor
	petMonitorsMutex sync.RWMutex
}

func NewHandler(database *db.DB) *Handler {
	pms := initPetMonitors(database)

	return &Handler{
		database:    database,
		petMonitors: pms,
	}
}

func initPetMonitors(database *db.DB) map[string]*monitors.PetMonitor {
	pms := make(map[string]*monitors.PetMonitor)
	users := database.GetAllUsers()
	for _, user := range users {
		for _, pet := range user.Pets {
			pm := monitors.NewPetMonitor(pet.Name, user.Phone)
			pms[pet.LeashID] = pm
			go pm.Monitor()
		}
	}

	return pms
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
