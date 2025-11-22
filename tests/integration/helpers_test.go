package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/todo-api-go-sda/internal/handlers"
	"github.com/todo-api-go-sda/internal/models"
	"github.com/todo-api-go-sda/internal/repository"
	"github.com/todo-api-go-sda/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	testDB     *gorm.DB
	testRouter *gin.Engine
)

// TestMain sets up and tears down the test environment
func TestMain(m *testing.M) {
	// Setup
	setup()

	// Run tests
	code := m.Run()

	// Teardown
	teardown()

	os.Exit(code)
}

func setup() {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Connect to test database
	dsn := getTestDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database: " + err.Error())
	}
	testDB = db

	// Migrate schema
	if err := testDB.AutoMigrate(&models.Task{}); err != nil {
		panic("Failed to migrate test database: " + err.Error())
	}

	// Setup router
	testRouter = setupTestRouter(testDB)
}

func teardown() {
	// Clean up database
	if testDB != nil {
		sqlDB, _ := testDB.DB()
		sqlDB.Close()
	}
}

func getTestDSN() string {
	host := getEnvOrDefault("TEST_DB_HOST", "localhost")
	port := getEnvOrDefault("TEST_DB_PORT", "5432")
	user := getEnvOrDefault("TEST_DB_USER", "postgres")
	password := getEnvOrDefault("TEST_DB_PASSWORD", "postgres")
	dbname := getEnvOrDefault("TEST_DB_NAME", "todoapi_test")

	return "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func setupTestRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()

	// Initialize dependencies
	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Setup routes
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

	return router
}

// cleanupTasks removes all tasks from the test database
func cleanupTasks(t *testing.T) {
	t.Helper()
	if err := testDB.Exec("DELETE FROM tasks").Error; err != nil {
		t.Fatalf("Failed to cleanup tasks: %v", err)
	}
}

// makeRequest is a helper to make HTTP requests to the test router
func makeRequest(method, path string, body interface{}) *httptest.ResponseRecorder {
	var reqBody *bytes.Buffer
	if body != nil {
		jsonBody, _ := json.Marshal(body)
		reqBody = bytes.NewBuffer(jsonBody)
	} else {
		reqBody = bytes.NewBuffer(nil)
	}

	req, _ := http.NewRequest(method, path, reqBody)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	return w
}

// parseResponse parses a JSON response into the given target
func parseResponse(t *testing.T, w *httptest.ResponseRecorder, target interface{}) {
	t.Helper()
	if err := json.Unmarshal(w.Body.Bytes(), target); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}
}
