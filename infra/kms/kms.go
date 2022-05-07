package kms

import (
	"context"
	"encoding/base64"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

type Config struct {
	Region      string
	MasterKeyID string
	EndPoint    string
}

func New(cfg *Config) *Conn {
	fResolver := func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{URL: cfg.EndPoint}, nil
	}
	fCred := func(context.Context) (aws.Credentials, error) {
		return aws.Credentials{AccessKeyID: "gosample2022", SecretAccessKey: "gosample2022"}, nil
	}

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithDefaultRegion(cfg.Region),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(fResolver)),
		config.WithCredentialsProvider(aws.CredentialsProviderFunc(fCred)),
		//config.WithClientLogMode(aws.LogSigning),
	)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := kms.NewFromConfig(awsCfg)
	return &Conn{client: client, masterKeyID: cfg.MasterKeyID}
}

type Conn struct {
	client      *kms.Client
	masterKeyID string
}

func (c *Conn) GenerateDataKeyPairWithoutPlaintext() (string, error) {
	input := &kms.EncryptInput{
		KeyId:     &c.masterKeyID,
		Plaintext: []byte("hoge"),
	}
	out, err := c.client.Encrypt(context.TODO(), input)
	if err != nil {
		return "", err
	}

	blobString := base64.StdEncoding.EncodeToString(out.CiphertextBlob)

	return blobString, nil
}
