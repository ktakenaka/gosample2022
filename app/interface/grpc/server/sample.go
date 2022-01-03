package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
	"github.com/ktakenaka/gosample2022/app/pkg/notifier"
)

// LisaSamples implements samplePb.LisaSamples
func (s *server) ListSamples(ctx context.Context, in *samplePb.Request) (*samplePb.Response, error) {
	_, err := s.sampleViewer.List(ctx)

	if err != nil {
		// Temporary implementation
		s.ntfr.Message(notifier.ERR, err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &samplePb.Response{}, err
}
