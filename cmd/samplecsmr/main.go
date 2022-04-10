package main

import (
	"context"
	"encoding/json"
	"fmt"

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
		hoge := &Hoge{}
		if err := json.Unmarshal(msg.Value, hoge); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("=== ID ===", hoge.Payload.After.Id)
		fmt.Println("Code", hoge.Payload.After.Code)
		fmt.Println("Category", hoge.Payload.After.Category)
		fmt.Println("Amount", hoge.Payload.After.Amount)
		fmt.Println("ValidFrom", hoge.Payload.After.ValidFrom)
		fmt.Println("ValidTo", hoge.Payload.After.ValidTo)
		fmt.Println("CreatedAt", hoge.Payload.After.CreatedAt)
		fmt.Println("DeletedAt", hoge.Payload.After.DeletedAt)
	}
}

type Hoge struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	After DSample `json:"after"`
}

type DSample struct {
	Id        uint                  `json:"id"`
	Biid      string                `json:"biid"`
	OfficeId  string                `json:"office_id"`
	Code      string                `json:"code"`
	Category  string                `json:"category"`
	Amount    debeziumcsmr.Decimal  `json:"amount"`
	ValidFrom debeziumcsmr.Date     `json:"valid_from"`
	ValidTo   debeziumcsmr.Date     `json:"valid_to"`
	CreatedAt debeziumcsmr.Time     `json:"created_at"`
	DeletedAt debeziumcsmr.NullTime `json:"deleted_at"`
}
