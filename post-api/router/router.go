package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nawarajshah/grpc-post-service/post-api/controller"
)

func SetupRouter(postController *controller.PostController, commentController *controller.CommentController) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		// Post routes
		api.POST("/posts", postController.CreatePost)
		api.GET("/posts/:postId", postController.GetPost)
		api.PUT("/posts/:postId", postController.UpdatePost)
		api.DELETE("/posts/:postId", postController.DeletePost)

		// Comment routes
		api.POST("/posts/:postId/comments", commentController.CreateComment)
		api.GET("/posts/:postId/comments/:commentId", commentController.GetComment)
		api.PUT("/posts/:postId/comments/:commentId", commentController.UpdateComment)
		api.DELETE("/posts/:postId/comments/:commentId", commentController.DeleteComment)
		api.GET("/posts/:postId/comments", commentController.ListComments)
	}

	return router
}
