package main

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	cfg, _ = config.Initialize()
	ctx    = context.Background()
)

func init() {
	boil.DebugMode = true
}

func main() {
	db, task, _ := mysql.Init(ctx, cfg)
	defer task.Shutdown(ctx)

	id := ulid.MustNew()
	office := &models.Office{ID: id, Name: id}
	office.InsertP(ctx, db, boil.Infer())
}
