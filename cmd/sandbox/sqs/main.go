package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			Config: aws.Config{
				Credentials:      credentials.NewStaticCredentials("gosample2022", "gosample2022", ""),
				Region:           aws.String("ap-northeast-1"),
				Endpoint:         aws.String("http://localstack:4566"),
				S3ForcePathStyle: aws.Bool(true),
			},
		},
	))
	var client sqsiface.SQSAPI = sqs.New(sess)
	out, err := client.ListQueues(&sqs.ListQueuesInput{})
	if err != nil {
		panic(err)
	}

	for _, url := range out.QueueUrls {
		fmt.Println(*url)
	}

	result, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: aws.String("http://localhost:4566/000000000000/sandbox"),
		AttributeNames: aws.StringSlice([]string{
			"SentTimestamp",
		}),
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: aws.StringSlice([]string{
			"All",
		}),
		WaitTimeSeconds: aws.Int64(20),
	})

	for _, msg := range result.Messages {
		fmt.Println(*msg.MessageId, ":", *msg.Body)
	}
}
