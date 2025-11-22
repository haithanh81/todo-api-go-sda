//go:build integration

package integration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/todo-api-go-sda/internal/models"
)

func TestListTasks_WithTasks(t *testing.T) {
	cleanupTasks(t)

	// Create some tasks
	req1 := models.CreateTaskRequest{Content: "Task 1"}
	req2 := models.CreateTaskRequest{Content: "Task 2"}
	makeRequest(http.MethodPost, "/api/v1/tasks", req1)
	makeRequest(http.MethodPost, "/api/v1/tasks", req2)

	// List tasks
	w := makeRequest(http.MethodGet, "/api/v1/tasks", nil)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.TaskListResponse
	parseResponse(t, w, &response)

	assert.Equal(t, 2, response.Count)
	assert.Len(t, response.Tasks, 2)
}

func TestListTasks_EmptyList(t *testing.T) {
	cleanupTasks(t)

	w := makeRequest(http.MethodGet, "/api/v1/tasks", nil)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.TaskListResponse
	parseResponse(t, w, &response)

	assert.Equal(t, 0, response.Count)
	assert.Empty(t, response.Tasks)
}
