package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/todo-api-go-sda/internal/models"
	apperrors "github.com/todo-api-go-sda/pkg/errors"
)

// MockTaskService is a mock implementation of TaskService
type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) CreateTask(req *models.CreateTaskRequest) (*models.Task, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskService) GetAllTasks() ([]models.Task, error) {
	args := m.Called()
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockTaskService) GetTaskByID(id uint) (*models.Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskService) UpdateTask(id uint, req *models.UpdateTaskRequest) (*models.Task, error) {
	args := m.Called(id, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskService) DeleteTask(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupTestRouter(handler *TaskHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/api/v1")
	tasks := v1.Group("/tasks")
	tasks.POST("", handler.CreateTask)
	tasks.GET("", handler.ListTasks)
	tasks.GET("/:id", handler.GetTask)
	tasks.PUT("/:id", handler.UpdateTask)
	tasks.DELETE("/:id", handler.DeleteTask)
	return router
}

func TestCreateTask_Success(t *testing.T) {
	mockService := new(MockTaskService)
	handler := NewTaskHandler(mockService)
	router := setupTestRouter(handler)

	task := &models.Task{ID: 1, Content: "Test task", Completed: false}
	mockService.On("CreateTask", mock.AnythingOfType("*models.CreateTaskRequest")).Return(task, nil)

	body, _ := json.Marshal(models.CreateTaskRequest{Content: "Test task"})
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestCreateTask_ValidationError(t *testing.T) {
	mockService := new(MockTaskService)
	handler := NewTaskHandler(mockService)
	router := setupTestRouter(handler)

	body, _ := json.Marshal(map[string]interface{}{})
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListTasks_Success(t *testing.T) {
	mockService := new(MockTaskService)
	handler := NewTaskHandler(mockService)
	router := setupTestRouter(handler)

	tasks := []models.Task{{ID: 1, Content: "Task 1"}}
	mockService.On("GetAllTasks").Return(tasks, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/tasks", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetTask_Success(t *testing.T) {
	mockService := new(MockTaskService)
	handler := NewTaskHandler(mockService)
	router := setupTestRouter(handler)

	task := &models.Task{ID: 1, Content: "Test task"}
	mockService.On("GetTaskByID", uint(1)).Return(task, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/tasks/1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetTask_NotFound(t *testing.T) {
	mockService := new(MockTaskService)
	handler := NewTaskHandler(mockService)
	router := setupTestRouter(handler)

	mockService.On("GetTaskByID", uint(999)).Return(nil, &apperrors.TaskNotFoundError{ID: 999})

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/tasks/999", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteTask_Success(t *testing.T) {
	mockService := new(MockTaskService)
	handler := NewTaskHandler(mockService)
	router := setupTestRouter(handler)

	mockService.On("DeleteTask", uint(1)).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/tasks/1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}
