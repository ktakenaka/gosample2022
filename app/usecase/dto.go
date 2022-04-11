package usecase

import (
	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/pkg/historydate"
	"github.com/shopspring/decimal"
)

type BiTemporalSampleRequest struct {
	ID        uint
	Biid      string
	Code      string
	Category  models.SamplesCategory
	Amount    decimal.Decimal
	ValidFrom historydate.Date
	ValidTo   historydate.Date
}
