package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}

func Unauthorized(ctx *gin.Context, message string, errors interface{}) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"message": message,
		"errors":  errors,
	})
}

func NotFound(ctx *gin.Context, message string, errors interface{}) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": message,
		"errors":  errors,
	})
}

func BadRequest(ctx *gin.Context, message string, errors interface{}) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": message,
		"errors":  errors,
	})
}

func InternalServerError(ctx *gin.Context, message string, errors interface{}) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": message,
		"errors":  errors,
	})
}
