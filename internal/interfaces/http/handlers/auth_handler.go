package handlers

import (
	"github.com/andrMaulana/employee-management-api/internal/application/services"
	"github.com/andrMaulana/employee-management-api/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input struct {
		EmployeeCode string `json:"employee_code"`
		Password     string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	token, err := h.service.Login(c.Context(), input.EmployeeCode, input.Password)
	if err != nil {
		if err == errors.ErrInvalidCredentials {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}
