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
	queueName = "sandbox"
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

	go func() {
		body := aws.String(time.Now().Format(time.RFC3339) + "hello")
		client.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(0),
			MessageBody:  body,
			QueueUrl:     queueURL.QueueUrl,
		})
	}()

	result, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:              queueURL.QueueUrl,
		AttributeNames:        aws.StringSlice([]string{"SentTimestamp"}),
		MaxNumberOfMessages:   aws.Int64(1),
		MessageAttributeNames: aws.StringSlice([]string{sqs.QueueAttributeNameAll}),
		WaitTimeSeconds:       aws.Int64(20),
	})
	if err != nil {
		panic(err)
	}

	for _, msg := range result.Messages {
		fmt.Println(*msg.MessageId, ":", *msg.Body)
		_, err := client.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      queueURL.QueueUrl,
			ReceiptHandle: msg.ReceiptHandle,
		})
		if err != nil {
			panic(err)
		}
	}
}
