package handlers

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"

	"website/internal/db"
	"website/web/template"
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
	database *db.DB
}

func NewHandler(database *db.DB) *Handler {
	return &Handler{database}
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
		http.Error(w, "Failed to fetch user data", http.StatusInternalServerError)
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
