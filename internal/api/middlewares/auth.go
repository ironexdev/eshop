package middlewares

import (
	"net/http"
	"strings"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type TokenParser interface {
	ParseToken(tokenString string) (*jwt.Token, error)
}

func AuthMiddleware(tokenParser TokenParser) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authorization header"})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid authorization header format"})
		}

		tokenString := tokenParts[1]
		token, err := tokenParser.ParseToken(tokenString)
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := uint(claims["user_id"].(float64))
			ctx.Locals("userID", userID)
			return ctx.Next()
		}

		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}
}
