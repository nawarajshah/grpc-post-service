package controller

import (
	"context"
	"net/http"
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
		Post struct {
			PostId      string `json:"postId"`
			Title       string `json:"title"`
			Description string `json:"description"`
			CreatedBy   string `json:"createdBy"`
			CreatedAt   string `json:"createdAt"`
		} `json:"post"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAt, err := time.Parse(time.RFC3339, req.Post.CreatedAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid timestamp format"})
		return
	}

	post := &pb.Post{
		PostId:      req.Post.PostId,
		Title:       req.Post.Title,
		Description: req.Post.Description,
		CreatedBy:   req.Post.CreatedBy,
		CreatedAt:   timestamppb.New(createdAt),
		UpdatedAt:   timestamppb.Now(),
	}

	grpcReq := &pb.CreatePostRequest{Post: post}

	res, err := c.PostService.CreatePost(context.Background(), grpcReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *PostController) GetPost(ctx *gin.Context) {
	postID := ctx.Param("id")
	req := &pb.GetPostRequest{PostId: postID}

	res, err := c.PostService.GetPost(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *PostController) UpdatePost(ctx *gin.Context) {
	var req struct {
		Post struct {
			PostId      string `json:"postId"`
			Title       string `json:"title"`
			Description string `json:"description"`
			CreatedBy   string `json:"createdBy"`
			CreatedAt   string `json:"createdAt"`
		} `json:"post"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAt, err := time.Parse(time.RFC3339, req.Post.CreatedAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid timestamp format"})
		return
	}

	post := &pb.Post{
		PostId:      req.Post.PostId,
		Title:       req.Post.Title,
		Description: req.Post.Description,
		CreatedBy:   req.Post.CreatedBy,
		CreatedAt:   timestamppb.New(createdAt),
		UpdatedAt:   timestamppb.Now(),
	}

	grpcReq := &pb.UpdatePostRequest{Post: post}

	res, err := c.PostService.UpdatePost(context.Background(), grpcReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *PostController) DeletePost(ctx *gin.Context) {
	postID := ctx.Param("id")
	req := &pb.DeletePostRequest{PostId: postID}

	_, err := c.PostService.DeletePost(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
