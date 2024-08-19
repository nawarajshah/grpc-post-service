package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nawarajshah/grpc-post-service/post-api/controller"
	"github.com/nawarajshah/grpc-post-service/post-api/middleware"
)

func SetupRouter(postController *controller.PostController, commentController *controller.CommentController, authController *controller.AuthController, verificationController *controller.VerificationController) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		// Authentication routes
		api.POST("/signup", authController.SignUp)
		api.POST("/signin", authController.SignIn)
		api.POST("/verify-email", verificationController.VerifyEmail)

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// Post routes
			protected.POST("/posts", postController.CreatePost)
			protected.GET("/posts/:postId", postController.GetPost)
			protected.PUT("/posts/:postId", postController.UpdatePost)
			protected.DELETE("/posts/:postId", postController.DeletePost)

			// Comment routes
			protected.POST("/posts/:postId/comments", commentController.CreateComment)
			protected.GET("/posts/:postId/comments/:commentId", commentController.GetComment)
			protected.PUT("/posts/:postId/comments/:commentId", commentController.UpdateComment)
			protected.DELETE("/posts/:postId/comments/:commentId", commentController.DeleteComment)
			protected.GET("/posts/:postId/comments", commentController.ListComments)
		}
	}

	return router
}
