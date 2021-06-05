package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ktakenaka/gomsx/app/pkg/logger"
)

var defaultStogSigs = []os.Signal{os.Interrupt, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM}

type Tasks struct {
	log   logger.Logger
	tasks []Task
}

type Task interface {
	Shutdown(ctx context.Context) error
	Name() string
}

func NewShutdownTasks(l logger.Logger) *Tasks {
	return &Tasks{
		log:   l,
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

		t.log.Infof(ctx, "Shutting down %s...", task.Name())
		if err := task.Shutdown(ctx); err != nil {
			t.log.Errorf(ctx, "Failed to shutdown %s", task.Name())
		}
		t.tasks[i] = nil
	}
}

func (t *Tasks) WaitForServerStop(ctx context.Context) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, defaultStogSigs...)
	sig := <-sigChan
	t.log.Infof(ctx, "got stop sig: %s", sig.String())
}
