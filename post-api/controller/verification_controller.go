package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-api/service"
)

type VerificationController struct {
	VerificationService service.VerificationService
}

func NewVerificationController(verificationService service.VerificationService) *VerificationController {
	return &VerificationController{
		VerificationService: verificationService,
	}
}

func (c *VerificationController) VerifyEmail(ctx *gin.Context) {
	var req pb.VerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.VerificationService.VerifyEmail(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
