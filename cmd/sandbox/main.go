package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	pkgNotifier "github.com/ktakenaka/gosample2022/app/pkg/notifier"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/notifier"
	"github.com/ktakenaka/gosample2022/infra/database"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	cfg, err := config.Initialize()
	if err != nil {
		panic(err)
	}

	db, err := database.New(cfg.DB.Write)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	boil.DebugMode = true
	users, err := models.Users().All(ctx, db)
	if err != nil {
		panic(err)
	}

	user := users[0]

	ntfr, _ := notifier.Init(cfg)

	ntfr.ErrorWithExtrasAndContext(
		pkgNotifier.NewPersonContext(ctx, ulid.ULID(user.ID).String()),
		pkgNotifier.WARN,
		fmt.Errorf("hellow"),
		map[string]interface{}{"user": user},
	)
	time.Sleep(time.Second * 3)
}
