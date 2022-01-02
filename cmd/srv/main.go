package main

import (
	"context"

	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/database"
	"github.com/ktakenaka/gosample2022/cmd/internal/grpc"
	"github.com/ktakenaka/gosample2022/cmd/internal/notifier"
	"github.com/ktakenaka/gosample2022/cmd/internal/shutdown"
)

func main() {
	cfg, err := config.Initialize()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tasks := shutdown.New()
	defer tasks.Shutdown(ctx)

	read, write, task, err := database.Init(cfg)
	if err != nil {
		panic(err)
	}
	tasks.Add(task)

	ntfr, task := notifier.Init(cfg)
	tasks.Add(task)

	task, err = grpc.New(cfg, read, write, ntfr)
	if err != nil {
		panic(err)
	}
	tasks.Add(task)

	tasks.WaitForStopSignal(ctx)
}
