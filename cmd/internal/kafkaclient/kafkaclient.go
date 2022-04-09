package kafkaclient

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/ktakenaka/gosample2022/app/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/shutdown"
)

type task struct {
	c sarama.Client
}

func (t *task) Name() string {
	return "kafkaconsumer"
}
func (t *task) Shutdown(ctx context.Context) error {
	return t.c.Close()
}

func Init(ctx context.Context, cfg *config.Config) (sarama.Client, shutdown.Task, error) {
	saramaCfg := sarama.NewConfig()
	c, err := sarama.NewClient([]string{cfg.Kafka.Address}, saramaCfg)

	return c, &task{c: c}, err
}
