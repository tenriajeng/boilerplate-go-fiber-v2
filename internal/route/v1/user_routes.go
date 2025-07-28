package v1

import (
	"boilerplate-go-fiber-v2/config"
	"boilerplate-go-fiber-v2/internal/container"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// SetupUserRoutes configures user-related routes
func SetupUserRoutes(router fiber.Router, container *container.Container, cfg *config.Config, redis *redis.Client) {
	// TODO: Implement user routes when UserHandler is created
	// user := router.Group("/users")
	// authMiddleware := middleware.NewAuthMiddleware(container.GetAuthService(), cfg)
	// protected := user.Group("/", authMiddleware.Authenticate())
}
