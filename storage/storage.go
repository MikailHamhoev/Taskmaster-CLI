// taskmaster-cli/storage/storage.go
package storage

import (
	"encoding/json"
	"os"
	"taskmaster-cli/tasks"
)

type JSONStorage struct {
	filename string
}

func NewJSONStorage(filename string) *JSONStorage {
	return &JSONStorage{filename: filename}
}

func (s *JSONStorage) Load() (*tasks.TaskList, error) {
	data, err := os.ReadFile(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return &tasks.TaskList{Tasks: []tasks.Task{}, NextID: 1}, nil
		}
		return nil, err
	}

	var taskList tasks.TaskList
	if err := json.Unmarshal(data, &taskList); err != nil {
		return nil, err
	}

	// Initialize NextID if it's zero (for backward compatibility)
	if taskList.NextID == 0 {
		maxID := 0
		for _, task := range taskList.Tasks {
			if task.ID > maxID {
				maxID = task.ID
			}
		}
		taskList.NextID = maxID + 1
	}

	return &taskList, nil
}

func (s *JSONStorage) Save(taskList *tasks.TaskList) error {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, data, 0644)
}
