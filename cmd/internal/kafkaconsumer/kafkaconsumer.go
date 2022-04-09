package kafkaconsumer

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ktakenaka/gosample2022/app/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/shutdown"
)

type task struct {
	c *kafka.Consumer
}

func (t *task) Name() string {
	return "kafkaconsumer"
}
func (t *task) Shutdown(ctx context.Context) error {
	return t.c.Close()
}

func Init(ctx context.Context, cfg *config.Config) (*kafka.Consumer, shutdown.Task, error) {
	c, err := kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers": cfg.KafkaConsumer.BootStrapServer,
			"group.id":          "gosample2022",
			"auto.offset.reset": cfg.KafkaConsumer.AutoOffsetReset,
		},
	)
	if err != nil {
		return nil, nil, err
	}

	if err := c.SubscribeTopics([]string{cfg.KafkaConsumer.TopicSamples, cfg.KafkaConsumer.TopicTransactions}, nil); err != nil {
		c.Close()
		return nil, nil, err
	}

	return c, &task{c: c}, nil
}
