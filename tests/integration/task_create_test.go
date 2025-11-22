package integration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/todo-api-go-sda/internal/models"
)

func TestCreateTask_Success(t *testing.T) {
	cleanupTasks(t)

	req := models.CreateTaskRequest{
		Content: "Buy groceries",
	}

	w := makeRequest(http.MethodPost, "/api/v1/tasks", req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.TaskResponse
	parseResponse(t, w, &response)

	assert.NotZero(t, response.ID)
	assert.Equal(t, "Buy groceries", response.Content)
	assert.False(t, response.Completed)
	assert.NotZero(t, response.CreatedAt)
	assert.NotZero(t, response.UpdatedAt)
}

func TestCreateTask_MissingContent(t *testing.T) {
	cleanupTasks(t)

	req := map[string]interface{}{}

	w := makeRequest(http.MethodPost, "/api/v1/tasks", req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response models.ErrorResponse
	parseResponse(t, w, &response)

	assert.Equal(t, "VALIDATION_ERROR", response.Error.Code)
	assert.Contains(t, response.Error.Message, "content")
}

func TestCreateTask_EmptyContent(t *testing.T) {
	cleanupTasks(t)

	req := models.CreateTaskRequest{
		Content: "",
	}

	w := makeRequest(http.MethodPost, "/api/v1/tasks", req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response models.ErrorResponse
	parseResponse(t, w, &response)

	assert.Equal(t, "VALIDATION_ERROR", response.Error.Code)
}
