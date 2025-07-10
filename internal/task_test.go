package internal

import "testing"

var TaskSlice TaskList = []Task{
	{
		ID:    1,
		Title: "hello world",
		Done:  true,
	},
	{
		ID:    2,
		Title: "Good luck",
		Done:  true,
	},
	{
		ID:    3,
		Title: "Life is good",
		Done:  true,
	},
	{
		ID:    4,
		Title: "Stepik course!",
		Done:  false,
	},
	{
		ID:    5,
		Title: "Go lang task",
		Done:  false,
	},
}

var manager = TaskManager{
	TMap: map[int]Task{
		1: {
			ID:    1,
			Title: "hello world",
			Done:  true,
		},
		5: {
			ID:    5,
			Title: "Go lang task",
			Done:  false,
		},
		4: {
			ID:    4,
			Title: "Stepik course!",
			Done:  false,
		},
		2: {
			ID:    2,
			Title: "Good luck",
			Done:  true,
		},
		3: {
			ID:    3,
			Title: "Life is good",
			Done:  true,
		},
	},
	TSlice: &TaskSlice,
}

func TestGetMaxKey(t *testing.T) {
	got := manager.GetMaxKey()
	want := 5
	if want != got {
		t.Errorf("Expected max key = %v", want)
	}
}

func TestAdd(t *testing.T) {
	countTask := len(*manager.TSlice)
	wantCount := countTask + 1
	wantNewId := manager.GetMaxKey() + 1
	manager.Add("testing func")
	if len(*manager.TSlice) != wantCount {
		t.Errorf("expected size to remain %d, but got %d", wantCount, countTask)
	}

	if _, exists := manager.TMap[wantNewId]; !exists {
		t.Errorf("expected key %v to be present in the map, but it was not found", wantNewId)
	}
}

func TestDelete(t *testing.T) {
	err := manager.Delete(4)

	if err != nil {
		t.Errorf("expected err %v, but got %v", nil, err)
	}

	if _, exists := manager.TMap[4]; exists {
		t.Errorf("expected key %v deleted, but it was not deleted", 4)
	}

	for _, task := range *manager.TSlice {
		if task.ID == 4 {
			t.Errorf("expected Task with id 4 deleted, but it was not deleted")
		}
	}
}
