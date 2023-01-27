package test

import (
	"testing"
	"testtask"
)

func TestCreateAndGet(t *testing.T) {
	// Create a store and a single task.
	ts := main.New()
	id := ts.CreateTask("Hola", "nil", make(map[string]string))

	// We should be able to retrieve this task by ID, but nothing with other
	// IDs.
	task, err := ts.GetTask(id)
	if err != nil {
		t.Fatal(err)
	}

	if task.Id != id {
		t.Errorf("got task.Id=%d, id=%d", task.Id, id)
	}
	if task.Method != "Hola" {
		t.Errorf("got Method=%v, want %v", task.Method, "Hola")
	}
	if task.Url != "Hola" {
		t.Errorf("got Method=%v, want %v", task.Url, "Hola")
	}

	_, err = ts.GetTask(id + 1)
	if err == nil {
		t.Fatal("got nil, want error")
	}
}
