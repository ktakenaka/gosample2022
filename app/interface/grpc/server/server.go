package server

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/domain/repository"
	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
	"github.com/ktakenaka/gosample2022/app/pkg/notifier"
	ucUpdater "github.com/ktakenaka/gosample2022/app/usecase/updater"
	ucViewer "github.com/ktakenaka/gosample2022/app/usecase/viewer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	samplePb.UnimplementedSampleServer

	ntfr notifier.Notifier

	updater ucUpdater.Updater
	viewer  ucViewer.Viewer
}

func NewServer(read repository.DBReadFunc, write repository.DBWriteFunc, ntfr notifier.Notifier) *server {
	srv := &server{
		ntfr:    ntfr,
		updater: ucUpdater.NewUpdater(write),
		viewer:  ucViewer.NewViewer(read),
	}
	return srv
}

func (s *server) notifyError(_ context.Context, err error) error {
	// TODO: Implement the logic
	s.ntfr.Message(notifier.DEBUG, err.Error())
	return status.Errorf(codes.Internal, err.Error())
}
