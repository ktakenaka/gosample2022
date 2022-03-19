package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	cmdRedis "github.com/ktakenaka/gosample2022/cmd/internal/redis"
	"github.com/ktakenaka/gosample2022/cmd/internal/shutdown"
)

const (
	chanName = "sandbox-redis-pubsub"
)

var (
	cfg, _ = config.Initialize()
	ctx    = context.Background()

	stopSigs = []os.Signal{syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM}
)

type task struct {
	name    string
	release func()
}

func (t *task) Name() string {
	return t.name
}
func (t *task) Shutdown(ctx context.Context) error {
	t.release()
	return nil
}

func main() {
	tasks := shutdown.New()
	defer tasks.Shutdown(ctx)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	client, task, err := cmdRedis.Init(ctx, cfg)
	if err != nil {
		panic(err)
	}
	tasks.Add(task)

	tasks.Add(publish(ctx, client))
	tasks.Add(subscribe(ctx, client))

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, stopSigs...)
	sig := <-sigChan
	fmt.Println("-----get-----", sig)
}

func publish(ctx context.Context, client *redis.Client) shutdown.Task {
	finish := make(chan struct{})
	t := &task{name: "publisher", release: func() { <-finish }}

	go func() {
		defer func() {
			finish <- struct{}{}
		}()

		var i int
		for {
			select {
			case <-ctx.Done():
				return
			default:
				client.Publish(ctx, chanName, fmt.Sprintf("hello %d", i))
				fmt.Println("publish", i)
				time.Sleep(2 * time.Second)
				i++
			}
		}
	}()

	return t
}

func subscribe(ctx context.Context, client *redis.Client) shutdown.Task {
	finish := make(chan struct{})
	t := &task{name: "subscriber", release: func() { <-finish }}
	pubsub := client.Subscribe(ctx, chanName)

	go func() {
		defer func() {
			fmt.Println("close pubsub:", pubsub.Close())
			finish <- struct{}{}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-pubsub.Channel():
				select {
				case <-ctx.Done():
					fmt.Println("subscribe", msg.Payload)
					return
				default:
					fmt.Println("subscribe", msg.Payload)
				}
			}
		}
	}()
	return t
}
