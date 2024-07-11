package main

import (
	"log"
	"os"

	"github.com/andrMaulana/employee-management-api/internal/application/services"
	"github.com/andrMaulana/employee-management-api/internal/domain/attendance"
	"github.com/andrMaulana/employee-management-api/internal/domain/department"
	"github.com/andrMaulana/employee-management-api/internal/domain/employee"
	"github.com/andrMaulana/employee-management-api/internal/domain/location"
	"github.com/andrMaulana/employee-management-api/internal/domain/position"
	"github.com/andrMaulana/employee-management-api/internal/infrastructure/database"
	"github.com/andrMaulana/employee-management-api/internal/interfaces/http/handlers"
	"github.com/andrMaulana/employee-management-api/internal/interfaces/http/routes"
	"github.com/andrMaulana/employee-management-api/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("configs/config.yaml")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	viper.AutomaticEnv() // Automatically use environment variables where available

	// Set default values from .env file if environment variables are not set
	viper.SetDefault("database.host", os.Getenv("DB_HOST"))
	viper.SetDefault("database.port", os.Getenv("DB_PORT"))
	viper.SetDefault("database.user", os.Getenv("DB_USER"))
	viper.SetDefault("database.password", os.Getenv("DB_PASSWORD"))
	viper.SetDefault("database.name", os.Getenv("DB_NAME"))
	viper.SetDefault("jwt.secret", os.Getenv("JWT_SECRET"))

	logger.Init()
}

func main() {
	app := fiber.New()

	db := database.NewPostgresDatabase()

	// Setup repositories
	departmentRepo := department.NewRepository(db)
	positionRepo := position.NewRepository(db)
	locationRepo := location.NewRepository(db)
	employeeRepo := employee.NewRepository(db)
	attendanceRepo := attendance.NewRepository(db)

	// Setup services
	departmentService := services.NewDepartmentService(departmentRepo)
	positionService := services.NewPositionService(positionRepo)
	locationService := services.NewLocationService(locationRepo)
	employeeService := services.NewEmployeeService(employeeRepo)
	attendanceService := services.NewAttendanceService(attendanceRepo)
	reportService := services.NewReportService(attendanceRepo, employeeRepo, departmentRepo, positionRepo, locationRepo)
	authService := services.NewAuthService(employeeRepo)

	// Setup handlers
	departmentHandler := handlers.NewDepartmentHandler(departmentService)
	positionHandler := handlers.NewPositionHandler(positionService)
	locationHandler := handlers.NewLocationHandler(locationService)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	attendanceHandler := handlers.NewAttendanceHandler(attendanceService)
	reportHandler := handlers.NewReportHandler(reportService)
	authHandler := handlers.NewAuthHandler(authService)

	// Setup routes
	routes.SetupRoutes(app, departmentHandler, positionHandler, locationHandler, employeeHandler, attendanceHandler, reportHandler, authHandler)

	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}

	logger.InfoLogger.Printf("Starting server on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		logger.ErrorLogger.Fatalf("Error starting server: %s", err)
	}
}
