package services

import (
	"github.com/todo-api-go-sda/internal/models"
	"github.com/todo-api-go-sda/internal/repository"
)

// TaskService defines the interface for task business logic
type TaskService interface {
	CreateTask(req *models.CreateTaskRequest) (*models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(id uint, req *models.UpdateTaskRequest) (*models.Task, error)
	DeleteTask(id uint) error
}

// taskService implements TaskService
type taskService struct {
	repo repository.TaskRepository
}

// NewTaskService creates a new TaskService instance
func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

// CreateTask creates a new task
func (s *taskService) CreateTask(req *models.CreateTaskRequest) (*models.Task, error) {
	task := &models.Task{
		Content:   req.Content,
		Completed: false,
	}
	err := s.repo.Create(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// GetAllTasks retrieves all tasks
func (s *taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.FindAll()
}

// GetTaskByID retrieves a task by its ID
func (s *taskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.repo.FindByID(id)
}

// UpdateTask updates an existing task
func (s *taskService) UpdateTask(id uint, req *models.UpdateTaskRequest) (*models.Task, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Content != nil {
		task.Content = *req.Content
	}
	if req.Completed != nil {
		task.Completed = *req.Completed
	}

	err = s.repo.Update(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// DeleteTask deletes a task by its ID
func (s *taskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
