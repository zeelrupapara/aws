package s3

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct {
	S3Client *s3.Client
	Cfg      S3Config
}

type S3Config struct {
	Endpoint        string
	AccessKey       string
	SecretAccessKey string
	Region          string
	Bucket          string
}

// configS3 creates the S3 client
func New(cfg S3Config) (*S3, error) {
	fmt.Println("aws===>", cfg)

	endpointresolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{PartitionID: "aws",
			URL:               cfg.Endpoint,
			SigningRegion:     cfg.Region,
			HostnameImmutable: true}, nil
	})
	creds := credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretAccessKey, "")

	loadedCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(cfg.Region), config.WithEndpointResolverWithOptions(endpointresolver))

	awsS3Client := s3.NewFromConfig(loadedCfg)
	if err != nil {
		return &S3{S3Client: awsS3Client, Cfg: cfg}, err
	}
	return &S3{S3Client: awsS3Client, Cfg: cfg}, nil
}

// Upload log file to Bucket
func (s3conn *S3) UploadFile(filepath string) error {
	uploader := manager.NewUploader(s3conn.S3Client)
	_, err := uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(s3conn.Cfg.Bucket),
		Key:    aws.String(filepath),
		Body:   strings.NewReader("Hello World!"),
	})

	if err != nil {
		return err
	}

	return nil
}

// Download log file from Bucket
func (s3conn *S3) DownloadS3File(filepath string) ([]byte, error) {
	buffer := manager.NewWriteAtBuffer([]byte{})
	downloader := manager.NewDownloader(s3conn.S3Client)
	numBytes, err := downloader.Download(context.Background(), buffer, &s3.GetObjectInput{
		Bucket: aws.String(s3conn.Cfg.Bucket),
		Key:    aws.String(filepath),
	})
	if err != nil {
		return nil, err
	}
	if numBytes == 0 {
		return nil, errors.New("zero bytes written to memory")
	}
	return buffer.Bytes(), nil
}
