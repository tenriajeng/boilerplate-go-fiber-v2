package route

import (
	"boilerplate-go-fiber-v2/config"
	"boilerplate-go-fiber-v2/internal/container"
	v1Routes "boilerplate-go-fiber-v2/internal/route/v1"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, db *gorm.DB, redis *redis.Client, cfg *config.Config) {
	// Health check endpoint
	app.Get("/health", healthCheck)

	// Initialize dependency container
	container := container.NewContainer(db, redis, cfg)
	log.Println("Dependency container initialized successfully")

	// API routes
	api := app.Group("/api")

	// v1 routes
	v1 := api.Group("/v1")
	setupV1Routes(v1, container, cfg, redis)

	// // v2 routes (future)
	// v2 := api.Group("/v2")
	// setupV2Routes(v2, container, cfg, redis)
}

// setupV1Routes configures v1 API routes
func setupV1Routes(router fiber.Router, container *container.Container, cfg *config.Config, redis *redis.Client) {
	log.Println("Setting up v1 routes...")

	// Setup v1 route modules
	v1Routes.SetupAuthRoutes(router, container, cfg, redis)
	v1Routes.SetupUserRoutes(router, container, cfg, redis)

	// v1 test endpoint
	router.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API v1 is working!",
			"version": "1.0.0",
			"status":  "stable",
		})
	})
}

// setupV2Routes configures v2 API routes (commented until needed)
// func setupV2Routes(router fiber.Router, container *container.Container, cfg *config.Config, redis *redis.Client) {
// 	log.Println("Setting up v2 routes...")
//
// 	// TODO: Add v2 specific routes when needed
// 	// v2Routes.SetupAuthRoutes(router, container, cfg, redis)
// 	// v2Routes.SetupUserRoutes(router, container, cfg, redis)
//
// 	// v2 test endpoint
// 	router.Get("/test", func(c *fiber.Ctx) error {
// 		return c.JSON(fiber.Map{
// 			"message": "API v2 is working!",
// 			"version": "2.0.0",
// 			"status":  "beta",
// 		})
// 	})
// }

// healthCheck handles the health check endpoint
func healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Server is running",
		"version": "1.0.0",
	})
}
