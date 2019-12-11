package main

func errorResp(message string) map[string]string {
	return map[string]string{
		"error": message,
	}
}
