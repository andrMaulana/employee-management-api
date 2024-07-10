package middleware

import (
	"strings"

	"github.com/andrMaulana/employee-management-api/pkg/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authorization header"})
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
		}

		claims, err := auth.ValidateJWT(bearerToken[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		c.Locals("employeeID", claims.EmployeeID)
		return c.Next()
	}
}
