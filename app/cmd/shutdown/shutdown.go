package shutdown

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var defaultStogSigs = []os.Signal{os.Interrupt, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM}

type Tasks struct {
	// TODO: define logger interface in app/pkg/log
	log   log.Logger
	tasks []Task
}

type Task interface {
	Shutdown(ctx context.Context) error
	Name() string
}

func NewShutdownTasks() *Tasks {
	return &Tasks{
		log:   log.Logger{},
		tasks: make([]Task, 0),
	}
}

func (t *Tasks) Add(task Task) {
	t.tasks = append(t.tasks, task)
}

func (t *Tasks) ExecuteAll(ctx context.Context) {
	for i := len(t.tasks) - 1; i >= 0; i-- {
		task := t.tasks[i]
		if task == nil {
			continue
		}

		if err := task.Shutdown(ctx); err != nil {
			t.log.Printf("Failed to shutdown %s", task.Name())
		}
		t.tasks[i] = nil
	}
}

func WaitForServerStop(ctx context.Context) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, defaultStogSigs...)
	sig := <-sigChan
	log.Printf("got stop sig: %s", sig.String())
}
