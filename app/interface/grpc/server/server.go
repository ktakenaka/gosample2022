package server

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/domain/repository"
	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
	pkgNotifier "github.com/ktakenaka/gosample2022/app/pkg/notifier"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	ucUpdater "github.com/ktakenaka/gosample2022/app/usecase/updater"
	ucViewer "github.com/ktakenaka/gosample2022/app/usecase/viewer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	samplePb.UnimplementedSampleServer

	ntfr pkgNotifier.Notifier

	updater ucUpdater.Updater
	viewer  ucViewer.Viewer
}

func NewServer(read repository.DBReadFunc, write repository.DBWriteFunc, ntfr pkgNotifier.Notifier) *server {
	srv := &server{
		ntfr:    ntfr,
		updater: ucUpdater.NewUpdater(write),
		viewer:  ucViewer.NewViewer(read),
	}
	return srv
}

func (s *server) getCurrentOffice(ctx context.Context) (context.Context, *models.Office, error) {
	// TODO: Implement authN logic
	currentOffice, err := s.viewer.OfficeOne(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	ctx = pkgNotifier.NewPersonContext(ctx, ulid.ULID(currentOffice.ID), ulid.ULID(currentOffice.ID))
	return ctx, currentOffice, nil
}

func (s *server) notifyError(ctx context.Context, err error) error {
	// TODO: Implement the logic
	s.ntfr.ErrorWithExtrasAndContext(ctx, pkgNotifier.ERR, err, nil)
	return status.Errorf(codes.Internal, err.Error())
}
