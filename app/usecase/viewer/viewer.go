package viewer

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/domain/repository"
)

type Viewer interface {
	CurrentOffice(ctx context.Context) (*models.Office, error)

	SampleList(ctx context.Context, office *models.Office) (models.SampleSlice, error)
}

type viewer struct {
	getReadFunc repository.DBReadFunc
}

func NewViewer(getReadFunc repository.DBReadFunc) Viewer {
	return &viewer{getReadFunc: getReadFunc}
}

func (v *viewer) CurrentOffice(ctx context.Context) (*models.Office, error) {
	// TODO: Implement auth logic
	return models.Offices().One(ctx, v.getReadFunc())
}

func (v *viewer) SampleList(ctx context.Context, office *models.Office) (models.SampleSlice, error) {
	db := v.getReadFunc()
	return office.Samples().All(ctx, db)
}
