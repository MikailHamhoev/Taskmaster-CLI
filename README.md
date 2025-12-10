# TaskMaster CLI

A lightweight, command-line task manager built with Go's standard library.

## Skills
Go standard library (`flag`, `encoding/json`, `os`), CLI development, file I/O, JSON serialization, project structure, error handling, time manipulation

## Features
- Add tasks with descriptions  
- List pending or all tasks  
- Mark tasks as complete  
- Delete tasks by ID  
- Persistent JSON storage  
- Intuitive CLI interface  

## Project Structure
```
taskmaster-cli/
├── main.go
├── cli/          # Command parsing
├── tasks/        # Task logic and structs
├── storage/      # JSON persistence
└── tasks.json    # Auto-created data file
```

## Usage
```bash
# Build
go build -o taskmaster

# Commands
taskmaster add "Buy groceries"
taskmaster list
taskmaster list -all
taskmaster complete -id 1
taskmaster delete -id 1
taskmaster help
```

## Data Storage
Tasks saved to `tasks.json` in current directory using human-readable JSON format.

## Dependencies
Go 1.21+  
No external dependencies — pure standard library.

MIT License