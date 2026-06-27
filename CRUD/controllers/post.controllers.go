package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"crud/initializers"
	"crud/models"
)

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

func GetPost(c *gin.Context) {
var post models.Post

id := c.Param("id")

result := initializers.DB.First(&post, id)

if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

	c.JSON(http.StatusOK, gin.H{
		"post":    post,
	})
}

func UpdatePost(c *gin.Context) {
var post models.Post

id := c.Param("id")

result := initializers.DB.First(&post, id)

if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

	if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	result = initializers.DB.Save(&post)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{
		"post":    post,
	})
}

func DeletePost(c *gin.Context) {

	var post models.Post

	id := c.Param("id")

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}

		result = initializers.DB.Delete(&post)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Post deleted successfully",
		})
}
    