package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahrialardians/rekeningku/internal/app/handlers"
	"github.com/sahrialardians/rekeningku/internal/app/middlewares"
	"github.com/sahrialardians/rekeningku/internal/utils"
)

func NewRouter(authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, accountHandler *handlers.AccountHandler) *gin.Engine {
	routes := gin.Default()

	routes.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Hello, Welcome to Rekeningku API.")
	})

	// Define when page not found
	routes.NoRoute(func(ctx *gin.Context) {
		utils.NotFound(ctx, "Page not found", nil)
	})

	router := routes.Group("/api/v1")

	// Auth routes
	router.POST("/auth/register", authHandler.Register)
	router.POST("/auth/login", authHandler.Login)

	// User routes
	router.Use(middlewares.Authenticated())
	{
		router.GET("/users", userHandler.GetProfile)
		router.PATCH("/users", userHandler.UpdateProfile)
	}

	return routes
}
