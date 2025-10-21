package routes

import (
	"go-gin/handlers"
	"go-gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		posts := api.Group("/posts")
		{
			posts.GET("", handlers.GetPosts)
			posts.GET("/:id", handlers.GetPost)

			authenticated := posts.Group("")
			authenticated.Use(middleware.AuthMiddleware())
			{
				authenticated.POST("", handlers.CreatePost)
				authenticated.PUT("/:id", handlers.UpdatePost)
				authenticated.DELETE("/:id", handlers.DeletePost)
			}

			comments := posts.Group("/:post_id/comments")
			{
				comments.GET("", handlers.GetComments)
				comments.Use(middleware.AuthMiddleware())
				{
					comments.POST("", handlers.CreateComment)
				}
			}
		}
	}

	return router
}
