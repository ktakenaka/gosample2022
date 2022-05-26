package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"

	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type Config struct {
	ID       string
	Secret   string
	Region   string
	Endpoint string
}

type Session interface {
	client.ConfigProvider
}

type SQS interface {
	sqsiface.SQSAPI
}

type KMS interface {
	kmsiface.KMSAPI
}

type S3 interface {
	s3iface.S3API
}

func NewSession(cfg *Config) (Session, error) {
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

func NewSQS(sess Session) SQS {
	return sqs.New(sess)
}

func NewKMS(sess Session) KMS {
	return kms.New(sess)
}

func NewS3(sess Session) S3 {
	return s3.New(sess)
}
