# Quickstart Guide: Todo API

**Branch**: `001-todo-api` | **Date**: 2025-11-22

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose (for PostgreSQL)
- Make (optional, for convenience commands)

## Quick Setup

### 1. Start PostgreSQL

```bash
# Start PostgreSQL using Docker Compose
docker-compose up -d postgres
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Run the API

```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/tasks` | List all tasks |
| POST | `/api/v1/tasks` | Create a new task |
| GET | `/api/v1/tasks/{id}` | Get a task by ID |
| PUT | `/api/v1/tasks/{id}` | Update a task |
| DELETE | `/api/v1/tasks/{id}` | Delete a task |

## Example Requests

### Create a Task

```bash
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"content": "Buy groceries"}'
```

Response:
```json
{
  "id": 1,
  "content": "Buy groceries",
  "completed": false,
  "created_at": "2025-11-22T10:00:00Z",
  "updated_at": "2025-11-22T10:00:00Z"
}
```

### List All Tasks

```bash
curl http://localhost:8080/api/v1/tasks
```

Response:
```json
{
  "tasks": [
    {
      "id": 1,
      "content": "Buy groceries",
      "completed": false,
      "created_at": "2025-11-22T10:00:00Z",
      "updated_at": "2025-11-22T10:00:00Z"
    }
  ],
  "count": 1
}
```

### Get a Task

```bash
curl http://localhost:8080/api/v1/tasks/1
```

### Update a Task

```bash
curl -X PUT http://localhost:8080/api/v1/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"completed": true}'
```

### Delete a Task

```bash
curl -X DELETE http://localhost:8080/api/v1/tasks/1
```

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `DB_HOST` | `localhost` | PostgreSQL host |
| `DB_PORT` | `5432` | PostgreSQL port |
| `DB_USER` | `postgres` | PostgreSQL user |
| `DB_PASSWORD` | `postgres` | PostgreSQL password |
| `DB_NAME` | `todoapi` | PostgreSQL database name |

## Docker Deployment

### Build the Image

```bash
docker build -t todo-api .
```

### Run with Docker Compose

```bash
docker-compose up
```

This starts both the API and PostgreSQL database.

## Project Structure

```
.
├── cmd/api/main.go           # Application entry point
├── internal/
│   ├── config/               # Configuration management
│   ├── handlers/             # HTTP handlers (Gin)
│   ├── models/               # GORM models
│   ├── repository/           # Data access layer
│   └── services/             # Business logic
├── tests/
│   └── integration/          # API integration tests
├── openapi.yaml              # API specification
├── Dockerfile                # Container definition
├── docker-compose.yml        # Local development setup
└── go.mod                    # Go module definition
```

## Common Issues

### Database Connection Error

Ensure PostgreSQL is running:
```bash
docker-compose ps
docker-compose logs postgres
```

### Port Already in Use

Change the port via environment variable:
```bash
PORT=3000 go run cmd/api/main.go
```

## Next Steps

1. Run `/speckit.tasks` to generate implementation tasks
2. Follow the tasks in `specs/001-todo-api/tasks.md`
3. Implement, test, and iterate
