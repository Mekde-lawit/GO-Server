package main

import (
	"log"
	"net/http"
	"udemygo/routes"
	"udemygo/services"
	"udemygo/utils"
)

func main() {
	log.Println("Server Main...")

	var dbConn = utils.GetConnection()
	services.SetDB(dbConn)

	Router := routes.MainRouter()

	log.Fatal(http.ListenAndServe(":8000", Router))
	log.Println("Listen on PORT 8000")
}
