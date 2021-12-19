package server

import (
	"context"
	"log"

	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
)

// LisaSamples implements samplePb.LisaSamples
func (s *server) ListSamples(ctx context.Context, in *samplePb.Request) (*samplePb.Response, error) {
	samples, err := s.sampleViewer.List(ctx)
	log.Printf("%v\n", samples)
	return &samplePb.Response{}, err
}
