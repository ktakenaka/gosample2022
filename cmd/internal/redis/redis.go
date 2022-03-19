package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/ktakenaka/gosample2022/app/config"
	"github.com/ktakenaka/gosample2022/app/domain/repository"
	"github.com/ktakenaka/gosample2022/cmd/internal/shutdown"
	infraRedis "github.com/ktakenaka/gosample2022/infra/redis"
)

type task struct {
	client *redis.Client
}

func (t *task) Shutdown(ctx context.Context) error {
	return t.client.Close()
}

func (t *task) Name() string {
	return "redis"
}

func Init(ctx context.Context, cfg *config.Config) (repository.Redis, shutdown.Task, error) {
	client, err := infraRedis.New(
		&infraRedis.Config{
			URL:          cfg.Redis.URL,
			PoolSize:     cfg.Redis.PoolSize,
			MinIdleConns: cfg.Redis.MinIdleConns,
			ReadTimeout:  cfg.Redis.ReadTimeout,
			WriteTimeout: cfg.Redis.WriteTimeout,
			UseTLS:       cfg.Redis.UseTLS,
		},
	)
	if err != nil {
		return nil, nil, err
	}

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, nil, err
	}
	return client, &task{client: client}, nil
}
