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

	var appRouter = routes.CreateRouter()

	log.Println("Listen on PORT 8000")

	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
