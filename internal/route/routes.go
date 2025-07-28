package route

import (
	"boilerplate-go-fiber-v2/config"
	"boilerplate-go-fiber-v2/internal/domain/repository"
	"boilerplate-go-fiber-v2/internal/domain/service"
	repo "boilerplate-go-fiber-v2/internal/repository"
	v1 "boilerplate-go-fiber-v2/internal/route/v1"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, db *gorm.DB, redis *redis.Client, cfg *config.Config) {
	// Health check endpoint
	app.Get("/health", healthCheck)

	// API routes
	api := app.Group("/api")

	// v1 routes
	v1 := api.Group("/v1")
	setupV1Routes(v1, db, redis, cfg)

	// v2 routes (for future use)
	v2 := api.Group("/v2")
	setupV2Routes(v2, db, redis, cfg)
}

// setupV1Routes configures v1 API routes
func setupV1Routes(router fiber.Router, db *gorm.DB, redis *redis.Client, cfg *config.Config) {
	log.Println("Setting up v1 routes...")

	// Initialize repositories
	var userRepo repository.UserRepository
	var authRepo repository.AuthRepository

	if db != nil {
		log.Println("Database is not nil, initializing repositories...")
		userRepo = repo.NewUserRepository(db)
		authRepo = repo.NewAuthRepository(db)
		log.Println("Repositories initialized successfully")
	} else {
		log.Println("WARNING: Database is nil!")
	}

	// Initialize services
	var userService service.UserService
	var authService service.AuthService

	if userRepo != nil {
		log.Println("UserRepo is not nil, initializing services...")
		userService = service.NewUserService(userRepo)
		authService = service.NewAuthService(userRepo, authRepo, userService, cfg)
		log.Println("Services initialized successfully")
	} else {
		log.Println("WARNING: UserRepo is nil!")
	}

	// Setup v1 routes
	v1.SetupAuthRoutes(router, authService, userService, cfg, redis)

	// Temporary test endpoint
	router.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API v1 is working!",
		})
	})
}

// setupV2Routes configures v2 API routes
func setupV2Routes(router fiber.Router, db *gorm.DB, redis *redis.Client, cfg *config.Config) {
	// TODO: Initialize handlers and setup v2 routes
	// This will be implemented in the next steps

	// Temporary test endpoint
	router.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API v2 is working!",
		})
	})
}

// healthCheck handles the health check endpoint
func healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Server is running",
		"version": "1.0.0",
	})
}
