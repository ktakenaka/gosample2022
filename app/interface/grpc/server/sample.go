package server

import (
	"context"

	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
)

// SampleList implements samplePb.SampleList
func (s *server) SampleList(ctx context.Context, in *samplePb.ListRequest) (*samplePb.ListResponse, error) {
	office, err := s.getCurrentOffice(ctx)
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
		// TODO: Implement with all of fields
		pbSamples[i] = &samplePb.OneSample{
			Id:       uint32(samples[i].ID),
			Biid:     samples[i].Biid,
			Code:     samples[i].Code,
			Category: string(samples[i].Category),
		}
	}
	return &samplePb.ListResponse{Values: pbSamples}, nil
}
