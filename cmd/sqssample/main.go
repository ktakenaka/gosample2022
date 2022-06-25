package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Result struct {
	ID          string `json:"id"`
	EventSource string `json:"event_source"`
	Body        string `json:"body"`
}

func handler() func(ctx context.Context, sqsEvent events.SQSEvent) ([]Result, error) {
	cfg, err := config.Initialize()
	if err != nil {
		panic(err)
	}
	db, _, _ := mysql.Init(context.TODO(), cfg)

	return func(ctx context.Context, sqsEvent events.SQSEvent) ([]Result, error) {
		results := []Result{}
		for _, msg := range sqsEvent.Records {
			results = append(results, Result{ID: msg.MessageId, EventSource: msg.EventSource, Body: msg.Body})
		}
		// No need to delete message by ourselves in the case of SQS + Lambda

		fmt.Printf("%+v\n", sqsEvent)

		// Check if the function can connect to DB
		id := ulid.MustNew()
		user := models.User{ID: id, Email: id + "@hoge.com"}
		return results, user.Insert(ctx, db, boil.Infer())
	}
}

func main() {
	dir, _ := os.Getwd()
	fmt.Println("===== current dir", dir)

	files, _ := ioutil.ReadDir("./")
	for i, f := range files {
		fmt.Println("file", i, ":", f.Name())
	}

	lambda.Start(handler())
}
