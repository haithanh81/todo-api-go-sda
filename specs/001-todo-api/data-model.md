# Data Model: Todo API

**Branch**: `001-todo-api` | **Date**: 2025-11-22

## Overview

This document defines the data model for the Todo API, extracted from the feature specification requirements.

## Entities

### Task

Represents a todo item that a user wants to track.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | integer | Primary Key, Auto-increment | Unique identifier for the task (system-generated) |
| content | string | Required, Not empty, Max 1000 chars | Text description of what needs to be done |
| completed | boolean | Default: false | Whether the task has been finished |
| created_at | timestamp | Auto-generated | When the task was created |
| updated_at | timestamp | Auto-updated | When the task was last modified |

### Validation Rules

| Field | Rule | Error Message |
|-------|------|---------------|
| content | Required | "content is required" |
| content | Non-empty | "content cannot be empty" |
| content | Max length 1000 | "content exceeds maximum length of 1000 characters" |

### State Transitions

```
┌─────────────┐
│   Created   │
│ completed=  │
│   false     │
└──────┬──────┘
       │ Update (completed=true)
       ▼
┌─────────────┐
│  Completed  │
│ completed=  │
│   true      │
└──────┬──────┘
       │ Update (completed=false)
       ▼
┌─────────────┐
│ Incomplete  │
│ completed=  │
│   false     │
└─────────────┘
```

Tasks can freely transition between completed and incomplete states via update operations.

## Database Schema

### PostgreSQL Table Definition

```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    content VARCHAR(1000) NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Index for listing tasks (optional optimization)
CREATE INDEX idx_tasks_created_at ON tasks(created_at DESC);
```

### GORM Model

```go
package models

import (
    "time"
)

type Task struct {
    ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Content   string    `gorm:"type:varchar(1000);not null" json:"content"`
    Completed bool      `gorm:"default:false;not null" json:"completed"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Task) TableName() string {
    return "tasks"
}
```

## Request/Response DTOs

### CreateTaskRequest

```go
type CreateTaskRequest struct {
    Content string `json:"content" binding:"required,min=1,max=1000"`
}
```

### UpdateTaskRequest

```go
type UpdateTaskRequest struct {
    Content   *string `json:"content,omitempty" binding:"omitempty,min=1,max=1000"`
    Completed *bool   `json:"completed,omitempty"`
}
```

### TaskResponse

```go
type TaskResponse struct {
    ID        uint      `json:"id"`
    Content   string    `json:"content"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### TaskListResponse

```go
type TaskListResponse struct {
    Tasks []TaskResponse `json:"tasks"`
    Count int            `json:"count"`
}
```

### ErrorResponse

```go
type ErrorResponse struct {
    Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}
```

## Relationships

This is a simple single-entity model with no relationships. Future extensions might include:
- User entity (for multi-user support)
- Tag/Category entity (for task organization)
- Priority entity (for task prioritization)

## Notes

- `created_at` and `updated_at` are automatically managed by GORM
- The `id` field uses PostgreSQL's SERIAL type (auto-incrementing integer)
- Content validation happens at the API layer using Gin's binding validation
- Database constraints provide a second layer of validation
