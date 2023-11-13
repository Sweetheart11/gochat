package main

import (
	"log"

	"github.com/Sweetheart11/gochat/internal/user"
	"github.com/Sweetheart11/gochat/router"
	"github.com/Sweetheart11/gochat/storage"
)

func main() {
	dbConn, err := storage.NewDB()
	if err != nil {
		log.Fatalf("could not initialize database connection: %v", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userServ := user.NewService(userRep)
	userHandler := user.NewHandler(userServ)

	router.InitRouter(userHandler)

	router.Start(":8080")
}
