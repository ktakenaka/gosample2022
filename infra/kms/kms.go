package kms

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

type Config struct {
	Region string
}

func New(cfg *aws.Config) {
	svc := kms.New(session.Must(session.NewSession(cfg)))
	dataKey, err := svc.GenerateDataKeyPairWithoutPlaintext()
	if err != nil {
		return 
	}
}
