package main

import (
	"fmt"

	"github.com/leakedmemory/prototyping-class-project-web/internal/db"
	"github.com/leakedmemory/prototyping-class-project-web/internal/server"
)

func main() {
	database, err := db.NewDB("user.db.json")
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	server := server.NewServer(database)
	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
