package server

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
	"github.com/ktakenaka/gosample2022/app/registry"
	"github.com/ktakenaka/gosample2022/app/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	samplePb.UnimplementedSampleServer

	interactor usecase.Interactor
}

func NewServer(provider *registry.Provider) *server {
	srv := &server{interactor: usecase.NewInteractor(provider.DB, provider.Redis)}
	return srv
}

func (s *server) getCurrentOffice(ctx context.Context) (*models.Office, error) {
	// TODO: Implement authN logic
	currentOffice, err := s.interactor.OfficeOne(ctx, "")
	if err != nil {
		return nil, err
	}
	// TODO: Add Person context for error tracing (Rollbar)
	return currentOffice, nil
}

func (s *server) notifyError(ctx context.Context, err error) error {
	// TODO: Implement the logic
	// s.ntfr.ErrorWithExtrasAndContext(ctx, pkgNotifier.ERR, err, nil)
	return status.Errorf(codes.Internal, err.Error())
}
