package errors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error codes
const (
	CodeTaskNotFound    = "TASK_NOT_FOUND"
	CodeValidationError = "VALIDATION_ERROR"
	CodeInternalError   = "INTERNAL_ERROR"
)

// TaskNotFoundError represents a task not found error
type TaskNotFoundError struct {
	ID uint
}

func (e *TaskNotFoundError) Error() string {
	return fmt.Sprintf("Task with id %d not found", e.ID)
}

// ValidationError represents a validation error
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// ErrorResponse represents the error response structure
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail contains error details
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponse creates a new error response
func NewErrorResponse(code, message string) ErrorResponse {
	return ErrorResponse{
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
	}
}

// RespondWithError sends an error response to the client
func RespondWithError(c *gin.Context, statusCode int, code, message string) {
	c.JSON(statusCode, NewErrorResponse(code, message))
}

// HandleError handles different error types and responds appropriately
func HandleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *TaskNotFoundError:
		RespondWithError(c, http.StatusNotFound, CodeTaskNotFound, e.Error())
	case *ValidationError:
		RespondWithError(c, http.StatusBadRequest, CodeValidationError, e.Error())
	default:
		RespondWithError(c, http.StatusInternalServerError, CodeInternalError, "An internal error occurred")
	}
}
