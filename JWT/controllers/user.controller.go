package controllers

import (
	service "jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// "jwt/configs"

// "github.com/go-playground/validator/v10"
// "go.mongodb.org/mongo-driver/v2/mongo"

// var userCollection *mongo.Collection = configs.OpenCollection(configs.DBinstance(), "users")
// var validate = validator.New()

func GetUser() gin.HandlerFunc{
return  func(c *gin.Context){
	userID := c.Param("user_id")

	if err := service.MatchUserTypeToUid(c, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H("error":""))
	}
}
}

func GetUsers() {
return  func(c *gin.Context){
	
}
}
