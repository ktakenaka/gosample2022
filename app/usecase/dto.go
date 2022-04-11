package usecase

import (
	"encoding/json"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/pkg/historydate"
	"github.com/shopspring/decimal"
)

type Office struct {
	*models.Office
}

type BiTemporalSampleRequest struct {
	ID        uint
	Biid      string
	Code      string
	Category  models.SamplesCategory
	Amount    decimal.Decimal
	ValidFrom historydate.Date
	ValidTo   historydate.Date
}

type Sample struct {
	*models.Sample
}

type SampleCopy struct {
	*models.SampleCopy
}

func (s *SampleCopy) MarshalBinary() (data []byte, err error) {
	return json.Marshal(s)
}

func (s *SampleCopy) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
