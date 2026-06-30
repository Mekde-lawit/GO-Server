package routes

import (
	controller "jwt/controllers"

	"github.com/gin-gonic/gin"
)


func AuthRoutes(router *gin.Engine) {
router.Use(middleware.Authenticate)

router.GET("/users", controller.GetUsers)
router.GET("/user/:user_id", controller.GetUser)
}