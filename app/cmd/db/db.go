package db

import (
	"context"

	"github.com/ktakenaka/gomsx/app/config"
	"github.com/ktakenaka/gomsx/app/pkg/sqls"
)

type Task struct {
	conf *config.DB
	conn *sqls.DB
}

func Initialize(ctx context.Context, conf *config.DB) (*Task, error) {
	db, err := sqls.Connect(conf.SqlsConf())
	if err != nil {
		return nil, err
	}

	return &Task{conf: conf, conn: db}, nil
}

func (t *Task) Name() string {
	return "db"
}

func (t *Task) Shutdown(ctx context.Context) error {
	return t.conn.Close()
}
