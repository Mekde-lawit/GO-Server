package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"crud/initializers"
	"crud/models"
)

func GetPosts(c *gin.Context) {
var posts []models.Post

result := initializers.DB.Find(&posts)

if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{
		"posts":   posts,
	})
}

func CreatePosts(c *gin.Context) {

	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	result := initializers.DB.Create(&post)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

	c.JSON(http.StatusCreated, gin.H{
		"post":    post,
	})
}