package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Id         int               `json:"id"`
	Status     string            `json:"status"`
	Method     string            `json:"method"`
	Url        string            `json:"url"`
	Headers    map[string]string `json:"headers"`
	CreateDate time.Time         `json:"create_date"`
}

// simple In-Memory concurrency storage
type TaskStore struct {
	sync.Mutex

	tasks  map[int]Task
	nextId int
}

func New() *TaskStore {
	ts := &TaskStore{}
	ts.tasks = make(map[int]Task)
	ts.nextId = 0
	return ts
}

func (ts *TaskStore) CreateTask(method string, url string, headers map[string]string) int {
	ts.Lock()
	defer ts.Unlock()

	task := Task{
		Id:         ts.nextId,
		Method:     method,
		Url:        url,
		CreateDate: time.Now(),
	}

	task.Headers = make(map[string]string, len(headers))
	for k, v := range headers {
		task.Headers[k] = v
	}

	ts.tasks[ts.nextId] = task
	ts.nextId++
	return task.Id
}

// GetTask retrieves a task from the store, by id. If no such id exists, an
// error is returned.
func (ts *TaskStore) GetTask(id int) (Task, error) {
	ts.Lock()
	defer ts.Unlock()

	t, ok := ts.tasks[id]
	if ok {
		return t, nil
	} else {
		return Task{}, fmt.Errorf("task with id=%d not found", id)
	}
}
