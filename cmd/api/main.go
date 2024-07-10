package main

import (
	"log"

	"github.com/andrMaulana/employee-management-api/internal/domain/department"
	"github.com/andrMaulana/employee-management-api/internal/infrastucture/database"
	"github.com/andrMaulana/employee-management-api/internal/interfaces/http"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := database.NewPostgresDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	departmentRepo := department.NewRepository(db)
	departmentService := department.NewService(departmentRepo)
	departmentHandler := http.NewDepartmentHandler(departmentService)

	app := fiber.New()

	// v1 := app.Group("/api/v1", middleware.AuthMiddleware())

	// Department routes
	app.Post("/departments", departmentHandler.Create)
	app.Get("/departments", departmentHandler.GetAll)
	app.Get("/departments/:id", departmentHandler.GetByID)
	app.Put("/departments/:id", departmentHandler.Update)
	app.Delete("/departments/:id", departmentHandler.Delete)

	app.Post("/departments/batch", departmentHandler.BatchCreate)
	app.Put("/departments/batch", departmentHandler.BatchUpdate)
	app.Delete("/departments/batch", departmentHandler.BatchDelete)

	log.Fatal(app.Listen(":8080"))
}
