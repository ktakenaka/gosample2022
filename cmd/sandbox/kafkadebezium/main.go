package main

import (
	"context"

	"github.com/google/uuid"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/interface/infrastructure"
	"github.com/ktakenaka/gosample2022/app/pkg/historydate"
	"github.com/ktakenaka/gosample2022/app/usecase"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/mysql"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	cfg, _       = config.Initialize()
	ctx          = context.Background()
	validFrom, _ = historydate.ParseDate("2021-08-02")
	maxDate, _   = historydate.ParseDate("9999-12-31")
)

func init() {
	boil.DebugMode = true
}

func main() {
	db, task, _ := mysql.Init(ctx, cfg)
	defer task.Shutdown(ctx)

	office := models.Offices().OneP(ctx, db)

	i := usecase.NewInteractor(&infrastructure.Provider{DB: db})
	sample, err := i.SampleCreate(ctx, &usecase.Office{Office: office}, &usecase.BiTemporalSampleRequest{
		Code:      uuid.New().String()[:9],
		Category:  models.SamplesCategoryMedium,
		Amount:    decimal.New(1234, -2),
		ValidFrom: validFrom,
		ValidTo:   maxDate,
	})
	if err != nil {
		panic(err)
	}

	err = i.SampleAddFirst(ctx, &usecase.Office{Office: office}, &usecase.BiTemporalSampleRequest{
		Biid:      sample.Biid,
		Code:      sample.Code,
		Category:  sample.Category,
		Amount:    decimal.New(12345, -3),
		ValidFrom: validFrom,
		ValidTo:   maxDate,
	})
	if err != nil {
		panic(err)
	}
}
