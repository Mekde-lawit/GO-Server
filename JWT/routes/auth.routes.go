package routes

import (
	controller "jwt/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("users/signup", controller.Signup())
	router.POST("users/login", controller.Login())
}