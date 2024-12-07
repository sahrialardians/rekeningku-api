package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sahrialardians/rekeningku/internal/app/handlers"
	"github.com/sahrialardians/rekeningku/internal/configs"
	"github.com/sahrialardians/rekeningku/internal/repositories"
	"github.com/sahrialardians/rekeningku/internal/routes"
	"github.com/sahrialardians/rekeningku/internal/services"
	"github.com/sahrialardians/rekeningku/internal/utils"
)

func main() {
	//Database
	db := configs.DatabaseConnection()

	fmt.Println("Migrating tables...")

	// Migrasi tabel
	// db.AutoMigrate(&models.User{}, &models.Account{})

	fmt.Println("Successfully migrated tables.")

	//Init Repository
	userRepo := repositories.NewUserRepositoryImpl(db)
	accountRepo := repositories.NewAccountRepositoryImpl(db)

	//Init Service
	userService := services.NewUserServiceImpl(userRepo)
	authService := services.NewAuthServiceImpl(userRepo)
	accountService := services.NewAccountServiceImpl(accountRepo)

	//Init controller
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	accountHandler := handlers.NewAccountHandler(accountService)

	//Router
	routes := routes.NewRouter(authHandler, userHandler, accountHandler)

	server := &http.Server{
		Addr:           ":9001",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	utils.ErrorPanic(err)
}
