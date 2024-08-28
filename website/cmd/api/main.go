package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"website/internal/db"
	"website/internal/server"
)

func main() {
	userFile := os.Getenv("USER_DB")
	petFile := os.Getenv("PET_DB")
	database, err := db.NewDB(userFile, petFile)
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	server := server.NewServer(database)
	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
