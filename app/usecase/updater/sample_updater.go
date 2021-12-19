package updater

import "github.com/ktakenaka/gosample2022/app/domain/repository"

type SampleUpdater interface{}

type sampleUpdater struct {
	getWriteFunc repository.DBWriteFunc
}

func NewSampleUpdater(getWriteFunc repository.DBWriteFunc) SampleUpdater {
	return &sampleUpdater{getWriteFunc: getWriteFunc}
}
