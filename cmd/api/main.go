package main

import (
	"log"

	"github.com/andrMaulana/employee-management-api/internal/domain/department"
	"github.com/andrMaulana/employee-management-api/internal/domain/position"
	"github.com/andrMaulana/employee-management-api/internal/infrastucture/database"
	"github.com/andrMaulana/employee-management-api/internal/interfaces/http"
	"github.com/andrMaulana/employee-management-api/internal/interfaces/http/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := database.NewPostgresDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// instace department
	departmentRepo := department.NewRepository(db)
	departmentService := department.NewService(departmentRepo)
	departmentHandler := http.NewDepartmentHandler(departmentService)

	// instace position
	positionRepo := position.NewRepository(db)
	positionService := position.NewService(positionRepo)
	positionHandler := http.NewPositionHandler(positionService)

	app := fiber.New()

	// Public routes
	// app.Post("/login", employeeHandler.Login)

	// Protected routes
	api := app.Group("/api/v1", middleware.AuthMiddleware())

	// Department routes
	app.Post("/departments", departmentHandler.Create)
	api.Get("/departments", departmentHandler.GetAll)
	api.Get("/departments/:id", departmentHandler.GetByID)
	app.Put("/departments/:id", departmentHandler.Update)
	api.Delete("/departments/:id", departmentHandler.Delete)

	api.Post("/departments/batch", departmentHandler.BatchCreate)
	api.Put("/departments/batch", departmentHandler.BatchUpdate)
	api.Delete("/departments/batch", departmentHandler.BatchDelete)

	// Position routes
	app.Post("/positions", positionHandler.Create)
	app.Get("/positions", positionHandler.GetAll)
	app.Get("/positions/:id", positionHandler.GetByID)
	app.Put("/positions/:id", positionHandler.Update)
	app.Delete("/positions/:id", positionHandler.Delete)

	log.Fatal(app.Listen(":8080"))
}
