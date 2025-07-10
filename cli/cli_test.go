package cli

import (
	"bytes"
	"github.com/emomovg/task_manager/internal"
	"io"
	"os"
	"testing"
)

var testTasks = internal.TaskList{
	{ID: 1, Title: "First task", Done: false},
	{ID: 2, Title: "Second task", Done: true},
}

var testManager = internal.TaskManager{
	TMap: internal.TaskStore{
		1: {ID: 1, Title: "First task", Done: false},
		2: {ID: 2, Title: "Second task", Done: true},
	},
	TSlice: &testTasks,
}

func TestShowAllTasks(t *testing.T) {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	ShowAllTasks(&testManager)
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	os.Stdout = originalStdout

	got := buf.String()
	want := `Tasks:
ID: 1, Title: First task, Done: not done
ID: 2, Title: Second task, Done: done
`

	if got != want {
		t.Errorf("unexpected output:\nGot:\n%s\nWant:\n%s", got, want)
	}
}
