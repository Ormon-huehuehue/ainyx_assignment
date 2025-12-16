package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"go-backend-task/internal/logger"
	"go-backend-task/internal/middleware"
	"go-backend-task/internal/repository"
	"go-backend-task/internal/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize Logger
	logger.InitLogger()
	defer logger.Log.Sync()

	// Initialize Repository (Database)
	repo, err := repository.NewRepository()
	if err != nil {
		logger.Log.Fatal("Failed to initialize repository: " + err.Error())
	}
	defer repo.Close()

	// Initialize Fiber app
	app := fiber.New()

	// Middleware
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	// Setup Routes
	routes.SetupRoutes(app, repo)

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	logger.Log.Info("Starting server on port " + port)
	if err := app.Listen(":" + port); err != nil {
		logger.Log.Fatal("Server failed to start: " + err.Error())
	}
}
