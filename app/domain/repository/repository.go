package repository

import "github.com/volatiletech/sqlboiler/v4/boil"

// Transaction
type ReadExecutor interface {
	boil.ContextExecutor
}
type WriteExecutor interface {
	boil.ContextExecutor
	boil.Beginner
}
type DBWriteFunc func() WriteExecutor
type DBReadFunc func() ReadExecutor