package main

import (
	"context"
	"fmt"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"github.com/ktakenaka/gosample2022/infra/database"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	cfg := &database.Config{
		User:     "writer",
		Password: "writer_password",
		Host:     "db",
		DBName:   "gosample2022_development",
	}
	db, err := database.New(cfg)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	id, _ := ulid.GenerateID()
	println(id.String())
	fmt.Println(id)
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
