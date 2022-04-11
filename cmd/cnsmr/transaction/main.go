package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ktakenaka/gosample2022/app/interface/infrastructure"
	"github.com/ktakenaka/gosample2022/app/pkg/debeziumcsmr"
	"github.com/ktakenaka/gosample2022/app/usecase"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/kafkaclient"
	"github.com/ktakenaka/gosample2022/cmd/internal/mysql"
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

	db, task, _ := mysql.Init(ctx, cfg)
	defer task.Shutdown(ctx)

	interacter := usecase.NewInteractor(&infrastructure.Provider{Redis: redisClient, DB: db})

	csmer, _ := kafka.NewConsumer(kafkaClient)
	pcsmer, err := csmer.ConsumePartition("gosample2022_dbserver.transaction", 0, 0)
	if err != nil {
		panic(err)
	}
	for msg := range pcsmer.Messages() {
		payload := &debeziumcsmr.TransactionPayload{}
		if err := json.Unmarshal(msg.Value, payload); err != nil {
			fmt.Println(err)
			continue
		}

		if payload.Payload.Status != debeziumcsmr.TransactionStatusEnd {
			continue
		}
		if payload.Payload.EventCount == 0 {
			continue
		}

		if err := redisClient.Set(ctx, debeziumcsmr.RedisKeyCount(payload.Payload.ID), payload.Payload.EventCount, 0).Err(); err != nil {
			fmt.Println(err)
			continue
		}

		samples := []*usecase.SampleCopy{}
		if err := redisClient.SMembers(ctx, debeziumcsmr.RedisKeyRecords(payload.Payload.ID)).ScanSlice(&samples); err != nil {
			fmt.Println(err)
			continue
		}

		if int(payload.Payload.EventCount) != len(samples) {
			continue
		}

		if err := interacter.SyncSamples(ctx, payload.Payload.ID, samples); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
