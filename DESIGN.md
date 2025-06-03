# Design Document — Go TODO CLI

---

## Context
The project is implemented as part of the course *"Methodologies and Software Development Technologies"*.

**Team Members:**
- D-Vasylchenko – Team Lead, Developer  
- Anirina Ponaho – Developer, Reviewer

---

## Goals and Non-Goals

### Goals:
- Build a simple, user-friendly CLI-based task tracking application using Go.
- Implement essential task operations: add, list, mark as complete, and delete.
- Use JSON file for task persistence between application runs.
- Demonstrate team collaboration using GitHub (branches, pull requests, code review).
- Integrate GitHub Actions for CI and automatic test execution.

### Non-Goals:
- No user authentication or authorization (single-user only).
- No GUI or web interface.
- No cloud-based or SQL data storage.
- No notification system (e.g., email reminders).

---

## Subsystems

### Architecture Overview
The system consists of the following modules:
- `cmd/` — CLI entry point (main.go)
- `task/` — task model and manager logic
- `storage/` — file-based JSON persistence
- `tests/` — unit tests for all modules
- `.github/workflows/` — CI pipeline for test validation

```
[CLI] → [Task Manager] → [Storage]
```

### Responsibilities
- **CLI**: reads input from the user and maps commands to logic.
- **Task Manager**: core logic of adding, listing, deleting, toggling task status.
- **Storage**: responsible for loading and saving tasks in a file.

All subsystems interact via exported functions. Standard Go modules will be used without third-party dependency injection libraries.

---

## Authorization Subsystem

**Not applicable.**  
This is a local CLI utility designed for single-user use. Authentication, session management, or user roles are out of scope.

---

## Data Storage Subsystem

Tasks are saved to and loaded from a file `tasks.json`.

### Data Format Example:
```json
[
  {
    "id": 1,
    "name": "Buy groceries",
    "done": false
  },
  {
    "id": 2,
    "name": "Write lab report",
    "done": true
  }
]
```

The storage layer includes functions:
- `LoadTasks() ([]Task, error)` — reads from JSON file
- `SaveTasks([]Task) error` — writes to JSON file

To ensure data integrity, tasks will be sorted and validated before saving.

---

## Business Logic

### Main features:
- Add new task with automatic ID assignment
- List all tasks (separating done and not done)
- Delete task by ID
- Mark task as done or undone
- Clear all completed tasks (additional feature)

### Command examples:
```bash
go run ./cmd/main.go add "Read design doc"
go run ./cmd/main.go list
```

Each action is implemented as a function in the task manager:
- `AddTask(name string)`
- `ListTasks()`
- `CompleteTask(id int)`
- `DeleteTask(id int)`
- `ClearCompleted()` (extra functionality)

---

## HTTP Frontend

**Not applicable.**  
This is a CLI-only application.

---

## Alternatives Considered

- **Database (e.g., SQLite)**: unnecessary for small local use. JSON file is sufficient.
- **Web Interface**: rejected due to complexity and irrelevance to CLI goals.
- **Third-party CLI libraries**: avoided to keep dependencies minimal and maintain educational value.
- **CSV file format**: less structured and harder to maintain than JSON.

---

## Final Notes

This project is ideal for demonstrating:
- Practical use of Go modules and JSON handling
- Command-line UX principles
- GitHub collaboration and CI

The extra feature `ClearCompleted()` was added to demonstrate iterative improvement and better user convenience.
