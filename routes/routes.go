package routes

import (
	"blog-platform/handlers"
	"blog-platform/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		r.GET("/posts", handlers.GetPosts)
		r.GET("/posts/:id", handlers.GetPost)
		r.GET("/categories", handlers.GetCategories)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/posts", handlers.CreatePost)
			protected.PUT("/posts/:id", handlers.UpdatePost)
			protected.DELETE("/posts/:id", handlers.DeletePost)
			protected.POST("/comments", handlers.CreateCategory)
			protected.POST("/posts/:postid/comments", handlers.AddComment)
		}
	}
}
