package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/redis/go-redis/v9"
)

type RateLimitMiddleware struct {
	redis *redis.Client
}

// NewRateLimitMiddleware creates a new rate limit middleware
func NewRateLimitMiddleware(redis *redis.Client) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		redis: redis,
	}
}

// GeneralRateLimit applies general rate limiting
func (m *RateLimitMiddleware) GeneralRateLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        100,             // Maximum 100 requests
		Expiration: 1 * time.Minute, // Per minute
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Rate limit by IP
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"error":   "Rate limit exceeded",
			})
		},
	})
}

// AuthRateLimit applies stricter rate limiting for auth endpoints
func (m *RateLimitMiddleware) AuthRateLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        5,               // Maximum 5 requests
		Expiration: 1 * time.Minute, // Per minute
		KeyGenerator: func(c *fiber.Ctx) string {
			return "auth:" + c.IP() // Rate limit by IP for auth
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"error":   "Too many authentication attempts",
			})
		},
	})
}

// UserRateLimit applies rate limiting per user
func (m *RateLimitMiddleware) UserRateLimit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		if userID == nil {
			return c.Next()
		}

		// Use user-specific rate limiting
		limiter := limiter.New(limiter.Config{
			Max:        50,              // Maximum 50 requests
			Expiration: 1 * time.Minute, // Per minute
			KeyGenerator: func(c *fiber.Ctx) string {
				return "user:" + userID.(string)
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
					"success": false,
					"error":   "User rate limit exceeded",
				})
			},
		})

		return limiter(c)
	}
}
