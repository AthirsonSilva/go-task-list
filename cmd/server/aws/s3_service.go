package awsservice

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"net/http"
	"os"
)

var bucketName = "music-api-golang"

func getClient() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	c := s3.NewFromConfig(cfg)
	return c
}

func PutBucketObject(key string, fileName string) (err error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Error opening file: %s", err)
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Printf("Error getting file info: %s", err)
		return err
	}
	size := fileInfo.Size()

	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	if err != nil {
		log.Printf("Error reading file: %s", err)
		return err
	}

	putObject := &s3.PutObjectInput{
		Bucket:             aws.String(bucketName),
		Key:                aws.String(key),
		Body:               bytes.NewReader(buffer),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ContentLength:      aws.Int64(size),
		ContentDisposition: aws.String("attachment"),
		StorageClass:       "INTELLIGENT_TIERING",
	}

	output, err := getClient().PutObject(context.TODO(), putObject)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("Successfully uploaded to S3", output)
	return nil
}
