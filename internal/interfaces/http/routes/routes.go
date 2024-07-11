package routes

import (
	"github.com/andrMaulana/employee-management-api/internal/interfaces/http/handlers"
	"github.com/andrMaulana/employee-management-api/internal/interfaces/http/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, departmentHandler *handlers.DepartmentHandler, locationHandler *handlers.LocationHandler, positionHandler *handlers.PositionHandler, employeeHandler *handlers.EmployeeHandler, attendanceHandler *handlers.AttendanceHandler, reportHandler *handlers.ReportHandler, authHandler *handlers.AuthHandler) {

	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)

	// Department routes
	departments := api.Group("/departments")
	departments.Use(middleware.AuthMiddleware())
	departments.Post("/", departmentHandler.CreateDepartment)
	departments.Get("/", departmentHandler.GetAllDepartments)
	departments.Get("/:id", departmentHandler.GetDepartmentByID)
	departments.Put("/:id", departmentHandler.UpdateDepartment)
	departments.Delete("/:id", departmentHandler.DeleteDepartment)

	// Position routes
	locations := api.Group("/locations")
	locations.Use(middleware.AuthMiddleware())
	locations.Post("/", locationHandler.CreateLocation)
	locations.Get("/", locationHandler.GetAllLocations)
	locations.Get("/:id", locationHandler.GetLocationByID)
	locations.Put("/:id", locationHandler.UpdateLocation)
	locations.Delete("/:id", locationHandler.DeleteLocation)

	// Employee routes
	employees := api.Group("/employees")
	// employees.Use(middleware.AuthMiddleware())
	employees.Post("/", employeeHandler.CreateEmployee)
	employees.Get("/", employeeHandler.GetAllEmployees)
	employees.Get("/:id", employeeHandler.GetEmployeeByID)
	employees.Put("/:id", employeeHandler.UpdateEmployee)
	employees.Delete("/:id", employeeHandler.DeleteEmployee)

	// Attendance routes
	attendances := api.Group("/attendances")
	attendances.Use(middleware.AuthMiddleware())
	attendances.Post("/", attendanceHandler.CreateAttendance)
	attendances.Get("/", attendanceHandler.GetAllAttendances)
	attendances.Get("/:id", attendanceHandler.GetAttendanceByID)
	attendances.Put("/:id", attendanceHandler.UpdateAttendance)
	attendances.Delete("/:id", attendanceHandler.DeleteAttendance)
	attendances.Get("/report", attendanceHandler.GetAttendanceReport)

	// Report routes
	reports := api.Group("/reports")
	reports.Get("/attendance", reportHandler.GetAttendanceReport)
}
