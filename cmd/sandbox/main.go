package main

import (
	"context"
	"fmt"

	"github.com/ktakenaka/gosample2022/app/models"
	"github.com/ktakenaka/gosample2022/infra/database"
)

func main() {
	fmt.Println("hello")

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
	users, err := models.Users().All(ctx, db)
	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}
