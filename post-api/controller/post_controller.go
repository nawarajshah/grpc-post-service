package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-api/service"
	"net/http"
)

type PostController struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{
		PostService: postService,
	}
}

func (p *PostController) CreatePost(ctx *gin.Context) {
	// Extract the created_by (userId) from the context
	createdBy, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var reqBody struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := &pb.CreatePostRequest{
		Post: &pb.Post{
			PostId:      "", // This can be generated in the service layer if necessary
			Title:       reqBody.Title,
			Description: reqBody.Description,
			UserId:      createdBy.(string), // Update to use `created_by`
		},
	}

	res, err := p.PostService.CreatePost(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *PostController) GetPost(ctx *gin.Context) {
	postID := ctx.Param("postId") // Use the correct parameter name
	req := &pb.GetPostRequest{PostId: postID}

	res, err := c.PostService.GetPost(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (p *PostController) UpdatePost(ctx *gin.Context) {
	// Extract the PostId from the URL parameter
	postId := ctx.Param("postId")

	// Extract the userId from the context (set by middleware)
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Bind the incoming JSON request to a struct
	var reqBody struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare the UpdatePostRequest
	req := &pb.UpdatePostRequest{
		Post: &pb.Post{
			PostId:      postId, // Set the PostId from the URL parameter
			Title:       reqBody.Title,
			Description: reqBody.Description,
			UserId:      userId.(string),
		},
	}

	// Call the service layer
	res, err := p.PostService.UpdatePost(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// Return the response
	ctx.JSON(http.StatusOK, res)
}

func (p *PostController) DeletePost(ctx *gin.Context) {
	// Extract the PostId from the URL parameter
	postId := ctx.Param("postId")

	// Extract the userId from the context (set by middleware)
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Prepare the DeletePostRequest
	req := &pb.DeletePostRequest{
		PostId: postId,          // Set the PostId from the URL parameter
		UserId: userId.(string), // Set the UserId from the context
	}

	// Call the service layer
	res, err := p.PostService.DeletePost(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// Return the response
	ctx.JSON(http.StatusOK, res)
}
