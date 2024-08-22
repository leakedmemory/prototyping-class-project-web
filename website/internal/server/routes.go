package server

import (
	"net/http"

	"github.com/a-h/templ"

	"website/internal/handlers"
	"website/web/template"
)

func RegisterRoutes() http.Handler {
	staticDir := "web/static"

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(staticDir))

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.Handle("/", templ.Handler(template.Hello()))
	mux.HandleFunc("/hello", handlers.HelloWorld)

	return mux
}
