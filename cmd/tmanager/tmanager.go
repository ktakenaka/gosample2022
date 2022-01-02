package tmanager

import "github.com/ktakenaka/gosample2022/app/domain/repository"

// TManager: Task Manager
type TManager struct {
	WriteDBFactory repository.DBWriteFunc
	ReadDBFactory  repository.DBReadFunc
}
