package handlers

import (
	"strconv"

	"github.com/andrMaulana/employee-management-api/internal/application/services"
	"github.com/andrMaulana/employee-management-api/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type LocationHandler struct {
	service services.LocationService
}

func NewLocationHandler(service services.LocationService) *LocationHandler {
	return &LocationHandler{service: service}
}

func (h *LocationHandler) CreateLocation(c *fiber.Ctx) error {
	var input struct {
		LocationName string `json:"location_name"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	// TODO: Get the user from the context after implementing authentication
	createdBy, ok := c.Locals("username").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User information not found in the context"})
	}

	loc, err := h.service.CreateLocation(c.Context(), input.LocationName, createdBy)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(loc)
}

func (h *LocationHandler) GetAllLocations(c *fiber.Ctx) error {
	locations, err := h.service.GetAllLocations(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(locations)
}

func (h *LocationHandler) GetLocationByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	loc, err := h.service.GetLocationByID(c.Context(), uint(id))
	if err != nil {
		if err == errors.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(loc)
}

func (h *LocationHandler) UpdateLocation(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	var input struct {
		LocationName string `json:"location_name"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	// TODO: Get the user from the context after implementing authentication
	updatedBy, ok := c.Locals("username").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User information not found in the context"})
	}

	loc, err := h.service.UpdateLocation(c.Context(), uint(id), input.LocationName, updatedBy)
	if err != nil {
		if err == errors.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(loc)
}

func (h *LocationHandler) DeleteLocation(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	if err := h.service.DeleteLocation(c.Context(), uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
