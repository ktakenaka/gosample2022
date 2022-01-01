package viewer

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/domain/repository"
)

type SampleViewer interface {
	List(ctx context.Context) (models.SampleSlice, error)
}

type sampleViewer struct {
	getReadFunc repository.DBReadFunc
}

func NewSampleViewer(getReadFunc repository.DBReadFunc) SampleViewer {
	return &sampleViewer{getReadFunc: getReadFunc}
}

func (sv *sampleViewer) List(ctx context.Context) (models.SampleSlice, error) {
	db := sv.getReadFunc()
	return models.Samples().All(ctx, db)
}
