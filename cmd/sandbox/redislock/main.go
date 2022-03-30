package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	infraRedis "github.com/ktakenaka/gosample2022/cmd/internal/redis"
)

var (
	cfg, _ = config.Initialize()
	ctx    = context.Background()
)

const (
	key = "key"
)

func main() {
	client, task, _ := infraRedis.Init(ctx, cfg)
	defer task.Shutdown(ctx)

	client.Set(ctx, key, "fuga", 0)

	var wg sync.WaitGroup
	wg.Add(1)
	fn := func(tx *redis.Tx) error {
		defer wg.Done()
		tx.Set(ctx, key, "hoge", 0)
		time.Sleep(2 * time.Second)
		return nil
	}
	go client.Watch(ctx, fn, key)

	time.Sleep(1 * time.Second)
	fmt.Println(client.Set(ctx, key, "baz", 0).String())

	wg.Wait()
}
