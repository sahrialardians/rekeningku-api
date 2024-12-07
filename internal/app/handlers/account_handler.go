package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sahrialardians/rekeningku/internal/app/requests"
	"github.com/sahrialardians/rekeningku/internal/services"
	"github.com/sahrialardians/rekeningku/internal/utils"
)

type AccountHandler struct {
	accountService services.AccountService
}

// NewAccountHandler creates a new instance of AccountHandler
func NewAccountHandler(accountService services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

// GetAccounts for handle get all accounts
func (handler *AccountHandler) GetAccounts(ctx *gin.Context) {
	// get userId from Context
	ctxUserID, exists := ctx.Get("user_id")
	if !exists {
		utils.Unauthorized(ctx, "User not authenticated", nil)
		return
	}

	// convert userId to int
	userID, ok := ctxUserID.(int)
	if !ok {
		utils.InternalServerError(ctx, "Invalid user ID type", nil)
		return
	}

	// Get page number and page size from query parameters
	pageStr := ctx.Query("page")
	pageSizeStr := ctx.Query("page_size")

	// Default values
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	// Call service layer for paginated data
	accounts, totalRecords, err := handler.accountService.FindAll(userID, page, pageSize)
	if err != nil {
		utils.InternalServerError(ctx, "Failed to fetch accounts", err)
		return
	}

	// add pagination to response
	response := gin.H{
		"data": accounts,
		"pagination": gin.H{
			"current_page":  page,
			"page_size":     pageSize,
			"total_records": totalRecords,
			"total_pages":   (totalRecords + pageSize - 1) / pageSize,
		},
	}

	utils.Success(ctx, "Accounts retrieved successfully", response)
}

// GetAccount for handle get account by ID
func (handler *AccountHandler) GetAccount(ctx *gin.Context) { // Ambil user_id dari context
	ctxUserID, exists := ctx.Get("user_id")
	if !exists {
		utils.Unauthorized(ctx, "User not authenticated", nil)
		return
	}

	// convert userId to int
	userID, ok := ctxUserID.(int)
	if !ok {
		utils.InternalServerError(ctx, "Invalid user ID type", nil)
		return
	}

	accountID, err := strconv.Atoi(ctx.Param("accountId"))
	if err != nil {
		utils.BadRequest(ctx, "Invalid account ID", err)
		return
	}

	account, err := handler.accountService.FindById(userID, accountID)
	if err != nil {
		utils.NotFound(ctx, "Account not found", err)
		return
	}

	utils.Success(ctx, "Account retrieved successfully", account)
}

// CreateAccount for handle creating an account
func (handler *AccountHandler) CreateAccount(ctx *gin.Context) {
	var createAccountRequest requests.CreateAccountRequest

	if err := ctx.ShouldBindJSON(&createAccountRequest); err != nil {
		utils.BadRequest(ctx, "Invalid input", err)
		return
	}

	// get user id from context
	ctxUserID, exists := ctx.Get("user_id")
	if !exists {
		utils.Unauthorized(ctx, "User not authenticated", nil)
		return
	}

	// convert userId to int
	userID, ok := ctxUserID.(int)
	if !ok {
		utils.InternalServerError(ctx, "Invalid user ID type", nil)
		return
	}

	// assign userId to account request
	createAccountRequest.UserID = userID

	accountResponse, err := handler.accountService.Save(createAccountRequest)
	if err != nil {
		utils.InternalServerError(ctx, "Failed to create account", err)
		return
	}

	utils.Success(ctx, "Account created successfully", accountResponse)
}

// UpdateAccount for handle updating an account
func (handler *AccountHandler) UpdateAccount(ctx *gin.Context) {
	// get userId from context
	ctxUserID, exists := ctx.Get("user_id")
	if !exists {
		utils.Unauthorized(ctx, "User not authenticated", nil)
		return
	}

	// convert userId to int
	userID, ok := ctxUserID.(int)
	if !ok {
		utils.InternalServerError(ctx, "Invalid user ID type", nil)
		return
	}

	accountID, err := strconv.Atoi(ctx.Param("accountId"))
	if err != nil {
		utils.BadRequest(ctx, "Invalid account ID", err)
		return
	}

	var updateAccountRequest requests.UpdateAccountRequest
	if err := ctx.ShouldBindJSON(&updateAccountRequest); err != nil {
		utils.BadRequest(ctx, "Invalid input", err)
		return
	}

	// assign userId to struct account request
	updateAccountRequest.UserID = userID
	updateAccountRequest.ID = accountID

	err = handler.accountService.Update(updateAccountRequest)
	if err != nil {
		utils.InternalServerError(ctx, "Failed to update account", err)
		return
	}

	utils.Success(ctx, "Account updated successfully", nil)
}

// DeleteAccount for handle deleting an account
func (handler *AccountHandler) DeleteAccount(ctx *gin.Context) {
	ctxUserID, exists := ctx.Get("user_id")
	if !exists {
		utils.Unauthorized(ctx, "User not authenticated", nil)
		return
	}

	// Convert userId to int
	userID, ok := ctxUserID.(int)
	if !ok {
		utils.InternalServerError(ctx, "Invalid user ID type", nil)
		return
	}

	accountID, err := strconv.Atoi(ctx.Param("accountId"))
	if err != nil {
		utils.BadRequest(ctx, "Invalid account ID", err)
		return
	}

	err = handler.accountService.Delete(userID, accountID)
	if err != nil {
		utils.InternalServerError(ctx, "Failed to delete account", err)
		return
	}

	utils.Success(ctx, "Account deleted successfully", nil)
}
