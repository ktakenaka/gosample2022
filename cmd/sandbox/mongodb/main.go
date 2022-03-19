package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/mongodb"
	"github.com/ktakenaka/gosample2022/cmd/internal/shutdown"
)

var (
	cfg, _ = config.Initialize()
	ctx    = context.Background()
)

func main() {
	tasks := shutdown.New()
	defer tasks.Shutdown(ctx)

	db, task, err := mongodb.Init(ctx, cfg)
	if err != nil {
		panic(err)
	}
	tasks.Add(task)

	collection := db.Collection("numbers")
	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.InsertedID)
}
