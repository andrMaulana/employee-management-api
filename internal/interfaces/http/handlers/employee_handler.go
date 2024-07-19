package handlers

import (
	"strconv"

	"github.com/andrMaulana/employee-management-api/internal/application/services"
	"github.com/andrMaulana/employee-management-api/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	service services.EmployeeService
}

func NewEmployeeHandler(service services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) CreateEmployee(c *fiber.Ctx) error {
	var input struct {
		EmployeeName string `json:"employee_name"`
		Password     string `json:"password"`
		DepartmentID uint   `json:"department_id"`
		PositionID   uint   `json:"position_id"`
		Superior     uint   `json:"superior"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	// TODO: Get the user from the context after implementing authentication
	createdBy, ok := c.Locals("username").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User information not found in the context"})
	}

	emp, err := h.service.CreateEmployee(c.Context(), input.EmployeeName, input.Password, input.DepartmentID, input.PositionID, input.Superior, createdBy)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(emp)
}

func (h *EmployeeHandler) GetAllEmployees(c *fiber.Ctx) error {
	employees, err := h.service.GetAllEmployees(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(employees)
}

func (h *EmployeeHandler) GetEmployeeByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	emp, err := h.service.GetEmployeeByID(c.Context(), uint(id))
	if err != nil {
		if err == errors.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(emp)
}

func (h *EmployeeHandler) UpdateEmployee(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	var input struct {
		EmployeeName string `json:"employee_name"`
		DepartmentID uint   `json:"department_id"`
		PositionID   uint   `json:"position_id"`
		Superior     uint   `json:"superior"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	// TODO: Get the user from the context after implementing authentication
	updatedBy, ok := c.Locals("username").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User information not found in the context"})
	}

	emp, err := h.service.UpdateEmployee(c.Context(), uint(id), input.EmployeeName, input.DepartmentID, input.PositionID, input.Superior, updatedBy)
	if err != nil {
		if err == errors.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(emp)
}

func (h *EmployeeHandler) DeleteEmployee(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	if err := h.service.DeleteEmployee(c.Context(), uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
