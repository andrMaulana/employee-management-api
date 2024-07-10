package http

import (
	"strconv"
	"time"

	"github.com/andrMaulana/employee-management-api/internal/domain/department"
	"github.com/andrMaulana/employee-management-api/pkg/pagination"
	"github.com/gofiber/fiber/v2"
)

type DepartmentHandler struct {
	service department.Service
}

func NewDepartmentHandler(service department.Service) *DepartmentHandler {
	return &DepartmentHandler{service: service}
}

func (h *DepartmentHandler) Create(c *fiber.Ctx) error {
	var dept department.Department
	if err := c.BodyParser(&dept); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.Create(c.Context(), &dept); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(dept)
}

func (h *DepartmentHandler) GetAll(c *fiber.Ctx) error {
	params := &department.GetAllParams{
		Paginator: pagination.Paginator{
			Page:  c.QueryInt("page", 1),
			Limit: c.QueryInt("limit", 10),
		},
		Search:    c.Query("search"),
		SortBy:    c.Query("sort_by", "Department_name"),
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

func (h *DepartmentHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	dept, err := h.service.GetByID(c.Context(), uint(id))
	if err != nil {
		if err == department.ErrDepartmentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Department not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(dept)
}

func (h *DepartmentHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var dept department.Department
	if err := c.BodyParser(&dept); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	dept.ID = uint(id)
	if err := h.service.Update(c.Context(), &dept); err != nil {
		if err == department.ErrDepartmentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Department not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(dept)
}

func (h *DepartmentHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.service.Delete(c.Context(), uint(id)); err != nil {
		if err == department.ErrDepartmentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Department not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *DepartmentHandler) BatchCreate(c *fiber.Ctx) error {
	var departments []department.Department
	if err := c.BodyParser(&departments); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.BatchCreate(c.Context(), departments); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(departments)
}

func (h *DepartmentHandler) BatchUpdate(c *fiber.Ctx) error {
	var departments []department.Department
	if err := c.BodyParser(&departments); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.BatchUpdate(c.Context(), departments); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(departments)
}

func (h *DepartmentHandler) BatchDelete(c *fiber.Ctx) error {
	var ids []uint
	if err := c.BodyParser(&ids); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.BatchDelete(c.Context(), ids); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
