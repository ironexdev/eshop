package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement your authentication check logic here
		// If the user is authenticated, call c.Next() to proceed with the request
		// Otherwise, return an error or unauthorized status

		return c.Next()
	}
}
