package dao

import (
	"context"
	"database/sql"

	"github.com/ktakenaka/gomsx/app/internal/models/v1.0/entities"
)

type TxManagerFactory func(ctx context.Context, officeID uint32) (TxManager, error)

type TxManager interface {
	Factory
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error)
}

type Tx interface {
	Factory
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type Factory interface {
	SampleDAO() Sample
}

type Sample interface {
	FetchAllSamples(ctx context.Context) ([]*entities.Sample, error)
}
