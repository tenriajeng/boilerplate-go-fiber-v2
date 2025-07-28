package v1

import (
	"boilerplate-go-fiber-v2/config"
	"boilerplate-go-fiber-v2/internal/container"
	"boilerplate-go-fiber-v2/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// SetupAuthRoutes configures auth-related routes
func SetupAuthRoutes(router fiber.Router, container *container.Container, cfg *config.Config, redis *redis.Client) {
	auth := router.Group("/auth")

	// Public routes (no auth required)
	auth.Post("/register", container.GetAuthHandler().Register)
	auth.Post("/login", container.GetAuthHandler().Login)
	auth.Post("/refresh-token", container.GetAuthHandler().RefreshToken)
	auth.Post("/password-reset", container.GetAuthHandler().CreatePasswordReset)
	auth.Post("/reset-password", container.GetAuthHandler().ResetPassword)

	// Protected routes (auth required)
	authMiddleware := middleware.NewAuthMiddleware(container.GetAuthService(), cfg)
	protected := auth.Group("/", authMiddleware.Authenticate())
	protected.Post("/logout", container.GetAuthHandler().Logout)
	protected.Post("/tfa/create", container.GetAuthHandler().CreateTFACode)
	protected.Post("/tfa/enable", container.GetAuthHandler().EnableTFA)
	protected.Post("/tfa/disable", container.GetAuthHandler().DisableTFA)
	protected.Post("/tfa/verify", container.GetAuthHandler().VerifyTFA)
}
