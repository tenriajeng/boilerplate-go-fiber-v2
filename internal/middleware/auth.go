package middleware

import (
	"strings"

	"boilerplate-go-fiber-v2/config"
	"boilerplate-go-fiber-v2/internal/domain/service"
	"boilerplate-go-fiber-v2/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	authService service.AuthService
	config      *config.Config
}

// NewAuthMiddleware creates a new auth middleware
func NewAuthMiddleware(authService service.AuthService, config *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
		config:      config,
	}
}

// Authenticate validates JWT token and sets user context
func (m *AuthMiddleware) Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Unauthorized(c, "Authorization header required")
		}

		// Check Bearer token format
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return response.Unauthorized(c, "Invalid token format")
		}

		// Extract token
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate token
		claims, err := m.authService.ValidateToken(c.Context(), token)
		if err != nil {
			return response.Unauthorized(c, "Invalid or expired token")
		}

		// Set user context
		c.Locals("user_id", claims.UserID)
		c.Locals("user_email", claims.Email)
		c.Locals("user_role", claims.Role)

		return c.Next()
	}
}

// RequireRole checks if user has required role
func (m *AuthMiddleware) RequireRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("user_role").(string)
		if userRole != role {
			return response.Forbidden(c, "Insufficient permissions")
		}
		return c.Next()
	}
}

// RequireRoles checks if user has any of the required roles
func (m *AuthMiddleware) RequireRoles(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("user_role").(string)
		for _, role := range roles {
			if userRole == role {
				return c.Next()
			}
		}
		return response.Forbidden(c, "Insufficient permissions")
	}
}

// OptionalAuth validates token if present but doesn't require it
func (m *AuthMiddleware) OptionalAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Next()
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Next()
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := m.authService.ValidateToken(c.Context(), token)
		if err != nil {
			return c.Next()
		}

		// Set user context if valid
		c.Locals("user_id", claims.UserID)
		c.Locals("user_email", claims.Email)
		c.Locals("user_role", claims.Role)

		return c.Next()
	}
}
