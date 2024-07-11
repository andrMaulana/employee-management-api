package handlers

import (
	"strconv"

	"github.com/andrMaulana/employee-management-api/internal/application/services"
	"github.com/andrMaulana/employee-management-api/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type PositionHandler struct {
	service services.PositionService
}

func NewPositionHandler(service services.PositionService) *PositionHandler {
	return &PositionHandler{service: service}
}

func (h *PositionHandler) CreatePosition(c *fiber.Ctx) error {
	var input struct {
		PositionName string `json:"position_name"`
		DepartmentID uint   `json:"department_id"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	// TODO: Get the user from the context after implementing authentication
	createdBy := "system"

	pos, err := h.service.CreatePosition(c.Context(), input.PositionName, input.DepartmentID, createdBy)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(pos)
}

func (h *PositionHandler) GetAllPositions(c *fiber.Ctx) error {
	positions, err := h.service.GetAllPositions(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(positions)
}

func (h *PositionHandler) GetPositionByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	pos, err := h.service.GetPositionByID(c.Context(), uint(id))
	if err != nil {
		if err == errors.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(pos)
}

func (h *PositionHandler) UpdatePosition(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	var input struct {
		PositionName string `json:"position_name"`
		DepartmentID uint   `json:"department_id"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	// TODO: Get the user from the context after implementing authentication
	updatedBy := "system"

	pos, err := h.service.UpdatePosition(c.Context(), uint(id), input.PositionName, input.DepartmentID, updatedBy)
	if err != nil {
		if err == errors.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(pos)
}

func (h *PositionHandler) DeletePosition(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	if err := h.service.DeletePosition(c.Context(), uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
