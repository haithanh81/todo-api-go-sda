# Tasks: Todo API

**Input**: Design documents from `/specs/001-todo-api/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: Constitution requires >80% test coverage - test tasks are included in Foundational phase.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- **Project type**: Single API service (Go)
- **Source**: `cmd/`, `internal/`, `pkg/`
- **Tests**: `tests/integration/`, unit tests alongside source

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 Initialize Go module with `go mod init` in project root (go.mod)
- [x] T002 Create project directory structure per plan: cmd/api/, internal/config/, internal/handlers/, internal/models/, internal/repository/, internal/services/, pkg/errors/, tests/integration/
- [x] T003 [P] Add Gin dependency `github.com/gin-gonic/gin` to go.mod
- [x] T004 [P] Add GORM dependencies `gorm.io/gorm` and `gorm.io/driver/postgres` to go.mod
- [x] T005 [P] Add testify dependency `github.com/stretchr/testify` to go.mod
- [x] T006 [P] Copy OpenAPI specification from specs/001-todo-api/contracts/openapi.yaml to openapi.yaml at project root

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**CRITICAL**: No user story work can begin until this phase is complete

- [x] T007 Create configuration management in internal/config/config.go (database connection, server port from environment variables)
- [x] T008 Create Task GORM model in internal/models/task.go with id, content, completed, created_at, updated_at fields per data-model.md
- [x] T009 Create request/response DTOs in internal/models/dto.go (CreateTaskRequest, UpdateTaskRequest, TaskResponse, TaskListResponse, ErrorResponse)
- [x] T010 Create custom error types in pkg/errors/errors.go (TaskNotFoundError, ValidationError, error response helper)
- [x] T011 Create TaskRepository interface and implementation in internal/repository/task.go (Create, FindAll, FindByID, Update, Delete methods)
- [x] T012 Create TaskService interface and implementation in internal/services/task.go (business logic layer calling repository)
- [x] T013 Create Gin router setup with /api/v1 group and middleware (logging, recovery) in cmd/api/main.go
- [x] T014 Create database connection and AutoMigrate in cmd/api/main.go
- [x] T015 [P] Create Dockerfile with multi-stage build (build stage with Go, runtime stage with Alpine)
- [x] T016 [P] Create docker-compose.yml with PostgreSQL service and API service
- [x] T017 [P] Create test helper utilities in tests/integration/helpers_test.go (test database setup, cleanup, HTTP client)

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Create a New Task (Priority: P1)

**Goal**: API consumers can create a new task with content, receiving a task with auto-generated id and completed=false

**Independent Test**: Send POST /api/v1/tasks with content, verify response has id, content, and completed=false

### Implementation for User Story 1

- [x] T018 [US1] Implement CreateTask handler in internal/handlers/task.go (bind CreateTaskRequest, call service, return 201 with TaskResponse)
- [x] T019 [US1] Add validation error handling for missing/empty content in CreateTask handler (return 400 with ErrorResponse)
- [x] T020 [US1] Register POST /api/v1/tasks route in cmd/api/main.go
- [x] T021 [US1] Create integration test for successful task creation in tests/integration/task_create_test.go
- [x] T022 [US1] Create integration test for validation errors (missing content, empty content) in tests/integration/task_create_test.go

**Checkpoint**: User Story 1 (Create Task) is fully functional and testable independently

---

## Phase 4: User Story 2 - List All Tasks (Priority: P1)

**Goal**: API consumers can retrieve all tasks to see their complete todo list

**Independent Test**: Create tasks, then GET /api/v1/tasks, verify all tasks returned with count

### Implementation for User Story 2

- [x] T023 [US2] Implement ListTasks handler in internal/handlers/task.go (call service, return 200 with TaskListResponse)
- [x] T024 [US2] Register GET /api/v1/tasks route in cmd/api/main.go
- [x] T025 [US2] Create integration test for listing tasks (with tasks and empty list) in tests/integration/task_list_test.go

**Checkpoint**: User Stories 1 AND 2 form the MVP - create and list tasks both work independently

---

## Phase 5: User Story 3 - Update a Task (Priority: P2)

**Goal**: API consumers can update content and/or completed status of an existing task

**Independent Test**: Create task, update content/completed via PUT, verify changes persisted

### Implementation for User Story 3

- [x] T026 [US3] Implement UpdateTask handler in internal/handlers/task.go (parse id, bind UpdateTaskRequest, call service, return 200 with TaskResponse)
- [x] T027 [US3] Add not found error handling in UpdateTask handler (return 404 with ErrorResponse)
- [x] T028 [US3] Register PUT /api/v1/tasks/:id route in cmd/api/main.go
- [x] T029 [US3] Create integration test for successful task update (content, completed, both) in tests/integration/task_update_test.go
- [x] T030 [US3] Create integration test for update not found error in tests/integration/task_update_test.go

**Checkpoint**: User Story 3 (Update Task) is fully functional - tasks can be modified and marked complete

---

## Phase 6: User Story 4 - Delete a Task (Priority: P2)

**Goal**: API consumers can delete a task to remove it from the todo list

**Independent Test**: Create task, delete via DELETE, verify no longer in list

### Implementation for User Story 4

- [x] T031 [US4] Implement DeleteTask handler in internal/handlers/task.go (parse id, call service, return 204 no content)
- [x] T032 [US4] Add not found error handling in DeleteTask handler (return 404 with ErrorResponse)
- [x] T033 [US4] Register DELETE /api/v1/tasks/:id route in cmd/api/main.go
- [x] T034 [US4] Create integration test for successful task deletion in tests/integration/task_delete_test.go
- [x] T035 [US4] Create integration test for delete not found error in tests/integration/task_delete_test.go

**Checkpoint**: User Story 4 (Delete Task) is fully functional - full CRUD minus single-read

---

## Phase 7: User Story 5 - Get a Single Task (Priority: P3)

**Goal**: API consumers can retrieve a specific task by ID for efficient single-task lookup

**Independent Test**: Create task, GET /api/v1/tasks/{id}, verify task returned

### Implementation for User Story 5

- [x] T036 [US5] Implement GetTask handler in internal/handlers/task.go (parse id, call service, return 200 with TaskResponse)
- [x] T037 [US5] Add not found error handling in GetTask handler (return 404 with ErrorResponse)
- [x] T038 [US5] Register GET /api/v1/tasks/:id route in cmd/api/main.go
- [x] T039 [US5] Create integration test for successful task retrieval in tests/integration/task_get_test.go
- [x] T040 [US5] Create integration test for get not found error in tests/integration/task_get_test.go

**Checkpoint**: All CRUD operations complete - full Todo API functionality

---

## Phase 8: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [x] T041 [P] Add unit tests for TaskService in internal/services/task_test.go (mock repository)
- [x] T042 [P] Add unit tests for TaskRepository in internal/repository/task_test.go
- [x] T043 [P] Add unit tests for handlers in internal/handlers/task_test.go (mock service)
- [x] T044 Run gofmt on all Go files to ensure code formatting compliance
- [x] T045 Run go test -cover ./... to verify >80% test coverage (constitution requirement)
- [x] T046 [P] Add request logging middleware in cmd/api/main.go (gin.Default() includes Logger)
- [x] T047 [P] Add panic recovery middleware in cmd/api/main.go (gin.Default() includes Recovery)
- [x] T048 Validate API against openapi.yaml specification (manual or tool-based)
- [x] T049 Run quickstart.md validation (start API, execute all curl examples)

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3-7)**: All depend on Foundational phase completion
  - User stories can then proceed in parallel (if staffed)
  - Or sequentially in priority order (P1 → P2 → P3)
- **Polish (Phase 8)**: Depends on all desired user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2) - No dependencies on other stories
- **User Story 2 (P1)**: Can start after Foundational (Phase 2) - No dependencies on other stories
- **User Story 3 (P2)**: Can start after Foundational (Phase 2) - Independently testable
- **User Story 4 (P2)**: Can start after Foundational (Phase 2) - Independently testable
- **User Story 5 (P3)**: Can start after Foundational (Phase 2) - Independently testable

### Within Each User Story

- Handler implementation first
- Route registration follows handler
- Integration tests verify the endpoint

### Parallel Opportunities

- T003, T004, T005, T006 can run in parallel (dependencies)
- T015, T016, T017 can run in parallel (Docker, test helpers)
- User stories can all run in parallel after Phase 2 completes
- T041, T042, T043 can run in parallel (unit tests)
- T046, T047 can run in parallel (middleware)

---

## Parallel Example: Setup Phase

```bash
# Launch all dependency additions in parallel:
Task: "Add Gin dependency github.com/gin-gonic/gin to go.mod"
Task: "Add GORM dependencies gorm.io/gorm and gorm.io/driver/postgres to go.mod"
Task: "Add testify dependency github.com/stretchr/testify to go.mod"
Task: "Copy OpenAPI specification from specs/001-todo-api/contracts/openapi.yaml to openapi.yaml"
```

## Parallel Example: After Foundational Phase

```bash
# With multiple developers, all user stories can start in parallel:
Developer A: User Story 1 (T018-T022)
Developer B: User Story 2 (T023-T025)
Developer C: User Story 3 (T026-T030)
Developer D: User Story 4 (T031-T035)
Developer E: User Story 5 (T036-T040)
```

---

## Implementation Strategy

### MVP First (User Stories 1 & 2 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational (CRITICAL - blocks all stories)
3. Complete Phase 3: User Story 1 (Create Task)
4. Complete Phase 4: User Story 2 (List Tasks)
5. **STOP and VALIDATE**: Test US1 + US2 independently - this is the MVP!
6. Deploy/demo if ready

### Incremental Delivery

1. Complete Setup + Foundational → Foundation ready
2. Add User Story 1 → Test independently → Can create tasks
3. Add User Story 2 → Test independently → Can create and list tasks (MVP!)
4. Add User Story 3 → Test independently → Can update tasks
5. Add User Story 4 → Test independently → Can delete tasks
6. Add User Story 5 → Test independently → Can get single task (Full CRUD!)
7. Each story adds value without breaking previous stories

### Suggested MVP Scope

**MVP = User Story 1 + User Story 2** (both P1 priority)
- Create tasks (entry point for all data)
- List tasks (view what exists)
- This forms a minimal but complete todo system

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story for traceability
- Each user story should be independently completable and testable
- Constitution requires >80% test coverage - tests included in foundational and polish phases
- Constitution requires gofmt - included in polish phase
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
