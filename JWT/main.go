package main

import (
	"os"

	"jwt/configs"
	router "jwt/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.InitEnv()
	configs.DBinstance()
}

func main() {
	println("Server starting...")
	port := os.Getenv("PORT")

	app := gin.Default() // create new gin router with default middleware (logger and recovery)

	// handle multiple routes
	router.AuthRoutes(app)
	router.UserRoutes(app)

	app.Run(":" + port)
}
