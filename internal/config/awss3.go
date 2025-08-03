package config

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
}

func NewS3Config(region, accessKeyID, secretAccessKey, bucketName string) *S3Config {
	return &S3Config{
		Region:          region,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		BucketName:      bucketName,
	}
}

func (s *S3Config) NewS3Client() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(s.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			s.AccessKeyID,
			s.SecretAccessKey,
			"",
		)),
	)
	if err != nil {
		log.Printf("unable to load SDK config, %v", err)
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}

func (s *S3Config) UploadFile(client *s3.Client, key string, body []byte) error {
	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &s.BucketName,
		Key:    &key,
		Body:   bytes.NewReader(body),
	})
	return err
}

func (s *S3Config) GetObjectURL(key string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s",
		s.BucketName, s.Region, key)
}
