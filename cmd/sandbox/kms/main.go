package main

import (
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	infraAWS "github.com/ktakenaka/gosample2022/infra/aws"
)

var (
	cfg, _ = config.Initialize()
)

func main() {
	sess, _ := infraAWS.NewSession(&infraAWS.Config{
		ID: cfg.AWS.ID, Secret: cfg.AWS.Secret, Region: cfg.AWS.Region, Endpoint: cfg.AWS.Endpoint,
	})
	client := infraAWS.NewKMS(sess)

	out, err := client.Encrypt(&kms.EncryptInput{
		KeyId:     aws.String("alias/local-kms-key"),
		Plaintext: []byte("hoge"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(out.CiphertextBlob))
}
