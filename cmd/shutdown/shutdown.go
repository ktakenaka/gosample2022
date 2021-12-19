package shutdown

import (
	"context"
	"log"
)

type Task func(ctx context.Context) error

type tasks struct {
	tasks []Task
}

func New() *tasks {
	return &tasks{}
}

func (ts *tasks) Add(task Task) {
	ts.tasks = append(ts.tasks, task)
}

func (ts *tasks) Shutdown(ctx context.Context) {
	for i := range ts.tasks {
		err := ts.tasks[i](ctx)
		log.Printf("%v\n", err)
	}
}
