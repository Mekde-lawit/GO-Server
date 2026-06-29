package main

import (
	"log"
	"net/http"

	"auth/routes"
	"auth/services"
	"auth/utils"
)


func main() {
	log.Println("Server Starting...")

// 1. Initialize once
    if err := utils.InitDB(); err != nil {
        log.Fatalf("Could not set up database: %v", err)
    }
	dbConn := utils.GetDB()     
    // Inject the pool directly into the services package variables
    services.DB = dbConn

	// 2. Start your server...
	router := routes.MainRouter()
	log.Println("Listening on PORT 5959")
	err := http.ListenAndServe(":5959", router)
	if err != nil {
		log.Fatal(err)
	}
}