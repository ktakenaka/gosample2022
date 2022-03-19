package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URI string
}

func New(ctx context.Context, cfg *Config) (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
}
