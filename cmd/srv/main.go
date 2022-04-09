package main

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/interface/infrastructure"
	ntfr "github.com/ktakenaka/gosample2022/app/pkg/notifier"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/grpc"
	"github.com/ktakenaka/gosample2022/cmd/internal/kafkaclient"
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

	provider := &infrastructure.Provider{}
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

	provider.Kafka, task, err = kafkaclient.Init(ctx, cfg)
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
