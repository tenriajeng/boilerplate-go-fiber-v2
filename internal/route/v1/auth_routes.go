package v1

import (
	"boilerplate-go-fiber-v2/config"
	"boilerplate-go-fiber-v2/internal/domain/service"
	"boilerplate-go-fiber-v2/internal/dto/auth"
	"boilerplate-go-fiber-v2/internal/handler"
	"boilerplate-go-fiber-v2/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// SetupAuthRoutes sets up authentication routes
func SetupAuthRoutes(router fiber.Router, authService service.AuthService, userService service.UserService, config *config.Config, redis *redis.Client) {
	// Create handlers
	authHandler := handler.NewAuthHandler(authService, userService)

	// Create middleware
	authMiddleware := middleware.NewAuthMiddleware(authService, config)

	// Create rate limit middleware only if Redis is available
	var rateLimitMiddleware *middleware.RateLimitMiddleware
	if redis != nil {
		rateLimitMiddleware = middleware.NewRateLimitMiddleware(redis)
	}

	// Auth group
	authGroup := router.Group("/auth")

	// Public auth routes (with rate limiting if Redis available)
	registerMiddleware := []fiber.Handler{
		middleware.ValidateRequest(&auth.RegisterRequest{}),
		authHandler.Register,
	}
	if rateLimitMiddleware != nil {
		registerMiddleware = append([]fiber.Handler{rateLimitMiddleware.AuthRateLimit()}, registerMiddleware...)
	}
	authGroup.Post("/register", registerMiddleware...)

	loginMiddleware := []fiber.Handler{
		middleware.ValidateRequest(&auth.LoginRequest{}),
		authHandler.Login,
	}
	if rateLimitMiddleware != nil {
		loginMiddleware = append([]fiber.Handler{rateLimitMiddleware.AuthRateLimit()}, loginMiddleware...)
	}
	authGroup.Post("/login", loginMiddleware...)

	authGroup.Post("/refresh",
		middleware.ValidateRequest(&auth.RefreshTokenRequest{}),
		authHandler.RefreshToken,
	)

	passwordResetMiddleware := []fiber.Handler{
		middleware.ValidateRequest(&auth.PasswordResetRequest{}),
		authHandler.CreatePasswordReset,
	}
	if rateLimitMiddleware != nil {
		passwordResetMiddleware = append([]fiber.Handler{rateLimitMiddleware.AuthRateLimit()}, passwordResetMiddleware...)
	}
	authGroup.Post("/password-reset", passwordResetMiddleware...)

	resetPasswordMiddleware := []fiber.Handler{
		middleware.ValidateRequest(&auth.ResetPasswordRequest{}),
		authHandler.ResetPassword,
	}
	if rateLimitMiddleware != nil {
		resetPasswordMiddleware = append([]fiber.Handler{rateLimitMiddleware.AuthRateLimit()}, resetPasswordMiddleware...)
	}
	authGroup.Post("/reset-password", resetPasswordMiddleware...)

	// Protected auth routes
	protected := authGroup.Group("/", authMiddleware.Authenticate())

	protected.Post("/logout", authHandler.Logout)

	protected.Post("/tfa/code", authHandler.CreateTFACode)

	protected.Post("/tfa/enable",
		middleware.ValidateRequest(&auth.EnableTFARequest{}),
		authHandler.EnableTFA,
	)

	protected.Post("/tfa/disable",
		middleware.ValidateRequest(&auth.DisableTFARequest{}),
		authHandler.DisableTFA,
	)

	protected.Post("/tfa/verify",
		middleware.ValidateRequest(&auth.TFACodeRequest{}),
		authHandler.VerifyTFA,
	)
}
