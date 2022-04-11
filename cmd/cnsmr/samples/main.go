package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	pkgRedis "github.com/go-redis/redis/v8"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/interface/infrastructure"
	"github.com/ktakenaka/gosample2022/app/pkg/debeziumcsmr"
	"github.com/ktakenaka/gosample2022/app/usecase"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/kafkaclient"
	"github.com/ktakenaka/gosample2022/cmd/internal/mysql"
	"github.com/ktakenaka/gosample2022/cmd/internal/redis"
	"github.com/ktakenaka/gosample2022/infra/kafka"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
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
	pcsmer, err := csmer.ConsumePartition("gosample2022_dbserver.gosample2022_development.samples", 0, 0)
	if err != nil {
		panic(err)
	}
	for msg := range pcsmer.Messages() {
		payload := &debeziumcsmr.SamplePayload{}
		if err := json.Unmarshal(msg.Value, payload); err != nil {
			fmt.Println(err)
			continue
		}

		sample := &models.Sample{
			ID:        payload.Payload.After.ID,
			Biid:      payload.Payload.After.Biid,
			OfficeID:  payload.Payload.After.OfficeID,
			Code:      payload.Payload.After.Code,
			Category:  models.SamplesCategory(payload.Payload.After.Category),
			Amount:    decimal.Decimal(payload.Payload.After.Amount),
			ValidFrom: time.Time(payload.Payload.After.ValidFrom),
			ValidTo:   time.Time(payload.Payload.After.ValidTo),
			CreatedAt: time.Time(payload.Payload.After.CreatedAt),
			DeletedAt: null.Time(payload.Payload.After.DeletedAt),
		}

		if err := redisClient.SAdd(ctx, debeziumcsmr.RedisKeyRecords(payload.Payload.Transaction.ID), &usecase.Sample{Sample: sample}).Err(); err != nil {
			fmt.Println(err)
			continue
		}

		samples := []*usecase.Sample{}
		if err := redisClient.SMembers(ctx, debeziumcsmr.RedisKeyRecords(payload.Payload.Transaction.ID)).ScanSlice(&samples); err != nil {
			fmt.Println(err)
			continue
		}

		count, err := redisClient.Get(ctx, debeziumcsmr.RedisKeyCount(payload.Payload.Transaction.ID)).Int()
		if err != nil && !errors.Is(err, pkgRedis.Nil) {
			fmt.Println(err)
			continue
		}

		if count != len(samples) {
			continue
		}

		if err := interacter.SyncSamples(ctx, payload.Payload.Transaction.ID, samples); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
