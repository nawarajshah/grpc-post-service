package controller

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-api/service"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		PostId      string `json:"postId" binding:"required"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		CreatedBy   string `json:"createdBy" binding:"required"`
		CreatedAt   string `json:"createdAt" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAt, err := time.Parse(time.RFC3339, req.CreatedAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid createdAt format, must be RFC3339"})
		return
	}

	pbReq := &pb.CreatePostRequest{
		Post: &pb.Post{
			PostId:      req.PostId,
			Title:       req.Title,
			Description: req.Description,
			CreatedBy:   req.CreatedBy,
			CreatedAt:   timestamppb.New(createdAt),
			UpdatedAt:   timestamppb.Now(), // Set the updatedAt to current time on creation
		},
	}

	res, err := c.PostService.CreatePost(context.Background(), pbReq)
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
		PostId      string `json:"postId" binding:"required"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		UpdatedAt   string `json:"updatedAt" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAt, err := time.Parse(time.RFC3339, req.UpdatedAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid updatedAt format, must be RFC3339"})
		return
	}

	pbReq := &pb.UpdatePostRequest{
		Post: &pb.Post{
			PostId:      req.PostId,
			Title:       req.Title,
			Description: req.Description,
			UpdatedAt:   timestamppb.New(updatedAt),
		},
	}

	res, err := c.PostService.UpdatePost(context.Background(), pbReq)
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
