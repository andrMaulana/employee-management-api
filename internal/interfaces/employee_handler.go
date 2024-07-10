package http

import (
	"github.com/andrMaulana/employee-management-api/internal/domain/employee"
	"github.com/andrMaulana/employee-management-api/pkg/auth"
	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	service employee.Service
}

func (h *EmployeeHandler) Login(c *fiber.Ctx) error {
	var loginRequest struct {
		Code     string `json:"code"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	employee, err := h.service.Authenticate(loginRequest.Code, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, err := auth.GenerateJWT(employee.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}
