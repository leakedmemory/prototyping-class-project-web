package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/leakedmemory/prototyping-class-project-web/internal/db"
)

func NewServer(database *db.DB) *http.Server {
	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:8080"),
		Handler:      RegisterRoutes(database),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
