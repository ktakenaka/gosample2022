package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ktakenaka/gosample2022/app/pkg/debeziumcsmr"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/kafkaclient"
	"github.com/ktakenaka/gosample2022/cmd/internal/redis"
	"github.com/ktakenaka/gosample2022/infra/kafka"
)

func main() {
	ctx := context.Background()
	cfg, _ := config.Initialize()

	kafkaClient, task, _ := kafkaclient.Init(ctx, cfg)
	defer task.Shutdown(ctx)
	redisClient, task, _ := redis.Init(ctx, cfg)
	defer task.Shutdown(ctx)

	csmer, _ := kafka.NewConsumer(kafkaClient)
	pcsmer, err := csmer.ConsumePartition(
		"gosample2022_dbserver.transaction",
		0,
		0,
	)
	if err != nil {
		panic(err)
	}
	for msg := range pcsmer.Messages() {
		payload := &debeziumcsmr.TransactionPayload{}
		if err := json.Unmarshal(msg.Value, payload); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("=== ID", payload.Payload.ID)
		fmt.Println("Status", payload.Payload.Status)
		fmt.Println("EventCount", payload.Payload.EventCount)
		fmt.Println(payload.Payload.DataCollections)
	}
}
