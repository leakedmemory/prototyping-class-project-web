package handlers

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"

	"website/internal/db"
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

func (h *Handler) Root(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
