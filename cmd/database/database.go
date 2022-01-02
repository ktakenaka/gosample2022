package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ktakenaka/gosample2022/app/domain/repository"
	"github.com/ktakenaka/gosample2022/cmd/tmanager"
	infraDB "github.com/ktakenaka/gosample2022/infra/database"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type task struct {
	dbWrite *sql.DB
	dbRead *sql.DB
}

func (t *task) Name() string {
	return "database"
}

func (t *task) Shutdown(ctx context.Context) (err error) {
	if e := t.dbRead.Close(); e != nil {
		err = e
	}
	if e := t.dbWrite.Close(); e != nil {
		if err != nil {
			err = fmt.Errorf("%w, %w", err, e)
			return
		}
		err = e
		return
	}
	return nil
}

func Init(writeCfg, readCfg *infraDB.Config) (
	read repository.DBReadFunc,
	write repository.DBWriteFunc,
	t tmanager.Task,
	err error,
) {
	writeDB, err := infraDB.New(writeCfg)
	if err != nil {
		return
	}
	write = func() repository.WriteExecutor {
		return writeDB
	}

	readDB, err := infraDB.New(readCfg)
	if err != nil {
		return
	}
	read = func() repository.ReadExecutor {
		return readDB
	}

	t = &task{dbWrite: writeDB, dbRead: readDB}
	boil.DebugMode = true

	return read, write, t, nil
}
