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

	mux.HandleFunc("/", apiHandler.Root)
	mux.Handle("/home", handlers.IsAuthenticated(templ.Handler(template.Home())))
	mux.Handle("/signup", templ.Handler(template.UserSignUp()))
	mux.Handle("/login", templ.Handler(template.UserLogin()))

	mux.HandleFunc("/user/signup", apiHandler.UserSignUpHandler)
	mux.HandleFunc("/user/login", apiHandler.UserLoginHandler)
	mux.HandleFunc("/user/logout", apiHandler.UserLogoutHandler)

	return mux
}
