package models

import "time"

// CreateTaskRequest represents the request body for creating a task
type CreateTaskRequest struct {
	Content string `json:"content" binding:"required,min=1,max=1000"`
}

// UpdateTaskRequest represents the request body for updating a task
type UpdateTaskRequest struct {
	Content   *string `json:"content,omitempty" binding:"omitempty,min=1,max=1000"`
	Completed *bool   `json:"completed,omitempty"`
}

// TaskResponse represents a task in API responses
type TaskResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TaskListResponse represents a list of tasks in API responses
type TaskListResponse struct {
	Tasks []TaskResponse `json:"tasks"`
	Count int            `json:"count"`
}

// ErrorResponse represents an error in API responses
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail contains error details
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ToResponse converts a Task model to TaskResponse
func (t *Task) ToResponse() TaskResponse {
	return TaskResponse{
		ID:        t.ID,
		Content:   t.Content,
		Completed: t.Completed,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

// ToListResponse converts a slice of Tasks to TaskListResponse
func ToListResponse(tasks []Task) TaskListResponse {
	responses := make([]TaskResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = task.ToResponse()
	}
	return TaskListResponse{
		Tasks: responses,
		Count: len(responses),
	}
}
