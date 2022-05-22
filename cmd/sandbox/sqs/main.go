package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	infraAWS "github.com/ktakenaka/gosample2022/infra/aws"
)

const (
	queueName = "sqssample"
)

var (
	cfg, _ = config.Initialize()
)

func main() {
	sess, _ := infraAWS.NewSession(&infraAWS.Config{
		ID: cfg.AWS.ID, Secret: cfg.AWS.Secret, Region: cfg.AWS.Region, Endpoint: cfg.AWS.Endpoint,
	})
	client := infraAWS.NewSQS(sess)

	queueURL, err := client.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		panic(err)
	}

	body := aws.String(time.Now().Format(time.RFC3339) + "hello")
	out, err := client.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(0),
		MessageBody:  body,
		QueueUrl:     queueURL.QueueUrl,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(out.GoString())
}
