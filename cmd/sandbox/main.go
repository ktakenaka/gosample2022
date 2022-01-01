package main

import (
	"context"
	"fmt"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"github.com/ktakenaka/gosample2022/cmd/config"
	"github.com/ktakenaka/gosample2022/infra/database"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	cfg, err := config.Initialize()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", cfg.DB.Read)
	fmt.Printf("%v\n", cfg.DB.Write)

	db, err := database.New(cfg.DB.Write)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	id, _ := ulid.GenerateID()
	user := models.User{
		ID:    id,
		Email: "example@hoge.com",
	}

	boil.DebugMode = true
	err = user.Upsert(ctx, db, boil.Infer(), boil.Infer())
	if err != nil {
		panic(err)
	}
	users, err := models.Users().All(ctx, db)
	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}
