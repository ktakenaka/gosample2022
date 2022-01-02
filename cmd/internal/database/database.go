package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ktakenaka/gosample2022/app/config"
	"github.com/ktakenaka/gosample2022/app/domain/repository"
	"github.com/ktakenaka/gosample2022/cmd/internal/shutdown"
	infraDB "github.com/ktakenaka/gosample2022/infra/database"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type task struct {
	writeDB *sql.DB
	readDB  *sql.DB
}

func (t *task) Name() string {
	return "database"
}

func (t *task) Shutdown(ctx context.Context) (err error) {
	if err1 := t.writeDB.Close(); err1 != nil {
		err = err1
	}
	if err2 := t.readDB.Close(); err2 != nil {
		if err != nil {
			err = fmt.Errorf("%w, %s", err, err2)
			return
		}
		err = err2
	}
	return err
}

func Init(cfg *config.Config) (
	read repository.DBReadFunc,
	write repository.DBWriteFunc,
	t shutdown.Task,
	err error,
) {
	writeDB, err := infraDB.New(cfg.DB.Write)
	if err != nil {
		return
	}
	write = func() repository.WriteExecutor {
		return writeDB
	}

	readDB, err := infraDB.New(cfg.DB.Read)
	if err != nil {
		return
	}
	read = func() repository.ReadExecutor {
		return readDB
	}

	t = &task{writeDB: writeDB, readDB: readDB}
	boil.DebugMode = true

	return read, write, t, nil
}
