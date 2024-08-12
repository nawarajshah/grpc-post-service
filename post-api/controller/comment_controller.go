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
	postID := ctx.Param("postId")
	userID := ctx.GetHeader("X-User-ID") // Assume user ID is passed in a custom header

	var req pb.CreateCommentRequest
	if err := ctx.ShouldBindJSON(&req.Comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Comment.PostId = postID
	req.Comment.UserId = userID

	res, err := c.CommentService.CreateComment(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *CommentController) GetComment(ctx *gin.Context) {
	commentID := ctx.Param("commentId")

	req := &pb.GetCommentRequest{CommentId: commentID}

	res, err := c.CommentService.GetComment(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	commentID := ctx.Param("commentId")
	userID := ctx.GetHeader("X-User-ID") // Assume user ID is passed in a custom header

	var req pb.UpdateCommentRequest
	if err := ctx.ShouldBindJSON(&req.Comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Comment.CommentId = commentID
	req.Comment.UserId = userID

	res, err := c.CommentService.UpdateComment(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("commentId")
	userID := ctx.GetHeader("X-User-ID") // Assume user ID is passed in a custom header

	req := &pb.DeleteCommentRequest{CommentId: commentID, UserId: userID}

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
