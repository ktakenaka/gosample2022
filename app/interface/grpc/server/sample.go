package server

import (
	"context"

	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
	"github.com/ktakenaka/gosample2022/app/pkg/notifier"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LisaSamples implements samplePb.LisaSamples
func (s *server) ListSamples(ctx context.Context, in *samplePb.Request) (*samplePb.Response, error) {
	samples, err := s.sampleViewer.List(ctx)

	if err != nil {
		// Temporary implementation
		s.ntfr.Message(notifier.ERR, err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
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
	return &samplePb.Response{}, err
}
