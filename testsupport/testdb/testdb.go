package testdb

import (
	"context"
	"database/sql"

	"github.com/ktakenaka/gosample2022/app/domain/repository"
	"github.com/ktakenaka/gosample2022/infra/database"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

func GetFuncs() (read repository.DBReadFunc, write repository.DBWriteFunc, release func()) {
	// いったん一つのDBだけで使えるようにする
	// TODO: 8並列で別のDBに接続して使えるようにする
	dbname := "gosample2022_test1"

	// cleanup処理のためにrootで接続する
	writeCfg := &database.Config{
		User:     "root",
		Password: "root",
		Host:     "db",
		DBName:   dbname,
		Options:  map[string]string{"foreign_key_checks": "0"},
	}
	writeDB, err := database.New(writeCfg)
	if err != nil {
		panic(err)
	}
	write = func() repository.WriteExecutor {
		return writeDB
	}

	readCfg := &database.Config{
		User:     "reader",
		Password: "reader_password",
		Host:     "db",
		DBName:   dbname,
		Options:  map[string]string{"foreign_key_checks": "0"},
	}
	readDB, err := database.New(readCfg)
	if err != nil {
		panic(err)
	}
	read = func() repository.ReadExecutor {
		return readDB
	}

	release = func() {
		_ = cleanup(writeDB, dbname)
		_ = writeDB.Close()
		_ = readDB.Close()
	}

	return read, write, release
}

func cleanup(db *sql.DB, dbname string) error {
	var raws []string
	err := queries.
		Raw(
			"SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA=? AND TABLE_NAME != 'schema_migrations'",
			dbname,
		).
		Bind(context.TODO(), db, &raws)
	if err != nil {
		return err
	}

	for i := range raws {
		// 他に方法が思いつかないので、SQL Injection対策はしない
		queries.Raw(
			"DELETE FROM " + raws[i],
		).Exec(db)
	}
	return nil
}
