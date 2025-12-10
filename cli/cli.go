// taskmaster-cli/cli/cli.go
package cli

import (
	"flag"
	"fmt"
	"os"
	"taskmaster-cli/tasks"
)

func Execute(taskList *tasks.TaskList, store Storage) error {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	// Define flags for list command
	listAll := listCmd.Bool("all", false, "Show all tasks including completed")

	// Define flags for complete command
	completeID := completeCmd.Int("id", 0, "ID of task to mark as complete")

	// Define flags for delete command
	deleteID := deleteCmd.Int("id", 0, "ID of task to delete")

	if len(os.Args) < 2 {
		printUsage()
		return nil
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if addCmd.NArg() == 0 {
			return fmt.Errorf("description required for add command")
		}
		description := addCmd.Arg(0)
		task := taskList.Add(description)
		fmt.Printf("Added task #%d: %s\n", task.ID, task.Description)

	case "list":
		listCmd.Parse(os.Args[2:])
		tasksToShow := taskList.List(*listAll)
		if len(tasksToShow) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}

		fmt.Println("Tasks:")
		for _, task := range tasksToShow {
			status := "□"
			if task.Completed {
				status = "✓"
			}
			fmt.Printf("  %s #%d: %s (created: %s)\n",
				status, task.ID, task.Description,
				task.CreatedAt.Format("2006-01-02"))
		}

	case "complete":
		completeCmd.Parse(os.Args[2:])
		if *completeID == 0 {
			return fmt.Errorf("id is required for complete command")
		}
		if err := taskList.Complete(*completeID); err != nil {
			return err
		}
		fmt.Printf("Task #%d marked as complete\n", *completeID)

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteID == 0 {
			return fmt.Errorf("id is required for delete command")
		}
		if err := taskList.Delete(*deleteID); err != nil {
			return err
		}
		fmt.Printf("Task #%d deleted\n", *deleteID)

	case "help":
		printUsage()
		return nil

	default:
		printUsage()
		return fmt.Errorf("unknown command: %s", os.Args[1])
	}

	// Save changes after each command
	return store.Save(taskList)
}

func printUsage() {
	fmt.Println(`TaskMaster - Simple CLI Task Manager

Usage:
  taskmaster <command> [arguments]

Commands:
  add <description>    Add a new task
  list [-all]          List tasks (use -all to show completed)
  complete -id <id>    Mark a task as complete
  delete -id <id>      Delete a task
  help                 Show this help message

Examples:
  taskmaster add "Buy groceries"
  taskmaster list
  taskmaster complete -id 1
  taskmaster delete -id 1`)
}

// Storage interface for dependency injection
type Storage interface {
	Save(*tasks.TaskList) error
	Load() (*tasks.TaskList, error)
}
