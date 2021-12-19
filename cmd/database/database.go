package database

import (
	"context"
	"fmt"

	"github.com/ktakenaka/gosample2022/app/domain/repository"
	"github.com/ktakenaka/gosample2022/cmd/shutdown"
	infraDB "github.com/ktakenaka/gosample2022/infra/database"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Init() (read repository.DBReadFunc, write repository.DBWriteFunc, task shutdown.Task, err error) {
	// TODO: Pass app config from outside
	writeCfg := &infraDB.Config{
		User:     "writer",
		Password: "writer_password",
		Host:     "db",
		DBName:   "gosample2022_development",
	}
	writeDB, err := infraDB.New(writeCfg)
	if err != nil {
		return
	}
	write = func() repository.WriteExecutor {
		return writeDB
	}

	readCfg := &infraDB.Config{
		User:     "reader",
		Password: "reader_password",
		Host:     "db",
		DBName:   "gosample2022_development",
	}
	readDB, err := infraDB.New(readCfg)
	if err != nil {
		return
	}
	read = func() repository.ReadExecutor {
		return readDB
	}

	task = func(ctx context.Context) error {
		var err error
		if err1 := writeDB.Close(); err1 != nil {
			err = fmt.Errorf("write db close: %w", err1)
		}
		if err2 := readDB.Close(); err2 != nil {
			err = fmt.Errorf("read db close: %w", err2)
		}

		return err
	}

	boil.DebugMode = true

	return read, write, task, nil
}
