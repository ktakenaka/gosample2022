package dbresolver

import (
	"context"
	"database/sql"
)

type DB struct {
	primary *sql.DB
	replica *sql.DB
}

func New(primary, replica *sql.DB) *DB {
	return &DB{primary: primary, replica: replica}
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.primary.Exec(query, args...)
}

func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.replica.Query(query, args...)
}

func (db *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.replica.QueryRow(query, args...)
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.primary.ExecContext(ctx, query, args...)
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.replica.QueryContext(ctx, query, args...)
}

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return db.replica.QueryRowContext(ctx, query, args...)
}

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return db.primary.BeginTx(ctx, opts)
}
