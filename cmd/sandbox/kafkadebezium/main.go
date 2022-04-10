package main

import (
	"context"
	"time"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/mysql"
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

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer tx.Commit()

	office := &models.Office{ID: ulid.MustNew(), Name: "debezium sample"}
	if err := office.Insert(ctx, tx, boil.Infer()); err != nil {
		panic(err)
	}
	if err := office.AddSamples(ctx, tx, true, &models.Sample{
		Biid:      ulid.MustNew(),
		Code:      "code1",
		Category:  models.SamplesCategoryLarge,
		Amount:    decimal.NewFromFloat(1.2),
		ValidFrom: time.Now(),
	}); err != nil {
		panic(err)
	}
}
