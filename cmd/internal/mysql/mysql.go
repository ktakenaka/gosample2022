package mysql

import (
	"context"
	"database/sql"

	"github.com/ktakenaka/gosample2022/app/config"
	"github.com/ktakenaka/gosample2022/app/interface/infrastructure"
	"github.com/ktakenaka/gosample2022/app/pkg/dbresolver"
	"github.com/ktakenaka/gosample2022/cmd/internal/shutdown"
	infraDB "github.com/ktakenaka/gosample2022/infra/database"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/multierr"
)

type task struct {
	write *sql.DB
	read  *sql.DB
}

func (t *task) Name() string {
	return "mysql"
}

func (t *task) Shutdown(ctx context.Context) (err error) {
	if e := t.write.Close(); e != nil {
		err = multierr.Append(err, e)
	}
	if e := t.read.Close(); e != nil {
		err = multierr.Append(err, e)
	}
	return err
}

func Init(ctx context.Context, cfg *config.Config) (infrastructure.DB, shutdown.Task, error) {
	write, err := infraDB.New(cfg.DB.Write)
	if err != nil {
		return nil, nil, err
	}
	if err := write.Ping(); err != nil {
		return nil, nil, err
	}

	read, err := infraDB.New(cfg.DB.Read)
	if err != nil {
		write.Close()
		return nil, nil, err
	}
	if err := read.Ping(); err != nil {
		write.Close()
		return nil, nil, err
	}

	boil.DebugMode = true
	db := dbresolver.New(write, read)
	boil.SetDB(db)

	return db, &task{write: write, read: read}, nil
}
