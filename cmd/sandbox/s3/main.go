package main

import (
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	infraAWS "github.com/ktakenaka/gosample2022/infra/aws"
)

const (
	bucket = "s3sample"
)

var (
	cfg, _ = config.Initialize()
)

func main() {
	sess, _ := infraAWS.NewSession(&infraAWS.Config{
		ID: cfg.AWS.ID, Secret: cfg.AWS.Secret, Region: cfg.AWS.Region, Endpoint: cfg.AWS.Endpoint,
	})
	client := infraAWS.NewS3(sess)

	out, err := client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range out.Buckets {
		log.Println(*b.Name)
	}
}
