package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/leakedmemory/prototyping-class-project/internal/db"
	"github.com/leakedmemory/prototyping-class-project/internal/server"
)

func main() {
	dbPath := os.Getenv("USER_DB")
	if dbPath == "" {
		panic("'USER_DB' environment variable not set")
	}

	database, err := db.NewDB(dbPath)
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	server := server.NewServer(database)
	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
