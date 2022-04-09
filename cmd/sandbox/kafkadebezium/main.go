package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/interface/infrastructure"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/kafkaclient"
	"github.com/ktakenaka/gosample2022/cmd/internal/mysql"
	"github.com/ktakenaka/gosample2022/infra/kafka"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	cfg, _ = config.Initialize()
	ctx    = context.Background()
)

func init() {
	boil.DebugMode = true
}

func main() {
	db, task, _ := mysql.Init(ctx, cfg)
	defer task.Shutdown(ctx)

	kafkaClient, task, _ := kafkaclient.Init(ctx, cfg)
	defer task.Shutdown(ctx)

	go startConsumer(kafkaClient)

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer tx.Commit()
	defer time.Sleep(2 * time.Second)

	office := &models.Office{ID: ulid.MustNew(), Name: "debezium sample"}
	if err := office.Insert(ctx, tx, boil.Infer()); err != nil {
		panic(err)
	}
	if err := office.AddSamples(ctx, tx, true, &models.Sample{
		Biid:      ulid.MustNew(),
		Code:      "code",
		Category:  models.SamplesCategorySmall,
		Amount:    decimal.NewFromFloat(1.2),
		ValidFrom: time.Now(),
	}); err != nil {
		panic(err)
	}
}

func startConsumer(client infrastructure.Kafka) {
	csmer, _ := kafka.NewConsumer(client)
	pcsmer, err := csmer.ConsumePartition(
		"gosample2022_dbserver.gosample2022_development.samples", 0, 0,
	)
	if err != nil {
		panic(err)
	}
	for msg := range pcsmer.Messages() {
		fmt.Println("Key", string(msg.Key))
		fmt.Println("Value", string(msg.Value))
	}
}
