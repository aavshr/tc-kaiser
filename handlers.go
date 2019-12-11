package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// FetchS3 fetches S3 file based on domain
func FetchS3(c echo.Context) error {
	forwardedHost, ok := c.Request().Header["X-Forwarded-Host"]
	if !ok {
		return c.JSON(http.StatusBadRequest, errorResp("No X-Forwarded-Host header set in request"))
	}
	fmt.Println(forwardedHost)
	resp := map[string]string{
		"message": "you're good babe",
	}
	return c.JSON(http.StatusOK, resp)
}
