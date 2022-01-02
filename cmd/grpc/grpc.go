package grpc

import (
	"context"

	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
	"github.com/ktakenaka/gosample2022/app/interface/grpc/server"
	"github.com/ktakenaka/gosample2022/cmd/shutdown"
	"github.com/ktakenaka/gosample2022/cmd/tmanager"
	pkggrpc "google.golang.org/grpc"
)

func New(tm *tmanager.TManager) (*pkggrpc.Server, shutdown.Task) {
	srv := server.NewServer(tm.ReadDBFactory, tm.WriteDBFactory)

	s := pkggrpc.NewServer()
	samplePb.RegisterSampleServer(s, srv)

	task := func(ctx context.Context) error {
		s.Stop()
		return nil
	}
	return s, task
}
