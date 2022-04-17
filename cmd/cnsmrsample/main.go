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
		event := &maxwellcsmr.SampleEvent{}
		if err := json.Unmarshal(msg.Value, event); err != nil {
			panic(err)
		}

		if event.Database != targetDB || event.Table != targetTable {
			continue
		}

		sample := &models.SampleCopy{
			ID:        event.Data.ID,
			Biid:      event.Data.Biid,
			OfficeID:  event.Data.OfficeID,
			Code:      event.Data.Code,
			Category:  models.SampleCopiesCategory(event.Data.Category),
			Amount:    decimal.Decimal(event.Data.Amount),
			ValidFrom: time.Time(event.Data.ValidFrom),
			ValidTo:   time.Time(event.Data.ValidTo),
			CreatedAt: time.Time(event.Data.CreatedAt),
			DeletedAt: null.Time(event.Data.DeletedAt),
			Version:   event.Data.Version,
		}

		if err := redisClient.SAdd(ctx, maxwellcsmr.CacheKey(event.XID), &usecase.SampleCopy{SampleCopy: sample}).Err(); err != nil {
			panic(err)
		}

		if !event.Commit {
			continue
		}

		samples := []*usecase.SampleCopy{}
		if err := redisClient.SMembers(ctx, maxwellcsmr.CacheKey(event.XID)).ScanSlice(&samples); err != nil {
			panic(err)
		}

		if err := interactor.SyncSamples(ctx, maxwellcsmr.CacheKey(event.XID), samples); err != nil {
			panic(err)
		}
	}
}
