// taskmaster-cli/main.go
package main

import (
	"log"
	"taskmaster-cli/cli"
	"taskmaster-cli/storage"
	"taskmaster-cli/tasks"
)

func main() {
	// Initialize storage
	store := storage.NewJSONStorage("tasks.json")

	// Load existing tasks
	taskList, err := store.Load()
	if err != nil {
		log.Printf("Note: Could not load tasks: %v", err)
		taskList = &tasks.TaskList{Tasks: []tasks.Task{}}
	}

	// Parse and execute CLI command
	if err := cli.Execute(taskList, store); err != nil {
		log.Fatal(err)
	}
}
