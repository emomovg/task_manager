package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const filename = "tasks.json"

func LoadData() (*TaskManager, error) {
	taskManager := NewTaskManager()
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return taskManager, nil
		}
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer func(file *os.File) {
		er := file.Close()
		if er != nil {
			log.Printf("WARNING: Failed to close file %s: %v", filename, er)
		}
	}(file)

	err = json.NewDecoder(file).Decode(&taskManager.TSlice)
	if err != nil {
		return nil, fmt.Errorf("failed to decoder: %w", err)
	}

	for _, task := range *taskManager.TSlice {
		taskManager.TMap[task.ID] = task
	}

	return taskManager, nil
}

func SaveTasks(tasks []Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func(file *os.File) {
		er := file.Close()
		if er != nil {
			log.Printf("WARNING: Failed to close file %s: %v", filename, er)
		}
	}(file)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(tasks)
}
