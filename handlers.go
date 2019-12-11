package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// FetchS3 fetches S3 file based on domain
func FetchS3(c echo.Context) error {
	for k, v := range c.Request().Header {
		fmt.Println(k, v)
	}
	forwardedHost, ok := c.Request().Header["X-Forwarded-Host"]
	if !ok {
		return c.JSON(http.StatusBadRequest, errorResp("No X-Forwarded-Host header set in request"))
	}
	resp := map[string]string{
		"s3Key": getKeyFromSubDomain(forwardedHost[0]),
	}
	return c.JSON(http.StatusOK, resp)
}
