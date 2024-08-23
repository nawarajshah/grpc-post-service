package controller

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-api/service"
)

type PostController struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{
		PostService: postService,
	}
}

func (c *PostController) CreatePost(ctx *gin.Context) {
	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		UserId      string `json:"user_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAt := time.Now().Unix()
	pbReq := &pb.CreatePostRequest{
		Post: &pb.Post{
			PostId:      "", // Will be generated on the server
			Title:       req.Title,
			Description: req.Description,
			UserId:      req.UserId, // Correct field name
			CreatedAt:   createdAt,  // Pass Unix timestamp directly
			UpdatedAt:   createdAt,  // Pass Unix timestamp directly
		},
	}

	res, err := c.PostService.CreatePost(ctx, pbReq)
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

func (c *PostController) UpdatePost(ctx *gin.Context) {
	var req struct {
		PostId      string `json:"post_id" binding:"required"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAt := time.Now().Unix()
	pbReq := &pb.UpdatePostRequest{
		Post: &pb.Post{
			PostId:      req.PostId,
			Title:       req.Title,
			Description: req.Description,
			UpdatedAt:   updatedAt, // Pass Unix timestamp directly
		},
	}

	res, err := c.PostService.UpdatePost(ctx, pbReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *PostController) DeletePost(ctx *gin.Context) {
	postID := ctx.Param("postId")
	req := &pb.DeletePostRequest{PostId: postID}

	_, err := c.PostService.DeletePost(context.Background(), req)
	if err != nil {
		if strings.Contains(err.Error(), "post not found") {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
