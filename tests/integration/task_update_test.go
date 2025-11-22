package integration

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/todo-api-go-sda/internal/models"
)

func TestUpdateTask_UpdateContent(t *testing.T) {
	cleanupTasks(t)

	// Create a task
	createReq := models.CreateTaskRequest{Content: "Buy groceries"}
	createW := makeRequest(http.MethodPost, "/api/v1/tasks", createReq)
	var created models.TaskResponse
	parseResponse(t, createW, &created)

	// Update content
	newContent := "Buy organic groceries"
	updateReq := models.UpdateTaskRequest{Content: &newContent}
	w := makeRequest(http.MethodPut, fmt.Sprintf("/api/v1/tasks/%d", created.ID), updateReq)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.TaskResponse
	parseResponse(t, w, &response)

	assert.Equal(t, created.ID, response.ID)
	assert.Equal(t, "Buy organic groceries", response.Content)
	assert.False(t, response.Completed)
}

func TestUpdateTask_MarkCompleted(t *testing.T) {
	cleanupTasks(t)

	// Create a task
	createReq := models.CreateTaskRequest{Content: "Buy groceries"}
	createW := makeRequest(http.MethodPost, "/api/v1/tasks", createReq)
	var created models.TaskResponse
	parseResponse(t, createW, &created)

	// Mark as completed
	completed := true
	updateReq := models.UpdateTaskRequest{Completed: &completed}
	w := makeRequest(http.MethodPut, fmt.Sprintf("/api/v1/tasks/%d", created.ID), updateReq)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.TaskResponse
	parseResponse(t, w, &response)

	assert.Equal(t, created.ID, response.ID)
	assert.True(t, response.Completed)
}

func TestUpdateTask_UpdateBothFields(t *testing.T) {
	cleanupTasks(t)

	// Create a task
	createReq := models.CreateTaskRequest{Content: "Buy groceries"}
	createW := makeRequest(http.MethodPost, "/api/v1/tasks", createReq)
	var created models.TaskResponse
	parseResponse(t, createW, &created)

	// Update both fields
	newContent := "Buy organic groceries"
	completed := true
	updateReq := models.UpdateTaskRequest{Content: &newContent, Completed: &completed}
	w := makeRequest(http.MethodPut, fmt.Sprintf("/api/v1/tasks/%d", created.ID), updateReq)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.TaskResponse
	parseResponse(t, w, &response)

	assert.Equal(t, "Buy organic groceries", response.Content)
	assert.True(t, response.Completed)
}

func TestUpdateTask_NotFound(t *testing.T) {
	cleanupTasks(t)

	newContent := "Updated content"
	updateReq := models.UpdateTaskRequest{Content: &newContent}
	w := makeRequest(http.MethodPut, "/api/v1/tasks/999", updateReq)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response models.ErrorResponse
	parseResponse(t, w, &response)

	assert.Equal(t, "TASK_NOT_FOUND", response.Error.Code)
}
