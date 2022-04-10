package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-cmp/cmp"

	"github.com/ktakenaka/gosample2022/app/pkg/debeziumcsmr"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/kafkaclient"
	"github.com/ktakenaka/gosample2022/infra/kafka"
)

func main() {
	ctx := context.Background()
	cfg, _ := config.Initialize()

	kafkaClient, task, _ := kafkaclient.Init(ctx, cfg)
	defer task.Shutdown(ctx)

	csmer, _ := kafka.NewConsumer(kafkaClient)
	pcsmer, err := csmer.ConsumePartition(
		"gosample2022_dbserver.gosample2022_development.samples",
		0,
		0,
	)
	if err != nil {
		panic(err)
	}
	for msg := range pcsmer.Messages() {
		payload := &debeziumcsmr.SamplePayload{}
		if err := json.Unmarshal(msg.Value, payload); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Before     ", cmp.Diff(payload.Payload.Before, &debeziumcsmr.Sample{}))
		fmt.Println("After      ", cmp.Diff(payload.Payload.After, &debeziumcsmr.Sample{}))
		fmt.Println("Transaction", cmp.Diff(payload.Payload.Transaction, &debeziumcsmr.Transaction{}))
	}
}
