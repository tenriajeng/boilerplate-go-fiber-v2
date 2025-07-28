package utils

import (
	"log"
	"os"

	"boilerplate-go-fiber-v2/config"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// InitializeApp initializes the application with all configurations
func InitializeApp() (*fiber.App, *config.Config, error) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Load configuration
	cfg := config.Load()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "Boilerplate Go Fiber v2",
		ErrorHandler: nil, // Will be set in main.go
	})

	return app, cfg, nil
}

// InitializeDatabase initializes database connection
func InitializeDatabase(cfg *config.Config) *gorm.DB {
	db := config.NewDatabase(cfg)
	return db
}

// InitializeRedis initializes Redis connection
func InitializeRedis(cfg *config.Config) *redis.Client {
	redis := config.NewRedis(cfg)
	return redis
}

// GetPort returns the port from environment or default
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

// GracefulShutdown handles graceful shutdown of the application
func GracefulShutdown(app *fiber.App) {
	// TODO: Implement graceful shutdown logic
	log.Println("Shutting down server...")
}
