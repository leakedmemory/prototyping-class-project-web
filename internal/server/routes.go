package server

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/leakedmemory/prototyping-class-project/internal/db"
	"github.com/leakedmemory/prototyping-class-project/internal/handlers"
	"github.com/leakedmemory/prototyping-class-project/web/template"
)

func RegisterRoutes(database *db.DB) http.Handler {
	staticDir := "web/static"

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(staticDir))
	apiHandler := handlers.NewHandler(database)

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", apiHandler.RootHandler)
	mux.HandleFunc("/home", apiHandler.HomeHandler)
	mux.Handle("/signup", templ.Handler(template.UserSignUp()))
	mux.Handle("/login", templ.Handler(template.UserLogin()))

	mux.HandleFunc("/user/signup", apiHandler.UserSignUpHandler)
	mux.HandleFunc("/user/login", apiHandler.UserLoginHandler)
	mux.HandleFunc("/user/logout", apiHandler.UserLogoutHandler)

	mux.HandleFunc("/pet/create", apiHandler.AddPetHandler)
	mux.HandleFunc("/pet/delete/{id}", apiHandler.DeletePetHandler)
	mux.HandleFunc("/pet/ping", apiHandler.PingHandler)

	mux.HandleFunc("/test/ping", apiHandler.TestPingHandler)

	return mux
}
