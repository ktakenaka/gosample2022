package grpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/ktakenaka/gosample2022/app/config"
	"github.com/ktakenaka/gosample2022/app/domain/repository"
	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
	"github.com/ktakenaka/gosample2022/app/interface/grpc/server"
	"github.com/ktakenaka/gosample2022/app/pkg/notifier"
	"github.com/ktakenaka/gosample2022/cmd/internal/shutdown"
	pkggrpc "google.golang.org/grpc"
)

type task struct {
	srv *pkggrpc.Server
}

func (t *task) Name() string {
	return "grpc"
}
func (t *task) Shutdown(ctx context.Context) error {
	t.srv.Stop()
	return nil
}

func New(
	cfg *config.Config,
	read repository.DBReadFunc,
	write repository.DBWriteFunc,
	ntfr notifier.Notifier,
) (shutdown.Task, error) {
	srv := server.NewServer(read, write, ntfr)

	s := pkggrpc.NewServer()
	samplePb.RegisterSampleServer(s, srv)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.App.Port))
	if err != nil {
		return nil, err
	}

	go func() {
		if err := s.Serve(lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// TODO: don't panic
			panic(err)
		}
	}()

	return &task{srv: s}, nil
}
