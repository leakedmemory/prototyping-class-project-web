package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/leakedmemory/prototyping-class-project-web/internal/db"
	"github.com/leakedmemory/prototyping-class-project-web/internal/server"
)

func main() {
	userFile := os.Getenv("USER_DB")
	database, err := db.NewDB(userFile)
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	server := server.NewServer(database)
	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
