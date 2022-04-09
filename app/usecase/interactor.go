package usecase

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/domain/repository"
)

type Interactor interface {
	OfficeOne(ctx context.Context, officeID string) (*models.Office, error)

	SampleList(ctx context.Context, office *models.Office) (models.SampleSlice, error)
}

type interactor struct {
	db    repository.DB
	redis repository.Redis
}

func NewInteractor(db repository.DB, redis repository.Redis) Interactor {
	return &interactor{db: db, redis: redis}
}

func (i *interactor) OfficeOne(ctx context.Context, officeID string) (*models.Office, error) {
	return models.Offices(models.OfficeWhere.ID.EQ(officeID)).One(ctx, i.db)
}

func (i *interactor) SampleList(ctx context.Context, office *models.Office) (models.SampleSlice, error) {
	return office.Samples().All(ctx, i.db)
}
