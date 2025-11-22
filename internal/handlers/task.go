package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/todo-api-go-sda/internal/models"
	"github.com/todo-api-go-sda/internal/services"
	apperrors "github.com/todo-api-go-sda/pkg/errors"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	service services.TaskService
}

// NewTaskHandler creates a new TaskHandler instance
func NewTaskHandler(service services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// CreateTask handles POST /api/v1/tasks
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req models.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperrors.RespondWithError(c, http.StatusBadRequest, apperrors.CodeValidationError, "content is required")
		return
	}

	task, err := h.service.CreateTask(&req)
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, task.ToResponse())
}

// ListTasks handles GET /api/v1/tasks
func (h *TaskHandler) ListTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.ToListResponse(tasks))
}

// GetTask handles GET /api/v1/tasks/:id
func (h *TaskHandler) GetTask(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		apperrors.RespondWithError(c, http.StatusBadRequest, apperrors.CodeValidationError, "Invalid task ID")
		return
	}

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, task.ToResponse())
}

// UpdateTask handles PUT /api/v1/tasks/:id
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		apperrors.RespondWithError(c, http.StatusBadRequest, apperrors.CodeValidationError, "Invalid task ID")
		return
	}

	var req models.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperrors.RespondWithError(c, http.StatusBadRequest, apperrors.CodeValidationError, err.Error())
		return
	}

	task, err := h.service.UpdateTask(id, &req)
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, task.ToResponse())
}

// DeleteTask handles DELETE /api/v1/tasks/:id
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		apperrors.RespondWithError(c, http.StatusBadRequest, apperrors.CodeValidationError, "Invalid task ID")
		return
	}

	if err := h.service.DeleteTask(id); err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// parseID parses the ID from the URL parameter
func parseID(c *gin.Context) (uint, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
