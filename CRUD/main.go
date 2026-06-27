package main

import (
	"log"
	"os"
	
	"crud/initializers"
	"crud/routes"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.InitEnv()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()	
	routes.SetupRoutes(router)

	PORT := os.Getenv("PORT")
	log.Println("Server running on port: " + PORT)
	if err := router.Run(":" + PORT); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}