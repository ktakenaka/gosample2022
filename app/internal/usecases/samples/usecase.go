package sample

import (
	"context"

	"github.com/ktakenaka/gomsx/app/domain/samples"
	"github.com/ktakenaka/gomsx/app/internal/models/v1.0/dao"
	"github.com/ktakenaka/gomsx/app/internal/models/v1.0/entities"
)

type UseCase interface {
	ListSamples(ctx context.Context) ([]*entities.Sample, error)
}

type interactor struct {
	txmFact dao.TxManagerFactory
}

func NewUseCase(ctx context.Context, txmFact dao.TxManagerFactory) UseCase {
	return &interactor{txmFact: txmFact}
}

func (u *interactor) ListSamples(ctx context.Context) ([]*entities.Sample, error) {
	// TODO: when integrating office, replace "1"
	txm, err := u.txmFact(ctx, 1)
	if err != nil {
		return nil, err
	}

	srv := samples.NewService(txm)
	list, err := srv.GetAllSamples(ctx)
	if err != nil {
		return nil, err
	}

	return list, nil
}
