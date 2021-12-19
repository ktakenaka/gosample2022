package grpc

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/domain/repository"
	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
	"github.com/ktakenaka/gosample2022/app/interface/grpc/server"
	"github.com/ktakenaka/gosample2022/cmd/shutdown"
	pkggrpc "google.golang.org/grpc"
)

func New(read repository.DBReadFunc, write repository.DBWriteFunc) (*pkggrpc.Server, shutdown.Task) {
	srv := server.NewServer(read, write)

	s := pkggrpc.NewServer()
	samplePb.RegisterSampleServer(s, srv)

	task := func(ctx context.Context) error {
		s.Stop()
		return nil
	}
	return s, task
}
