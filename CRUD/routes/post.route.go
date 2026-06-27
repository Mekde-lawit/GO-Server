package routes

import (
	"crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(router *gin.RouterGroup) {
	// Public post routes (no authentication required)
	public := router.Group("/posts")
	{
		public.GET("/", controllers.GetPosts)
		public.POST("/", controllers.CreatePosts)
		// public.GET("/:id", controllers.GetPost)
		// public.GET("/search", controllers.SearchPosts)
		// public.GET("/tags/:tag", controllers.GetPostsByTag)
		// public.GET("/category/:category", controllers.GetPostsByCategory)
	}

	// Protected post routes (authentication required)
	// protected := router.Group("/posts")
	// protected.Use(middleware.AuthMiddleware())
	// {
	// 	protected.POST("/", controllers.CreatePost)
	// 	protected.PUT("/:id", controllers.UpdatePost)
	// 	protected.DELETE("/:id", controllers.DeletePost)
	// 	protected.POST("/:id/like", controllers.LikePost)
	// 	protected.DELETE("/:id/like", controllers.UnlikePost)
	// 	protected.POST("/:id/share", controllers.SharePost)
	// 	protected.POST("/:id/report", controllers.ReportPost)
		
	// 	// User's own posts
	// 	protected.GET("/my-posts", controllers.GetMyPosts)
	// 	protected.GET("/drafts", controllers.GetDrafts)
	// 	protected.POST("/:id/publish", controllers.PublishPost)
	// 	protected.POST("/:id/save-draft", controllers.SaveDraft)
	// }

	// Optional auth routes (works with or without auth)
	// optional := router.Group("/posts")
	// optional.Use(middleware.OptionalAuthMiddleware())
	// {
	// 	optional.GET("/feed", controllers.GetFeed)
	// 	optional.GET("/popular", controllers.GetPopularPosts)
	// 	optional.GET("/recommended", controllers.GetRecommendedPosts)
	// }
}