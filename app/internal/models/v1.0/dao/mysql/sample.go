package mysql

import (
	"context"
	"github.com/ktakenaka/gomsx/app/internal/models/v1.0/entities"
	"github.com/ktakenaka/gomsx/app/internal/models/v1.0/models"
)

type sampleDAO struct {
	daoBase
}

func (dao *sampleDAO) FetchAllSamples(ctx context.Context) ([]*entities.Sample, error) {
	list, err := models.Samples().All(ctx, dao.executor)
	if err != nil {
		return nil, err
	}

	samples := make([]*entities.Sample, len(list))
	for i, s := range list {
		samples[i] = &entities.Sample{Sample: s}
	}
	return samples, nil
}
