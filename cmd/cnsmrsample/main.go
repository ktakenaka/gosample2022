package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/interface/infrastructure"
	"github.com/ktakenaka/gosample2022/app/pkg/maxwellcsmr"
	"github.com/ktakenaka/gosample2022/app/usecase"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/kafkaclient"
	"github.com/ktakenaka/gosample2022/cmd/internal/mysql"
	"github.com/ktakenaka/gosample2022/cmd/internal/redis"
	"github.com/ktakenaka/gosample2022/infra/kafka"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
)

const (
	topic     = "gosample2022_maxwell"
	partition = 0

	targetDB    = "gosample2022_development"
	targetTable = "samples"
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

	interactor := usecase.NewInteractor(&infrastructure.Provider{Redis: redisClient, DB: db})

	csmer, _ := kafka.NewConsumer(kafkaClient)

	offset := int64(0)
	pcsmer, err := csmer.ConsumePartition(topic, partition, offset)
	if err != nil {
		panic(err)
	}
	for msg := range pcsmer.Messages() {
		event := &maxwellcsmr.Event{}
		if err := json.Unmarshal(msg.Value, event); err != nil {
			panic(err)
		}
		if event.Database != targetDB || event.Table != targetTable {
			continue
		}

		eventSample := &maxwellcsmr.Sample{}
		if err := json.Unmarshal(event.Data, eventSample); err != nil {
			panic(err)
		}

		sample := &models.SampleCopy{
			ID:        eventSample.ID,
			Biid:      eventSample.Biid,
			OfficeID:  eventSample.OfficeID,
			Code:      eventSample.Code,
			Category:  models.SampleCopiesCategory(eventSample.Category),
			Amount:    decimal.Decimal(eventSample.Amount),
			ValidFrom: time.Time(eventSample.ValidFrom),
			ValidTo:   time.Time(eventSample.ValidTo),
			CreatedAt: time.Time(eventSample.CreatedAt),
			DeletedAt: null.Time(eventSample.DeletedAt),
			Version:   eventSample.Version,
		}

		if !event.Commit {
			if err := redisClient.SAdd(ctx, maxwellcsmr.CacheKey(event.XID), &usecase.SampleCopy{SampleCopy: sample}).Err(); err != nil {
				panic(err)
			}
			continue
		}

		samples := []*usecase.SampleCopy{}
		if err := redisClient.SMembers(ctx, maxwellcsmr.CacheKey(event.XID)).ScanSlice(&samples); err != nil {
			panic(err)
		}
		samples = append(samples, &usecase.SampleCopy{SampleCopy: sample})

		if err := interactor.SyncSamples(ctx, uint(msg.Offset), samples); err != nil {
			panic(err)
		}

		redisClient.Del(ctx, maxwellcsmr.CacheKey(event.XID)).Err()
	}
}
