package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/todo-api-go-sda/internal/models"
	apperrors "github.com/todo-api-go-sda/pkg/errors"
)

// MockTaskRepository is a mock implementation of TaskRepository
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) FindAll() ([]models.Task, error) {
	args := m.Called()
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockTaskRepository) FindByID(id uint) (*models.Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskRepository) Update(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateTask_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	req := &models.CreateTaskRequest{Content: "Test task"}
	mockRepo.On("Create", mock.AnythingOfType("*models.Task")).Return(nil)

	task, err := service.CreateTask(req)

	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, "Test task", task.Content)
	assert.False(t, task.Completed)
	mockRepo.AssertExpectations(t)
}

func TestGetAllTasks_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	expectedTasks := []models.Task{
		{ID: 1, Content: "Task 1"},
		{ID: 2, Content: "Task 2"},
	}
	mockRepo.On("FindAll").Return(expectedTasks, nil)

	tasks, err := service.GetAllTasks()

	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	expectedTask := &models.Task{ID: 1, Content: "Task 1"}
	mockRepo.On("FindByID", uint(1)).Return(expectedTask, nil)

	task, err := service.GetTaskByID(1)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), task.ID)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID_NotFound(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	mockRepo.On("FindByID", uint(999)).Return(nil, &apperrors.TaskNotFoundError{ID: 999})

	task, err := service.GetTaskByID(999)

	assert.Error(t, err)
	assert.Nil(t, task)
	assert.IsType(t, &apperrors.TaskNotFoundError{}, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	existingTask := &models.Task{ID: 1, Content: "Old content", Completed: false}
	newContent := "New content"
	req := &models.UpdateTaskRequest{Content: &newContent}

	mockRepo.On("FindByID", uint(1)).Return(existingTask, nil)
	mockRepo.On("Update", mock.AnythingOfType("*models.Task")).Return(nil)

	task, err := service.UpdateTask(1, req)

	assert.NoError(t, err)
	assert.Equal(t, "New content", task.Content)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := service.DeleteTask(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
