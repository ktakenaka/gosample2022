package testdb

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"

	"github.com/ktakenaka/gosample2022/app/domain/repository"
	"github.com/ktakenaka/gosample2022/infra/database"
)

const (
	dbName = "gosample2022_test"
)

var (
	cfg = &database.Config{
		User:     "root",
		Password: "root",
		Host:     "mysql",
		DBName:   dbName,
		Options:  map[string]string{"foreign_key_checks": "0"},
	}
)

func init() {
	txdb.Register("txdb", "mysql", database.ConnStr(cfg))
}

func GetDB() (db repository.DB, release func()) {
	// TODO: no need to use uuid
	sqlDB, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		panic(err)
	}

	release = func() {
		if err := sqlDB.Close(); err != nil {
			panic(err)
		}
	}

	return sqlDB, release
}
