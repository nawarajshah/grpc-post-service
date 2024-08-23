package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-api/service"
)

type CommentController struct {
	CommentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *CommentController {
	return &CommentController{
		CommentService: commentService,
	}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	var req pb.CreateCommentRequest
	req.PostId = ctx.Param("postId")
	req.UserId = ctx.Param("userId") // If using JWT, this might come from the token
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.CommentService.CreateComment(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *CommentController) GetComment(ctx *gin.Context) {
	postID := ctx.Param("postId")
	commentID := ctx.Param("commentId")
	req := &pb.GetCommentRequest{PostId: postID, CommentId: commentID}

	res, err := c.CommentService.GetComment(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	var req pb.UpdateCommentRequest
	req.CommentId = ctx.Param("commentId")
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.CommentService.UpdateComment(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	req := &pb.DeleteCommentRequest{
		CommentId: ctx.Param("commentId"),
	}

	_, err := c.CommentService.DeleteComment(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (c *CommentController) ListComments(ctx *gin.Context) {
	postID := ctx.Param("postId")

	req := &pb.GetCommentsByPostIDRequest{PostId: postID}

	res, err := c.CommentService.ListComments(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *CommentController) ApproveComment(ctx *gin.Context) {
	var req pb.ApproveCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.CommentService.ApproveComment(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
