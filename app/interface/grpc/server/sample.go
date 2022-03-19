package server

import (
	"context"

	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
)

// SampleList implements samplePb.SampleList
func (s *server) SampleList(ctx context.Context, in *samplePb.ListRequest) (*samplePb.ListResponse, error) {
	ctx, office, err := s.getCurrentOffice(ctx)
	if err != nil {
		err = s.notifyError(ctx, err)
		return nil, err
	}

	samples, err := s.interactor.SampleList(ctx, office)
	if err != nil {
		err = s.notifyError(ctx, err)
		return nil, err
	}

	pbSamples := make([]*samplePb.OneSample, len(samples))
	for i := range samples {
		pbSamples[i] = &samplePb.OneSample{
			Id:       ulid.ULID(samples[i].ID).String(),
			Title:    samples[i].Title,
			Category: samples[i].Category,
			Memo:     samples[i].Memo,
			Date:     samples[i].Date.String(),
			Amount:   samples[i].Amount.String(),
		}
	}
	return &samplePb.ListResponse{Values: pbSamples}, nil
}
