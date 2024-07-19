package handlers

import (
	"strconv"

	"github.com/andrMaulana/employee-management-api/internal/application/services"
	"github.com/andrMaulana/employee-management-api/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type DepartmentHandler struct {
	service services.DepartmentService
}

func NewDepartmentHandler(service services.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{service: service}
}

func (h *DepartmentHandler) CreateDepartment(c *fiber.Ctx) error {
	var input struct {
		DepartmentName string `json:"department_name"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	// TODO: Get the user from the context after implementing authentication
	createdBy, ok := c.Locals("username").(string) // Assuming the user information is a string
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User information not found in the context"})
	}

	dept, err := h.service.CreateDepartment(c.Context(), input.DepartmentName, createdBy)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(dept)
}

func (h *DepartmentHandler) GetAllDepartments(c *fiber.Ctx) error {
	departments, err := h.service.GetAllDepartments(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(departments)
}

func (h *DepartmentHandler) GetDepartmentByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	dept, err := h.service.GetDepartmentByID(c.Context(), uint(id))
	if err != nil {
		if err == errors.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(dept)
}

func (h *DepartmentHandler) UpdateDepartment(c *fiber.Ctx) error {
	// Periksa apakah user sudah terotentikasi
	if c.Locals("username") == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not authenticated"})
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	var input struct {
		DepartmentName string `json:"department_name"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	// TODO: Get the user from the context after implementing authentication
	updatedBy, ok := c.Locals("username").(string) // Assuming the user information is a string
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User information not found in the context"})
	}

	dept, err := h.service.UpdateDepartment(c.Context(), uint(id), input.DepartmentName, updatedBy)
	if err != nil {
		if err == errors.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(dept)
}

func (h *DepartmentHandler) DeleteDepartment(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	if err := h.service.DeleteDepartment(c.Context(), uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
