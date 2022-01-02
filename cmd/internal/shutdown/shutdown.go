package shutdown

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var stopSigs = []os.Signal{syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM}

type Task interface {
	Name() string
	Shutdown(ctx context.Context) error
}

type tasks struct {
	tasks   []Task
	sigChan chan os.Signal
}

func New() *tasks {
	return &tasks{
		sigChan: make(chan os.Signal),
	}
}

func (ts *tasks) Add(task Task) {
	ts.tasks = append(ts.tasks, task)
}

func (ts *tasks) Shutdown(ctx context.Context) {
	for i := range ts.tasks {
		log.Printf("%s: shutting down...\n", ts.tasks[i].Name())
		if err := ts.tasks[i].Shutdown(ctx); err != nil {
			log.Printf("%s: %v\n", ts.tasks[i].Name(), err)
		}
	}
}

func (ts *tasks) WaitForStopSignal(ctx context.Context) {
	signal.Notify(ts.sigChan, stopSigs...)
	sig := <-ts.sigChan

	log.Printf("got stop sig: %s", sig.String())
}
