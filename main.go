package main

import (
	"log"
	"os"

	"github.com/labstack/echo"
)

var (
	s3Bucket  string
	awsRegion string
	ac        awsClient
)

func bootstrap() {
	if s3Bucket = os.Getenv("S3_BUCKET"); s3Bucket == "" {
		panic("No env var S3_BUCKET set")
	}
	if awsRegion = os.Getenv("AWS_REGION"); awsRegion == "" {
		panic("No env var AWS_REGION set")
	}
	ac.init(awsRegion)
}

func main() {
	bootstrap()
	api := echo.New()
	api.Any("/", FetchS3)
	log.Fatal(api.Start(":5000"))
	return
}
