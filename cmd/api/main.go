package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahrialardians/rekeningku/internal/utils"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Hello, this is Rekeningku API.")
	})

	server := &http.Server{
		Addr:    ":9001",
		Handler: ,
	}

	err := server.ListenAndServe()
	utils.ErrorPanic(err)
}
