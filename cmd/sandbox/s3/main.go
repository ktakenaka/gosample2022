package main

import (
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
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

	f, err := os.Open("go.mod")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	out2, err := client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("example"),
		Body:   f,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out2.ETag)

	out3, err := client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("example"),
	})
	if err != nil {
		log.Fatal(err)
	}
	filebytes, _ := io.ReadAll(out3.Body)
	log.Println(string(filebytes))
}
