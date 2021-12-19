package server

import (
	"github.com/ktakenaka/gosample2022/app/domain/repository"
	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
	"github.com/ktakenaka/gosample2022/app/usecase/updater"
	"github.com/ktakenaka/gosample2022/app/usecase/viewer"
)

type server struct {
	samplePb.UnimplementedSampleServer

	sampleUpdater updater.SampleUpdater
	sampleViewer  viewer.SampleViewer
}

func NewServer(read repository.DBReadFunc, write repository.DBWriteFunc) *server {
	srv := &server{
		sampleUpdater: updater.NewSampleUpdater(write),
		sampleViewer:  viewer.NewSampleViewer(read),
	}
	return srv
}
