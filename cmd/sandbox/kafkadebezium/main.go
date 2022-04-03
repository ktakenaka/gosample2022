package main

import (
	"context"
	"encoding/hex"
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

	officeID, _ = hex.DecodeString("01563E3AB5D3D6764C61EFB99302BD5B")
)

func main() {
	db, task, _ := mysql.Init(ctx, cfg)
	defer task.Shutdown(ctx)

	office := models.Office{ID: officeID, Name: "sample"}
	office.Upsert(ctx, db, boil.Infer(), boil.Infer())
	err := office.AddSamples(ctx, db, true, &models.Sample{
		ID:       ulid.MustNew(),
		Title:    "title",
		Category: models.SamplesCategorySmall,
		Memo:     "memo",
		Date:     time.Now(),
		Amount:   decimal.NewFromFloat(1.2),
	})
	fmt.Println(err)
}
