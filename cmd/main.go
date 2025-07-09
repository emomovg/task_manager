package main

import (
	"fmt"
	"github.com/emomovg/task_manager/cli"
	"github.com/emomovg/task_manager/internal"
	"log"
)

func main() {
	manager, err := internal.LoadData()
	if err != nil {
		log.Fatalf("failed to load tasks: %v\n", err)
	}
	fmt.Print("Task Manager\n-----------\n1. Add task\n2. List tasks\n3. Remove task\n4. Exit \n")
	err = cli.Run(manager)
	if err != nil {
		fmt.Println(err)
	}
}