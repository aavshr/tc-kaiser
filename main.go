package main

import (
	"log"
	"os"

	"github.com/labstack/echo"
)

var (
	s3Bucket string
)

func bootstrap() {
	if s3Bucket = os.Getenv("S3_BUCKET"); s3Bucket == "" {
		panic("No env var S3_BUCKET set")
	}
}

func main() {
	api := echo.New()
	api.Any("/", FetchS3)
	log.Fatal(api.Start(":5000"))
	return
}
