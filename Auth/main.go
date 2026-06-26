package main

import (
	"log"
	"net/http"
	
	"auth/routes"
	"auth/services"
	"auth/utils"
)

func main() {
	log.Println("Server Main...")

	var dbConn = utils.GetConnection()
	services.SetDB(dbConn)

	Router := routes.MainRouter()

	log.Fatal(http.ListenAndServe(":5959", Router))
	log.Println("Listen on PORT 5959")
}
