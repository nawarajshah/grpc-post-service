package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-api/service"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (a *AuthController) SignUp(ctx *gin.Context) {
	var req pb.SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := a.AuthService.SignUp(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (a *AuthController) SignIn(ctx *gin.Context) {
	var req pb.SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// Log and return the error if JSON binding fails
		fmt.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the request to ensure we're capturing the input
	fmt.Printf("AuthController received SignIn request with email: %s, password: %s\n", req.Email, req.Password)

	res, err := a.AuthService.SignIn(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
