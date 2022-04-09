package usecase

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/interface/infrastructure"
)

type Interactor interface {
	OfficeOne(ctx context.Context, officeID string) (*models.Office, error)

	SampleList(ctx context.Context, office *models.Office) (models.SampleSlice, error)
}

type interactor struct {
	p *infrastructure.Provider
}

func NewInteractor(p *infrastructure.Provider) Interactor {
	return &interactor{p: p}
}

func (i *interactor) OfficeOne(ctx context.Context, officeID string) (*models.Office, error) {
	return models.Offices(models.OfficeWhere.ID.EQ(officeID)).One(ctx, i.p.DB)
}

func (i *interactor) SampleList(ctx context.Context, office *models.Office) (models.SampleSlice, error) {
	return office.Samples().All(ctx, i.p.DB)
}
