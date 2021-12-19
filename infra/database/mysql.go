package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	defaultOptions = map[string]string{
		"collation":         "utf8mb4_bin",
		"charset":           "utf8mb4",
		"parseTime":         "true",
		"interpolateParams": "true",
	}
)

const (
	defaultMaxIdleConns    = 5
	defaultMaxOpenConns    = 10
	defaultConnMaxLifetime = 20 * time.Second

	conn = "%s:%s@tcp(%s)/%s"
)

// Config connection information
type Config struct {
	User            string
	Password        string
	Host            string
	DBName          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration

	Options map[string]string
}

// New connect to db
func New(cfg *Config) (*sql.DB, error) {
	optionsMap := defaultOptions
	if cfg.Options != nil {
		for k, v := range cfg.Options {
			optionsMap[k] = v
		}
	}

	var optionsSlice []string
	for k, v := range optionsMap {
		optionsSlice = append(optionsSlice, k+"="+v)
	}

	connStr := fmt.Sprintf(conn+"?"+strings.Join(optionsSlice, "&"), cfg.User, cfg.Password, cfg.Host, cfg.DBName)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	if cfg.MaxIdleConns == 0 {
		db.SetMaxIdleConns(defaultMaxIdleConns)
	} else {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}

	if cfg.MaxOpenConns == 0 {
		db.SetMaxOpenConns(defaultMaxOpenConns)
	} else {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}

	if cfg.ConnMaxLifetime == 0 {
		db.SetConnMaxLifetime(defaultConnMaxLifetime)
	} else {
		db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}

	return db, nil
}
