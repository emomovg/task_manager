package internal

import "fmt"

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskStore map[int]Task
type TaskList []Task

type TaskManager struct {
	TMap   TaskStore
	TSlice *TaskList
}

func (t *TaskManager) GetMaxKey() int {
	var maxKey int
	for key := range t.TMap {
		if maxKey <= key {
			maxKey = key
		}
	}

	return maxKey
}

func (t *TaskManager) Add(title string) {
	maxKey := t.GetMaxKey()
	newTask := Task{
		ID:    maxKey + 1,
		Title: title,
	}
	t.TMap[maxKey+1] = newTask
	*t.TSlice = append(*t.TSlice, newTask)
}

func (t *TaskManager) Delete(id int) error {
	if _, ok := t.TMap[id]; !ok {
		return fmt.Errorf("id = %v не существует в файле", id)

	}
	delete(t.TMap, id)
	var deleteKey = 0
	for i, task := range *t.TSlice {
		if task.ID == id {
			deleteKey = i
		}
	}
	*t.TSlice = append((*t.TSlice)[:deleteKey], (*t.TSlice)[(deleteKey+1):]...)

	return nil
}

func (t *TaskManager) GetAllTasks() []Task {
	return *t.TSlice
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		TMap:   make(TaskStore),
		TSlice: &TaskList{},
	}
}