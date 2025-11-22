package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/todo-api-go-sda/internal/models"
	apperrors "github.com/todo-api-go-sda/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		t.Fatalf("Failed to migrate: %v", err)
	}
	return db
}

func TestTaskRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)

	task := &models.Task{Content: "Test task"}
	err := repo.Create(task)

	assert.NoError(t, err)
	assert.NotZero(t, task.ID)
}

func TestTaskRepository_FindAll(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)

	// Create some tasks
	repo.Create(&models.Task{Content: "Task 1"})
	repo.Create(&models.Task{Content: "Task 2"})

	tasks, err := repo.FindAll()

	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
}

func TestTaskRepository_FindByID_Success(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)

	created := &models.Task{Content: "Test task"}
	repo.Create(created)

	task, err := repo.FindByID(created.ID)

	assert.NoError(t, err)
	assert.Equal(t, created.ID, task.ID)
	assert.Equal(t, "Test task", task.Content)
}

func TestTaskRepository_FindByID_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)

	task, err := repo.FindByID(999)

	assert.Error(t, err)
	assert.Nil(t, task)
	assert.IsType(t, &apperrors.TaskNotFoundError{}, err)
}

func TestTaskRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)

	task := &models.Task{Content: "Old content"}
	repo.Create(task)

	task.Content = "New content"
	err := repo.Update(task)

	assert.NoError(t, err)

	updated, _ := repo.FindByID(task.ID)
	assert.Equal(t, "New content", updated.Content)
}

func TestTaskRepository_Delete_Success(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)

	task := &models.Task{Content: "Task to delete"}
	repo.Create(task)

	err := repo.Delete(task.ID)

	assert.NoError(t, err)

	_, err = repo.FindByID(task.ID)
	assert.Error(t, err)
}

func TestTaskRepository_Delete_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)

	err := repo.Delete(999)

	assert.Error(t, err)
	assert.IsType(t, &apperrors.TaskNotFoundError{}, err)
}
