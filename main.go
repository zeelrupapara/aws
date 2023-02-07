package main

import (
	"aws-pcts/s3"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	S3Config := s3.S3Config{
		Endpoint:        os.ExpandEnv(`${S3_ENDPOINT}`),
		AccessKey:       os.ExpandEnv(`${S3_ACCESS_KEY}`),
		SecretAccessKey: os.ExpandEnv(`${S3_SECRET_KEY}`),
		Region:          os.ExpandEnv(`${S3_REGION}`),
		Bucket:          os.ExpandEnv(`${S3_BUCKET}`),
	}

	s3Client, err := s3.New(S3Config)
	if err != nil {
		log.Fatal(err)
	}
	println("S3 Client created", &s3Client)
	err = s3Client.UploadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
}
