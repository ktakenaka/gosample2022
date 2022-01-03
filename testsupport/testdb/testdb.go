package testdb

import (
	"context"
	"fmt"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/domain/repository"
	"github.com/ktakenaka/gosample2022/infra/database"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	protectedTables = []interface{}{"schema_migrations", "users", "offices", "office_users"}
)

func GetFuncs() (read repository.DBReadFunc, write repository.DBWriteFunc, release func()) {
	// いったん一つのDBだけで使えるようにする
	// TODO: 8並列で別のDBに接続して使えるようにする
	dbname := "gosample2022_test1"
	boil.DebugMode = true

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
		User:     "test_reader",
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

func cleanup(db repository.WriteExecutor, dbname string) error {
	var raws []*struct{ Name string }

	err := models.NewQuery(
		qm.Select("table_name AS 'name'"),
		qm.From("information_schema.tables"),
		qm.Where("table_schema=?", dbname),
		qm.WhereNotIn("table_name NOT IN ?", protectedTables...),
	).Bind(context.Background(), db, &raws)
	if err != nil {
		return err
	}

	var groupedErr error
	for i := range raws {
		// 他に方法が思いつかないので、SQL Injection対策はしない
		_, err = queries.Raw("TRUNCATE TABLE " + raws[i].Name).Exec(db)
		if err != nil && groupedErr == nil {
			groupedErr = err
			continue
		}
		if err != nil {
			groupedErr = fmt.Errorf("%s: %s", groupedErr, err)
		}
	}
	return groupedErr
}
