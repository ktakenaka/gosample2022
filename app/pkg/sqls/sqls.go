package sqls

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	defaultMaxIdleConns    = 5
	defaultMaxOpenConns    = 10
	defaultConnMaxLifetime = 20 * time.Second
)

type DB struct {
	*sqlx.DB
}

func Connect(conf *Config) (*DB, error) {
	db, err := sqlx.Connect(conf.Driver, conf.ToConnString())
	// TODO: retry severail times
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	if conf.MaxIdleConns == 0 {
		db.SetMaxIdleConns(defaultMaxIdleConns)
	} else {
		db.SetMaxIdleConns(conf.MaxIdleConns)
	}

	if conf.MaxOpenConns == 0 {
		db.SetMaxOpenConns(defaultMaxOpenConns)
	} else {
		db.SetMaxOpenConns(conf.MaxOpenConns)
	}

	if conf.ConnMaxLifetime == 0 {
		db.SetConnMaxLifetime(defaultConnMaxLifetime)
	} else {
		db.SetConnMaxLifetime(conf.ConnMaxLifetime)
	}

	return &DB{db}, nil
}
