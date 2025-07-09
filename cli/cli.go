package cli

import (
	"bufio"
	"fmt"
	"github.com/emomovg/task_manager/internal"
	"log"
	"os"
)

const (
	add         = 1
	show        = 2
	remove      = 3
	saveAndExit = 4
)

func Run(manager *internal.TaskManager) error {
	var command int
	for {
		fmt.Print("> ")
		_, err := fmt.Scan(&command)
		if err != nil {
			return fmt.Errorf("failed to read command: %w", err)
		}

		switch command {
		case add:
			addTask(manager)
		case show:
			showAllTasks(manager)
		case remove:
			deleteTask(manager)
		case saveAndExit:
			err = SaveAndExit(manager)
			return err
		default:
			fmt.Println("Wrong operation selected!")
		}
	}
}

func showAllTasks(tm *internal.TaskManager) {
	fmt.Println("Tasks:")
	tasks := tm.GetAllTasks()
	for _, task := range tasks {
		status := "not done"
		if task.Done {
			status = "done"
		}
		fmt.Printf("ID: %d, Title: %s, Done: %s\n", task.ID, task.Title, status)
	}
}

func addTask(manager *internal.TaskManager) {
	fmt.Print("Enter task title: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Printf("Error reading task title: %v\n", scanner)
		return
	}
	title := scanner.Text()
	manager.Add(title)
	fmt.Println("Task added!")
}

func deleteTask(manager *internal.TaskManager) {
	var id int
	fmt.Print("Enter id for remove: ")
	_, er := fmt.Scan(&id)
	if er != nil {
		log.Printf("failed to read command: %v\n", er)
		return
	}
	err := manager.Delete(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Task deleted!")
}

func SaveAndExit(manager *internal.TaskManager) error {
	err := internal.SaveTasks(*manager.TSlice)
	if err != nil {
		return err
	}
	fmt.Println("Saving tasks... Bye!")
	return nil
}
