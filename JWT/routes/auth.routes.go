package routes

import (
	controller "jwt/controllers"
	middleware "jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("users/signup", controller.Signup())
	router.POST("users/login", controller.Login())
}