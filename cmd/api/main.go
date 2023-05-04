package main

import (
	"github.com/your_project/config"
	"github.com/your_project/internal/api"
	"github.com/your_project/internal/api/controllers"
	"github.com/your_project/internal/api/services"
	"github.com/your_project/internal/models"
	"github.com/your_project/pkg/database"
	"github.com/your_project/pkg/jwt"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Initialize the database
	db, err := database.Initialize(cfg)
	if err != nil {
		panic(err)
	}

	// Initialize the UserRepository
	userRepo := models.NewUserRepository(db)

	// Initialize the TokenGenerator
	tokenGenerator := jwt.NewTokenGenerator(cfg.JWTSecret)

	// Create the AuthService
	authService := services.NewAuthService(userRepo, tokenGenerator)

	// Create the AuthController
	authController := controllers.NewAuthController(authService)

	// Create the server instance
	server := api.NewServer(cfg)

	// Register the AuthController routes
	server.RegisterAuthRoutes(authController)

	// Start the server
	server.Start()
}
