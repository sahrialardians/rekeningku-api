package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sahrialardians/rekeningku/internal/app/requests"
	"github.com/sahrialardians/rekeningku/internal/services"
	"github.com/sahrialardians/rekeningku/internal/utils"
)

type AuthHandler struct {
	authService services.AuthService
}

// NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (authHandler *AuthHandler) Register(ctx *gin.Context) {
	var registerRequest requests.RegisterUserRequest

	// validate input
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		utils.BadRequest(ctx, "Invalid input", []string{err.Error()})
		return
	}

	// Call service layer
	token, err := authHandler.authService.Register(registerRequest)
	if err != nil {
		utils.InternalServerError(ctx, "Failed to register user", []string{err.Error()})
		return
	}

	// Return success response with token
	utils.Success(ctx, "Succesfully registered user.", gin.H{"token": token})
}

func (authHandler *AuthHandler) Login(ctx *gin.Context) {
	var loginRequest requests.LoginUserRequest

	// validate input
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		utils.BadRequest(ctx, "Invalid input", err)
		return
	}

	// call service login
	token, err := authHandler.authService.Login(loginRequest)
	if err != nil {
		utils.Unauthorized(ctx, "Invalid credentials", err)
		return
	}

	// return token
	utils.Success(ctx, "Login successful", gin.H{"token": token})
}
