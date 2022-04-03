package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

const (
	queueName = "example_queue"
	taskName  = "example_task"
)

func main() {
	server, err := newServer()
	if err != nil {
		panic(err)
	}

	worker := server.NewWorker("worker_name", 10)
	go func() {
		if err := worker.Launch(); err != nil && !errors.Is(err, machinery.ErrWorkerQuitGracefully) {
			panic(err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.Tick(3 * time.Second):
				fmt.Println()
				sendTask(ctx, server)
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}

func newServer() (*machinery.Server, error) {
	cfg := &config.Config{
		DefaultQueue:  queueName,
		Broker:        "redis://@redis:6379",
		ResultBackend: "redis://@redis:6379",
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}
	server, err := machinery.NewServer(cfg)
	if err != nil {
		return nil, err
	}

	return server, server.RegisterTask(taskName, exampleTask)
}

func exampleTask(arg string) error {
	fmt.Println(arg)
	return nil
}

func sendTask(ctx context.Context, server *machinery.Server) {
	signature := &tasks.Signature{
		Name: taskName,
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: "hello",
			},
		},
	}

	asyncResult, err := server.SendTask(signature)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; true; i++ {
		state := asyncResult.GetState()
		if state.State == tasks.StateSuccess {
			fmt.Println("Checked...", i, "times")
			break
		}
		<-time.Tick(2 * time.Second)
	}
}
