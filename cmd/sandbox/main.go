package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	pkgNotifier "github.com/ktakenaka/gosample2022/app/pkg/notifier"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/notifier"
	"github.com/ktakenaka/gosample2022/infra/database"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func init() {
	boil.DebugMode = true
}

func main() {
	cfg, err := config.Initialize()
	if err != nil {
		panic(err)
	}

	db, err := database.New(cfg.DB.Write)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	office, _ := models.Offices().One(ctx, db)

	id, _ := ulid.GenerateID()
	office.AddSamples(ctx, db, true, &models.Sample{
		ID:       id,
		Title:    "title",
		Category: "small",
		Date:     time.Now(),
		Amount:   decimal.New(123, -2),
	})

	samples, _ := models.Samples().All(ctx, db)
	fmt.Printf("%v\n", samples)

	ntfr, _ := notifier.Init(cfg)
	ntfr.ErrorWithExtrasAndContext(
		pkgNotifier.NewPersonContext(ctx, ulid.ULID(office.ID).String()),
		pkgNotifier.WARN,
		fmt.Errorf("hellow"),
		map[string]interface{}{"user": office},
	)
}
