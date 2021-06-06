package mysql

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/ktakenaka/gomsx/app/internal/models/v1.0/dao"
	"github.com/ktakenaka/gomsx/app/pkg/sqls"
)

func TxManagerFactory(db *sqls.DB) dao.TxManagerFactory {
	return func(ctx context.Context, officeID uint32) (dao.TxManager, error) {
		return &mysqlTxManager{db: db, daoBase: daoBase{db}}, nil
	}
}

type daoBase struct {
	executor boil.ContextExecutor
}

func (d daoBase) SampleDAO() dao.Sample {
	return &sampleDAO{daoBase{d.executor}}
}

type mysqlTxManager struct {
	daoBase
	db *sqls.DB
}

func (mtm *mysqlTxManager) BeginTx(ctx context.Context, opts *sql.TxOptions) (dao.Tx, error) {
	tx, err := mtm.db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &mysqlTx{tx: tx, daoBase: daoBase{tx}}, nil
}

type mysqlTx struct {
	daoBase
	tx *sqls.Tx
}

func (mt *mysqlTx) Commit(ctx context.Context) error {
	return mt.tx.Commit()
}

func (mt *mysqlTx) Rollback(ctx context.Context) error {
	return mt.tx.Rollback()
}
