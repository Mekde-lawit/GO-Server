package controllers

import (
	"jwt/models"
	service "jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login() {

}

func Signup()  gin.HandlerFunc{
return func (c *gin.Context) {
	
 var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := service.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, result)
}
}