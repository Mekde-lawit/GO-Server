package controllers

import (
	"net/http"

	service "jwt/services"

	"github.com/gin-gonic/gin"
)

/*
func GetUser() gin.HandlerFunc{
return  func(c *gin.Context){

	userID := c.Param("user_id")

	if err := service.MatchUserTypeToUid(c, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error(),})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"user_id": userID},).Decode(&user)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Internal server error!",})
		return
	}
	 c.JSON(http.StatusOK, user)
}
}
*/

func GetUser() gin.HandlerFunc {
    return func(c *gin.Context) {

        userID := c.Param("user_id")

        user, err := service.GetUserByID(c, userID)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{
                "error": err.Error(),
            })
            return
        }

        c.JSON(http.StatusOK, user)
    }
}

/*
func GetUsers() gin.HandlerFunc{
return  func(c *gin.Context){
	if err := service.CheckUserType(c,"ADMIN"); err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Internal server error!" + err.Error(),})
		return
	}
	defer cursor.Close(ctx)
	var users []models.User

	if err = cursor.All(ctx, &users); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Internal server error!" + err.Error(),})
		return
	}
	c.JSON(http.StatusOK, users)
}
}
*/

func GetUsers() gin.HandlerFunc{
return  func(c *gin.Context){
  users, err := service.GetAllUsers(c)
  if err != nil {
	c.JSON(http.StatusUnauthorized, gin.H{
	  "error": err.Error(),
	})
	return
  }
  c.JSON(http.StatusAccepted, users)
}}