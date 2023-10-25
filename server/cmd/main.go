package main

import (
	"log"

	"github.com/Sweetheart11/gochat/storage"
)

func main() {
	_, err := storage.NewDB()
	if err != nil {
		log.Fatalf("could not initialize database connection: %v", err)
	}
}
