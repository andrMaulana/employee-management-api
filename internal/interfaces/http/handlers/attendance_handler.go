package handlers

import (
	"strconv"
	"time"

	"github.com/andrMaulana/employee-management-api/internal/application/services"
	"github.com/andrMaulana/employee-management-api/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type AttendanceHandler struct {
	service services.AttendanceService
}

func NewAttendanceHandler(service services.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{service: service}
}

func (h *AttendanceHandler) CreateAttendance(c *fiber.Ctx) error {
	var input struct {
		EmployeeID uint       `json:"employee_id"`
		LocationID uint       `json:"location_id"`
		AbsentIn   *time.Time `json:"absent_in"`
		AbsentOut  *time.Time `json:"absent_out"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	// TODO: Get the user from the context after implementing authentication
	createdBy, ok := c.Locals("username").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User information not found in the context"})
	}

	att, err := h.service.CreateAttendance(c.Context(), input.EmployeeID, input.LocationID, input.AbsentIn, input.AbsentOut, createdBy)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(att)
}

func (h *AttendanceHandler) GetAllAttendances(c *fiber.Ctx) error {
	attendances, err := h.service.GetAllAttendances(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(attendances)
}

func (h *AttendanceHandler) GetAttendanceByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	att, err := h.service.GetAttendanceByID(c.Context(), uint(id))
	if err != nil {
		if err == errors.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(att)
}

func (h *AttendanceHandler) UpdateAttendance(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	var input struct {
		EmployeeID uint       `json:"employee_id"`
		LocationID uint       `json:"location_id"`
		AbsentIn   *time.Time `json:"absent_in"`
		AbsentOut  *time.Time `json:"absent_out"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	// TODO: Get the user from the context after implementing authentication
	updatedBy, ok := c.Locals("username").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User information not found in the context"})
	}

	att, err := h.service.UpdateAttendance(c.Context(), uint(id), input.EmployeeID, input.LocationID, input.AbsentIn, input.AbsentOut, updatedBy)
	if err != nil {
		if err == errors.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(att)
}

func (h *AttendanceHandler) DeleteAttendance(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.ErrInvalidInput.Error()})
	}

	if err := h.service.DeleteAttendance(c.Context(), uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *AttendanceHandler) GetAttendanceReport(c *fiber.Ctx) error {
	startDate, err := time.Parse("2006-01-02", c.Query("start_date"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start date format"})
	}

	endDate, err := time.Parse("2006-01-02", c.Query("end_date"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end date format"})
	}

	attendances, err := h.service.GetAttendanceReport(c.Context(), startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(attendances)
}
