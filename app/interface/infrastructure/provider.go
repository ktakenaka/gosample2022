package infrastructure

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-redis/redis/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Executor interface {
	boil.ContextExecutor
}
type DB interface {
	boil.ContextExecutor
	boil.ContextBeginner
}

type Redis interface {
	redis.Cmdable
}

type Mongo interface {
}

type KafkaConsumer interface {
	ReadMessage(timeout time.Duration) (*kafka.Message, error)
}

type Provider struct {
	DB            DB
	Redis         Redis
	Mongo         Mongo
	KafkaConsumer KafkaConsumer
}
