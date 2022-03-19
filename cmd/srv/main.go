package main

import (
	"context"

	ntfr "github.com/ktakenaka/gosample2022/app/pkg/notifier"
	"github.com/ktakenaka/gosample2022/app/registry"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/grpc"
	"github.com/ktakenaka/gosample2022/cmd/internal/mongodb"
	"github.com/ktakenaka/gosample2022/cmd/internal/mysql"
	"github.com/ktakenaka/gosample2022/cmd/internal/notifier"
	"github.com/ktakenaka/gosample2022/cmd/internal/redis"
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

	task, err := notifier.Init(ctx, cfg)
	if err != nil {
		panic(err)
	}
	tasks.Add(task)

	provider := &registry.Provider{}
	provider.DB, task, err = mysql.Init(ctx, cfg)
	if err != nil {
		ntfr.Error(err)
		panic(err)
	}
	tasks.Add(task)

	provider.Redis, task, err = redis.Init(ctx, cfg)
	if err != nil {
		ntfr.Error(err)
		panic(err)
	}
	tasks.Add(task)

	provider.Mongo, task, err = mongodb.Init(ctx, cfg)
	if err != nil {
		ntfr.Error(err)
		panic(err)
	}
	tasks.Add(task)

	task, err = grpc.New(cfg, provider)
	if err != nil {
		ntfr.Error(err)
		panic(err)
	}
	tasks.Add(task)

	tasks.WaitForStopSignal(ctx)
}
