package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/todo-api-go-sda/internal/config"
	"github.com/todo-api-go-sda/internal/handlers"
	"github.com/todo-api-go-sda/internal/models"
	"github.com/todo-api-go-sda/internal/repository"
	"github.com/todo-api-go-sda/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize dependencies
	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Setup Gin router
	router := gin.Default()

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.POST("", taskHandler.CreateTask)
			tasks.GET("", taskHandler.ListTasks)
			tasks.GET("/:id", taskHandler.GetTask)
			tasks.PUT("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}
	}

	// Start server
	log.Printf("Starting server on port %s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
