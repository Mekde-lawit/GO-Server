package routes

import (
	controller "jwt/controllers"
	middleware "jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	router.Use(middleware.Authenticate())

	router.GET("/users", controller.GetUsers())
	router.GET("/user/:user_id", controller.GetUser())
}
