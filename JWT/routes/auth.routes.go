package routes

import (
	controller "jwt/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	router.POST("user/signup", controller.Signup())
	router.POST("user/login", controller.Login())

}
