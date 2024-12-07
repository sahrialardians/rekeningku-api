package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sahrialardians/rekeningku/internal/app/requests"
	"github.com/sahrialardians/rekeningku/internal/services"
	"github.com/sahrialardians/rekeningku/internal/utils"
)

type UserHandler struct {
	userService services.UserService
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// UpdateProfile handles updating user profiles
func (userHandler *UserHandler) UpdateProfile(ctx *gin.Context) {
	var updateUserRequest requests.UpdateUserRequest

	// Validate input
	if err := ctx.ShouldBindJSON(&updateUserRequest); err != nil {
		utils.BadRequest(ctx, "Invalid input", err)
		return
	}

	// Ambil user_id dari context
	userID, exists := ctx.Get("user_id")
	if !exists {
		utils.Unauthorized(ctx, "User not authenticated", nil)
		return
	}

	id, ok := userID.(int)
	if !ok {
		utils.InternalServerError(ctx, "Invalid user ID format", nil)
		return
	}

	updateUserRequest.ID = id

	// Call service layer
	if err := userHandler.userService.Update(updateUserRequest); err != nil {
		utils.InternalServerError(ctx, "Failed to update user profile", err)
		return
	}

	utils.Success(ctx, "Successfully updated user profile.", nil)
}

// GetUserProfile retrieves user profile by ID
func (userHandler *UserHandler) GetProfile(ctx *gin.Context) {

	// Ambil user_id dari context
	userID, exists := ctx.Get("user_id")
	if !exists {
		utils.Unauthorized(ctx, "User not authenticated", nil)
		return
	}

	id, ok := userID.(int)
	if !ok {
		utils.InternalServerError(ctx, "Invalid user ID format", nil)
		return
	}

	// Panggil service layer untuk mendapatkan data user
	userResponse, err := userHandler.userService.FindById(id)
	if err != nil {
		utils.InternalServerError(ctx, "Failed to retrieve user profile", err)
		return
	}

	// Return response ke client
	utils.Success(ctx, "User profile retrieved successfully", userResponse)
}
