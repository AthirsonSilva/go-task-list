package awsservice

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"net/http"
	"os"
)

var bucketName = "todo-list-golang"

func getClient() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	c := s3.NewFromConfig(cfg)
	return c
}

func PutBucketObject(key string, filePath string) (s3FilePath string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %s", err)
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Printf("Error getting file info: %s", err)
		return "", err
	}
	size := fileInfo.Size()

	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	if err != nil {
		log.Printf("Error reading file: %s", err)
		return "", err
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

	_, err = getClient().PutObject(context.TODO(), putObject)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	photoUrl := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, key)
	log.Printf("Successfully uploaded to S3 => %s", photoUrl)
	return photoUrl, nil
}
