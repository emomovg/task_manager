package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

const (
	add = iota + 1
	show
	remove
	saveAndExit
)

const filename = "tasks.json"

func main() {
	var command int
	myCont, err := loadData()

	if err != nil {
		fmt.Printf("File opening error: %v\n", err)
		return
	}

	fmt.Print("Task Manager\n-----------\n1. Add task\n2. List tasks\n3. Remove task\n4. Exit \n")
OuterLoop:
	for {
		fmt.Print("> ")
		fmt.Scan(&command)
		switch command {
		case add:
			fmt.Print("Enter task title: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				title := scanner.Text()
				myCont.Add(title)
				fmt.Println("Task added!")
				break
			}

			fmt.Println("something went wrong!")
		case show:
			myCont.ShowAll()
		case remove:
			var id int
			fmt.Print("Enter id for remove: ")
			fmt.Scan(&id)
			err := myCont.Delete(id)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("Task deleted!")

		case saveAndExit:
			saveErr := saveTasks(*myCont.TSlice)
			if saveErr != nil {
				fmt.Printf("File saving error: %v\n", saveErr)
				break OuterLoop
			}
			fmt.Println("Saving tasks... Bye!")
			break OuterLoop
		default:
			fmt.Println("Выбрана неправильная команда!")
		}
	}
}

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskMap map[int]Task
type TaskSlice []Task

type TaskContainer struct {
	TMap   TaskMap
	TSlice *TaskSlice
}

func (t *TaskContainer) GetMaxKey() int {
	var maxKey int
	for key, _ := range t.TMap {
		if maxKey <= key {
			maxKey = key
		}
	}

	return maxKey
}

func (t *TaskContainer) Add(title string) {
	maxKey := t.GetMaxKey()
	newTask := Task{
		ID:    maxKey + 1,
		Title: title,
	}
	t.TMap[maxKey+1] = newTask
	*t.TSlice = append(*t.TSlice, newTask)
}

func (t *TaskContainer) Delete(id int) error {
	if _, ok := t.TMap[id]; !ok {
		return fmt.Errorf("id = %v не существует в файле", id)

	}
	delete(t.TMap, id)
	var deleteKey = 0
	for i, value := range *t.TSlice {
		if value.ID == id {
			deleteKey = i
		}
	}
	*t.TSlice = append((*t.TSlice)[:deleteKey], (*t.TSlice)[(deleteKey+1):]...)

	return nil
}

func (t *TaskContainer) ShowAll() {
	fmt.Println("Tasks:")
	for _, task := range *t.TSlice {
		taskResolve := "not done"
		if task.Done {
			taskResolve = "done"
		}
		fmt.Printf("[%v] %v (%v)\n", task.ID, task.Title, taskResolve)
	}
	fmt.Printf("\n")
}

func NewTaskContainer() *TaskContainer {
	return &TaskContainer{
		TMap:   make(TaskMap),
		TSlice: &TaskSlice{},
	}
}

func loadData() (*TaskContainer, error) {
	taskCont := NewTaskContainer()
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return taskCont, nil
		}
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&taskCont.TSlice)
	if err != nil {
		return nil, err
	}

	for _, task := range *taskCont.TSlice {
		taskCont.TMap[task.ID] = task
	}

	return taskCont, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}
