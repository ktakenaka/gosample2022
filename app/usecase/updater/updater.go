package updater

import "github.com/ktakenaka/gosample2022/app/domain/repository"

type Updater interface{}

type updater struct {
	getWriteFunc repository.DBWriteFunc
}

func NewUpdater(getWriteFunc repository.DBWriteFunc) Updater {
	return &updater{getWriteFunc: getWriteFunc}
}
