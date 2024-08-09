package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nawarajshah/grpc-post-service/server/pkg/controller"
)

func SetupRouter(postController *controller.PostController) *gin.Engine {
	r := gin.Default()

	// Define routes
	r.POST("/posts", postController.CreatePost)
	r.GET("/posts/:id", postController.GetPost)
	r.PUT("/posts/:id", postController.UpdatePost)
	r.DELETE("/posts/:id", postController.DeletePost)

	return r
}
