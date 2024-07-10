package http

import (
	"strconv"
	"time"

	"github.com/andrMaulana/employee-management-api/internal/domain/position"
	"github.com/andrMaulana/employee-management-api/pkg/pagination"
	"github.com/gofiber/fiber/v2"
)

type PositionHandler struct {
	service position.Service
}

func NewPositionHandler(service position.Service) *PositionHandler {
	return &PositionHandler{service}
}

func (h *PositionHandler) Create(c *fiber.Ctx) error {

	var post position.Position
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.Create(c.Context(), &post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(post)

}

func (h *PositionHandler) GetAll(c *fiber.Ctx) error {
	params := &position.GetAllParams{
		Paginator: pagination.Paginator{
			Page:  c.QueryInt("page", 1),
			Limit: c.QueryInt("limit", 10),
		},
		Search:    c.Query("search"),
		SortBy:    c.Query("sort_by", "Position_name"),
		SortOrder: c.Query("sort_order", "asc"),
	}

	if createdAt := c.Query("created_at"); createdAt != "" {
		t, err := time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid created_at format"})
		}
		params.CreatedAt = &t
	}

	if updatedAt := c.Query("updated_at"); updatedAt != "" {
		t, err := time.Parse(time.RFC3339, updatedAt)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid updated_at format"})
		}
		params.UpdatedAt = &t
	}

	result, err := h.service.GetAll(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

func (h *PositionHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	dept, err := h.service.GetByID(c.Context(), uint(id))
	if err != nil {
		if err == position.ErrPositionNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Position not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(dept)
}

func (h *PositionHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var dept position.Position
	if err := c.BodyParser(&dept); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	dept.ID = uint(id)
	if err := h.service.Update(c.Context(), &dept); err != nil {
		if err == position.ErrPositionNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Position not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(dept)
}

func (h *PositionHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.service.Delete(c.Context(), uint(id)); err != nil {
		if err == position.ErrPositionNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Position not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
