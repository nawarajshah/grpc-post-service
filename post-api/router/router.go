package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nawarajshah/grpc-post-service/post-api/controller"
)

func SetupRouter(postController *controller.PostController) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/posts", postController.CreatePost)
		api.GET("/posts/:id", postController.GetPost)
		api.PUT("/posts/:id", postController.UpdatePost)
		api.DELETE("/posts/:id", postController.DeletePost)
	}

	return router
}
