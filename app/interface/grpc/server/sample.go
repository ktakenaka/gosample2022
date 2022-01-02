package server

import (
	"context"

	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
	"github.com/ktakenaka/gosample2022/app/pkg/notifier"
)

// LisaSamples implements samplePb.LisaSamples
func (s *server) ListSamples(ctx context.Context, in *samplePb.Request) (*samplePb.Response, error) {
	_, err := s.sampleViewer.List(ctx)
	if err == nil {
		s.ntfr.Message(notifier.INFO, "success")
	}
	return &samplePb.Response{}, err
}
