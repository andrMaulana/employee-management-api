package handlers

import (
	"time"

	"github.com/andrMaulana/employee-management-api/internal/application/services"
	"github.com/gofiber/fiber/v2"
)

type ReportHandler struct {
	service services.ReportService
}

func NewReportHandler(service services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetAttendanceReport(c *fiber.Ctx) error {
	startDate, err := time.Parse("2006-01-02", c.Query("start_date"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start date format"})
	}

	endDate, err := time.Parse("2006-01-02", c.Query("end_date"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end date format"})
	}

	report, err := h.service.GetAttendanceReport(c.Context(), startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(report)
}
