package integration

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/todo-api-go-sda/internal/models"
)

func TestDeleteTask_Success(t *testing.T) {
	cleanupTasks(t)

	// Create a task
	createReq := models.CreateTaskRequest{Content: "Task to delete"}
	createW := makeRequest(http.MethodPost, "/api/v1/tasks", createReq)
	var created models.TaskResponse
	parseResponse(t, createW, &created)

	// Delete the task
	w := makeRequest(http.MethodDelete, fmt.Sprintf("/api/v1/tasks/%d", created.ID), nil)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// Verify it's deleted by trying to get it
	getW := makeRequest(http.MethodGet, fmt.Sprintf("/api/v1/tasks/%d", created.ID), nil)
	assert.Equal(t, http.StatusNotFound, getW.Code)
}

func TestDeleteTask_NotFound(t *testing.T) {
	cleanupTasks(t)

	w := makeRequest(http.MethodDelete, "/api/v1/tasks/999", nil)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response models.ErrorResponse
	parseResponse(t, w, &response)

	assert.Equal(t, "TASK_NOT_FOUND", response.Error.Code)
}
