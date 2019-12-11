package main

import (
	"fmt"
	"strings"

	"github.com/labstack/echo"
)

func errorResp(message string) map[string]string {
	return map[string]string{
		"error": message,
	}
}

func getKeyFromSubDomain(subdomain string) string {
	folder := strings.Split(subdomain, ".")[0]
	return fmt.Sprintf("%s/index.html", folder)
}

func printHeaders(c echo.Context) {
	for k, v := range c.Request().Header {
		fmt.Println(k, v)
	}
}
