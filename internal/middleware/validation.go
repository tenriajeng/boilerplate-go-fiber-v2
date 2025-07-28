package middleware

import (
	"boilerplate-go-fiber-v2/pkg/response"
	"boilerplate-go-fiber-v2/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

// ValidateRequest validates request body against a struct
func ValidateRequest(model interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse request body
		if err := c.BodyParser(model); err != nil {
			return response.ValidationError(c, "Invalid request body")
		}

		// Validate struct
		if err := validator.ValidateStruct(model); err != nil {
			return response.ValidationError(c, err.Error())
		}

		// Store validated model in context
		c.Locals("validated_model", model)

		return c.Next()
	}
}

// ValidateQuery validates query parameters
func ValidateQuery(model interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse query parameters
		if err := c.QueryParser(model); err != nil {
			return response.ValidationError(c, "Invalid query parameters")
		}

		// Validate struct
		if err := validator.ValidateStruct(model); err != nil {
			return response.ValidationError(c, err.Error())
		}

		// Store validated model in context
		c.Locals("validated_query", model)

		return c.Next()
	}
}

// ValidateParams validates URL parameters
func ValidateParams(model interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse URL parameters
		if err := c.ParamsParser(model); err != nil {
			return response.ValidationError(c, "Invalid URL parameters")
		}

		// Validate struct
		if err := validator.ValidateStruct(model); err != nil {
			return response.ValidationError(c, err.Error())
		}

		// Store validated model in context
		c.Locals("validated_params", model)

		return c.Next()
	}
}
