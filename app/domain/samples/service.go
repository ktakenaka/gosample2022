package samples

import (
	"context"

	"github.com/ktakenaka/gomsx/app/internal/models/v1.0/dao"
	"github.com/ktakenaka/gomsx/app/internal/models/v1.0/entities"
)

type Service interface {
	GetAllSamples(ctx context.Context) ([]*entities.Sample, error)
}

type service struct {
	txm dao.TxManager
}

func NewService(txm dao.TxManager) Service {
	return &service{txm: txm}
}

func (s *service) GetAllSamples(ctx context.Context) ([]*entities.Sample, error) {
	return s.txm.SampleDAO().FetchAllSamples(ctx)
}
