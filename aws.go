package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type awsClient struct {
	sess *session.Session
	s3dl *s3manager.Downloader
}

func (c *awsClient) init(region string) {
	c.sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	c.s3dl = s3manager.NewDownloader(c.sess)
}

func (c *awsClient) DownloadFile(bucket, key string) ([]byte, error) {
	var buffer []byte
	awsBuffer := aws.NewWriteAtBuffer(buffer)
	_, err := c.s3dl.Download(awsBuffer, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Println("aws.DownloadFile:", err)
		return nil, err
	}
	return awsBuffer.Bytes(), nil
}
