package repository

import (
	"github.com/todo-api-go-sda/internal/models"
	apperrors "github.com/todo-api-go-sda/pkg/errors"
	"gorm.io/gorm"
)

// TaskRepository defines the interface for task data access
type TaskRepository interface {
	Create(task *models.Task) error
	FindAll() ([]models.Task, error)
	FindByID(id uint) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
}

// taskRepository implements TaskRepository using GORM
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository creates a new TaskRepository instance
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

// Create creates a new task in the database
func (r *taskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

// FindAll retrieves all tasks from the database
func (r *taskRepository) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

// FindByID retrieves a task by its ID
func (r *taskRepository) FindByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &apperrors.TaskNotFoundError{ID: id}
		}
		return nil, err
	}
	return &task, nil
}

// Update updates an existing task in the database
func (r *taskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

// Delete removes a task from the database
func (r *taskRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return &apperrors.TaskNotFoundError{ID: id}
	}
	return nil
}
