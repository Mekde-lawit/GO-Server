package controllers

import (
	"context"
	"jwt/models"
	service "jwt/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
func GetUser() gin.HandlerFunc{
return  func(c *gin.Context){
	userID := c.Param("user_id")

	if err := service.MatchUserTypeToUid(c, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H("error": err.Error()))
		return
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var user models.User 
	err := userCollection.FindOne(ctx, bson.M("user_id": userID)).Decode(&user)
	defer cancel()
	if err != nil{
		
	}
}
}

func GetUsers() {
return  func(c *gin.Context){
	
}
}
