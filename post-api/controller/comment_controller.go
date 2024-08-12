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
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postID := ctx.Param("postId")
	commentID := ctx.Param("commentId")

	req.Comment.PostId = postID
	req.Comment.CommentId = commentID

	res, err := c.CommentService.UpdateComment(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	postID := ctx.Param("postId")
	commentID := ctx.Param("commentId")
	userID := ctx.GetHeader("userId") // Assume userId is passed in the header

	req := &pb.DeleteCommentRequest{PostId: postID, CommentId: commentID, UserId: userID}

	_, err := c.CommentService.DeleteComment(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (c *CommentController) ListComments(ctx *gin.Context) {
	postID := ctx.Param("postId")
	req := &pb.ListCommentsRequest{PostId: postID}

	res, err := c.CommentService.ListComments(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
