package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type Config struct {
	ID       string
	Secret   string
	Region   string
	Endpoint string
}

type SQS interface {
	sqsiface.SQSAPI
}

func NewSession(cfg *Config) (*session.Session, error) {
	return session.NewSessionWithOptions(
		session.Options{
			Config: aws.Config{
				Credentials: credentials.NewStaticCredentials(cfg.ID, cfg.Secret, ""),
				Region:      aws.String(cfg.Region),
				Endpoint:    aws.String(cfg.Endpoint),
			},
		},
	)
}

func NewSQS(sess *session.Session) SQS {
	return sqs.New(sess)
}
