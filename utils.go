package main

import "strings"

func errorResp(message string) map[string]string {
	return map[string]string{
		"error": message,
	}
}

func getKeyFromSubDomain(subdomain string) string {
	return strings.Split(subdomain, ".")[0]
}
