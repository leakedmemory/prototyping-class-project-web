package server

import (
	"net/http"

	"github.com/a-h/templ"

	"website/internal/db"
	"website/internal/handlers"
	"website/web/template"
)

func RegisterRoutes(database *db.DB) http.Handler {
	staticDir := "web/static"

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(staticDir))
	apiHandler := handlers.NewHandler(database)

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.Handle("/user/signup", templ.Handler(template.UserSignUp()))
	mux.HandleFunc("/user", apiHandler.RegisterUserHandler)

	return mux
}
