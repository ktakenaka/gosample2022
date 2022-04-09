package kafka

import "github.com/Shopify/sarama"

type Config struct {
	Address string
}

func New(cfg *Config) (sarama.Client, error) {
	saramaCfg := sarama.NewConfig()
	return sarama.NewClient([]string{cfg.Address}, saramaCfg)
}

func NewConsumer(client sarama.Client) (sarama.Consumer, error) {
	return sarama.NewConsumerFromClient(client)
}

func NewProducer(client sarama.Client) (sarama.SyncProducer, error) {
	return sarama.NewSyncProducerFromClient(client)
}
