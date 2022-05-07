package kms

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/config"
	"github.com/ktakenaka/gosample2022/infra/kms"
)

func Init(ctx context.Context, cfg *config.Config) (*kms.Conn, error) {
	conn := kms.New(&kms.Config{
		Region:      "ap-northeast-1",
		MasterKeyID: "alias/local-kms-key",
		EndPoint:    "http://localstack:4566",
	})
	return conn, nil
}
