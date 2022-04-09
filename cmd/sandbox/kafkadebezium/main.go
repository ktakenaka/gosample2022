package main

import (
	"context"
	"fmt"
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

	officeID = ulid.MustNew()
)

func main() {
	db, task, _ := mysql.Init(ctx, cfg)
	defer task.Shutdown(ctx)

	office := &models.Office{ID: officeID, Name: "sample"}
	office.Upsert(ctx, db, boil.Infer(), boil.Infer())
	err := office.AddSamples(ctx, db, true, &models.Sample{
		Biid:      ulid.MustNew(),
		Code:      "code",
		Category:  models.SamplesCategorySmall,
		Amount:    decimal.NewFromFloat(1.2),
		ValidFrom: time.Now(),
	})
	fmt.Println(err)
}
