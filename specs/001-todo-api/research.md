# Research: Todo API

**Branch**: `001-todo-api` | **Date**: 2025-11-22

## Overview

This document captures research findings and decisions for implementing the Todo API. Technology stack is predetermined by the project constitution (Go, Gin, GORM, PostgreSQL).

## Research Tasks

### 1. Go Project Structure Best Practices

**Decision**: Standard Go project layout with `cmd/` and `internal/` directories

**Rationale**:
- `cmd/` contains entry points, following Go community conventions
- `internal/` prevents external imports of private packages
- Layered architecture (handlers → services → repository) promotes testability
- Separation of concerns makes the codebase maintainable

**Alternatives Considered**:
- Flat structure: Rejected - doesn't scale, harder to navigate
- Domain-driven design: Rejected - over-engineering for simple CRUD API
- Monolithic single package: Rejected - violates separation of concerns

### 2. Gin Framework Best Practices

**Decision**: Use Gin router groups, middleware for common concerns, and standard response patterns

**Rationale**:
- Router groups (`/api/v1/tasks`) enable API versioning
- Middleware handles cross-cutting concerns (logging, error recovery)
- Consistent response structure improves API usability
- Context-based request handling is idiomatic Gin

**Alternatives Considered**:
- Chi router: Rejected - constitution mandates Gin
- Standard library net/http: Rejected - constitution mandates Gin

**Best Practices Applied**:
- Use `gin.Context` for request/response handling
- Implement custom error middleware for consistent error responses
- Use binding tags for request validation
- Return appropriate HTTP status codes (201 Created, 200 OK, 404 Not Found, 400 Bad Request)

### 3. GORM with PostgreSQL Best Practices

**Decision**: Use GORM models with appropriate tags, repository pattern for data access

**Rationale**:
- GORM tags define schema (`gorm:"primaryKey;autoIncrement"`)
- Repository pattern abstracts database operations
- Connection pooling via GORM's built-in support
- AutoMigrate for development, manual migrations for production

**Alternatives Considered**:
- Raw SQL with database/sql: Rejected - constitution mandates GORM
- sqlx: Rejected - constitution mandates GORM

**Best Practices Applied**:
- Use `gorm:"not null"` for required fields
- Use `gorm:"default:false"` for boolean defaults
- Implement soft deletes if needed (using `gorm.DeletedAt`)
- Use transactions for multi-step operations (not needed for this simple API)

### 4. ID Generation Strategy

**Decision**: Use PostgreSQL auto-increment integers via GORM

**Rationale**:
- Simple and well-understood
- PostgreSQL SERIAL type handles generation
- GORM `autoIncrement` tag provides seamless integration
- Sufficient for moderate-scale application

**Alternatives Considered**:
- UUID: More complex, not required for simple API
- Snowflake IDs: Over-engineering for this use case
- Custom ID generation: Unnecessary complexity

### 5. Input Validation Strategy

**Decision**: Use Gin's binding validation with struct tags

**Rationale**:
- Declarative validation via struct tags (`binding:"required"`)
- Automatic request parsing and validation
- Clear error messages for invalid input
- Standard Go/Gin pattern

**Best Practices Applied**:
- Use `binding:"required"` for mandatory fields
- Use `binding:"max=1000"` for content length limit
- Return 400 Bad Request for validation failures
- Include specific error messages in response

### 6. Error Handling Strategy

**Decision**: Consistent error response format with appropriate HTTP status codes

**Rationale**:
- Standard error response structure aids API consumers
- HTTP status codes convey error category
- Detailed messages help debugging
- Consistent patterns across all endpoints

**Error Response Format**:
```json
{
  "error": {
    "code": "TASK_NOT_FOUND",
    "message": "Task with id 123 not found"
  }
}
```

**HTTP Status Code Mapping**:
- 400 Bad Request: Invalid input, validation failure
- 404 Not Found: Resource doesn't exist
- 500 Internal Server Error: Unexpected server errors

### 7. Testing Strategy

**Decision**: Unit tests for services/handlers, integration tests for API endpoints

**Rationale**:
- Unit tests verify business logic in isolation
- Integration tests verify API contract
- >80% coverage required by constitution
- Table-driven tests follow Go conventions

**Best Practices Applied**:
- Use `testing` package for unit tests
- Use `httptest` for handler testing
- Use test database for integration tests
- Mock repository for service unit tests

### 8. Docker Configuration

**Decision**: Multi-stage Dockerfile with Alpine base image

**Rationale**:
- Multi-stage builds reduce final image size
- Alpine provides minimal attack surface
- Consistent environment across development and production
- docker-compose for local development with PostgreSQL

## Resolved Clarifications

All technical decisions were predetermined by the project constitution. No additional clarifications were needed.

| Item | Resolution | Source |
|------|------------|--------|
| Language | Go 1.21+ | Constitution |
| Framework | Gin | Constitution |
| ORM | GORM | Constitution |
| Database | PostgreSQL | Constitution |
| API Spec | OpenAPI 3.0 | Constitution |
| Test Coverage | >80% | Constitution |
| Performance | <200ms P95 | Constitution |

## Dependencies

| Dependency | Version | Purpose |
|------------|---------|---------|
| github.com/gin-gonic/gin | v1.9+ | HTTP framework |
| gorm.io/gorm | v1.25+ | ORM |
| gorm.io/driver/postgres | v1.5+ | PostgreSQL driver |
| github.com/stretchr/testify | v1.8+ | Test assertions |

## Next Steps

1. Generate data model (`data-model.md`)
2. Generate API contracts (`contracts/openapi.yaml`)
3. Generate quickstart guide (`quickstart.md`)
