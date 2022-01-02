package tmanager

import (
	"context"
	"log"

	"github.com/ktakenaka/gosample2022/app/config"
	"github.com/ktakenaka/gosample2022/app/domain/repository"
	"github.com/ktakenaka/gosample2022/cmd/database"
)

type Task interface {
	Name() string
	Shutdown(ctx context.Context) error
}

// TManager: Task Manager
type TManager struct {
	WriteDBFactory repository.DBWriteFunc
	ReadDBFactory  repository.DBReadFunc

	tasks []Task
}

func (t *TManager) Add(task Task) {
	t.tasks = append(t.tasks, task)
}

func (t *TManager) Shutdown(ctx context.Context) {
	for i := range t.tasks {
		log.Printf("shutting down %s...", t.tasks[i].Name())
		err := t.tasks[i].Shutdown(ctx)
		log.Printf("%v", err)
	}
}

func (t *TManager) InitDB(cfg *config.Config) error {
	read, write, task, err := database.Init(cfg.DB.Write, cfg.DB.Read)
	if err != nil {
		return err
	}

	t.ReadDBFactory = read
	t.WriteDBFactory = write
	t.Add(task)
	return nil
}
