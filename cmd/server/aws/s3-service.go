package awsservice

import (
	"bytes"
	"context"
	"fmt"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"net/http"
	"os"
)

var bucketName = "todo-list-golang"

func getClient() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Error("getClient", "Error loading config: "+err.Error())
	}

	c := s3.NewFromConfig(cfg)
	return c
}

func PutBucketObject(key string, filePath string) (s3FilePath string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		logger.Error("PutBucketObject", "Error opening file: "+err.Error())
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		logger.Error("PutBucketObject", "Error getting file info: "+err.Error())
		return "", err
	}
	size := fileInfo.Size()

	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	if err != nil {
		logger.Error("PutBucketObject", "Error reading file: "+err.Error())
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
		logger.Error("PutBucketObject", "Error putting object: "+err.Error())
		return "", err
	}

	photoUrl := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, key)
	logger.Error("PutBucketObject", "Successfully uploaded to S3 => "+photoUrl)
	return photoUrl, nil
}
