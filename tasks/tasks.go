// taskmaster-cli/tasks/tasks.go
package tasks

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
}

type TaskList struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

func (tl *TaskList) Add(description string) Task {
	task := Task{
		ID:          tl.NextID,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	tl.Tasks = append(tl.Tasks, task)
	tl.NextID++
	return task
}

func (tl *TaskList) Complete(id int) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks[i].Completed = true
			tl.Tasks[i].CompletedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func (tl *TaskList) Delete(id int) error {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func (tl *TaskList) List(showCompleted bool) []Task {
	if showCompleted {
		return tl.Tasks
	}

	var pending []Task
	for _, task := range tl.Tasks {
		if !task.Completed {
			pending = append(pending, task)
		}
	}
	return pending
}

func (tl *TaskList) FindByID(id int) (*Task, error) {
	for _, task := range tl.Tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", id)
}
