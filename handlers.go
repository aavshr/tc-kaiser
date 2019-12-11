package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/labstack/echo"
)

// FetchS3 fetches S3 file based on domain
func FetchS3(c echo.Context) error {
	// printHeaders(c)
	forwardedHost, ok := c.Request().Header["X-Forwarded-Host"]
	if !ok {
		return c.JSON(http.StatusBadRequest, errorResp("No X-Forwarded-Host header set in request"))
	}
	key := getKeyFromSubDomain(forwardedHost[0])
	file, err := ac.DownloadFile(s3Bucket, key)
	if err != nil {
		var e awserr.Error
		if errors.As(err, &e) {
			if e.Code() == s3.ErrCodeNoSuchKey {
				fmt.Println("Bucket:", s3Bucket)
				fmt.Println("Key:", key)
				return c.JSON(http.StatusNotFound, errorResp("No file found for the host"))
			}
		}
		fmt.Println("handlers.FetchS3:", err)
		fmt.Println("Bucket:", s3Bucket)
		fmt.Println("Key:", key)
		return c.JSON(http.StatusInternalServerError, errorResp("Internal server error"))
	}
	return c.String(http.StatusOK, string(file))
}
