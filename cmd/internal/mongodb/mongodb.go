package mongodb

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/shutdown"
	"github.com/ktakenaka/gosample2022/infra/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type task struct {
	client *mongo.Client
}

func (t *task) Name() string {
	return "mongo"
}
func (t *task) Shutdown(ctx context.Context) error {
	return t.client.Disconnect(ctx)
}

func Init(ctx context.Context, cfg *config.Config) (*mongo.Client, shutdown.Task, error) {
	client, err := mongodb.New(ctx, &mongodb.Config{
		URI: cfg.Mongo.URL,
	})
	if err != nil {
		return nil, nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, nil, err
	}

	return client, &task{client: client}, nil
}
