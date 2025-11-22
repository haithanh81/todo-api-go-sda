//go:build integration

package integration

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/todo-api-go-sda/internal/models"
)

func TestGetTask_Success(t *testing.T) {
	cleanupTasks(t)

	// Create a task
	createReq := models.CreateTaskRequest{Content: "Buy groceries"}
	createW := makeRequest(http.MethodPost, "/api/v1/tasks", createReq)
	var created models.TaskResponse
	parseResponse(t, createW, &created)

	// Get the task
	w := makeRequest(http.MethodGet, fmt.Sprintf("/api/v1/tasks/%d", created.ID), nil)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.TaskResponse
	parseResponse(t, w, &response)

	assert.Equal(t, created.ID, response.ID)
	assert.Equal(t, "Buy groceries", response.Content)
	assert.False(t, response.Completed)
}

func TestGetTask_NotFound(t *testing.T) {
	cleanupTasks(t)

	w := makeRequest(http.MethodGet, "/api/v1/tasks/999", nil)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response models.ErrorResponse
	parseResponse(t, w, &response)

	assert.Equal(t, "TASK_NOT_FOUND", response.Error.Code)
}
