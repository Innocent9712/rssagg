package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an API Key from the headers of an HTTP request
// Example:
// 	Authorization: ApiKey {insert apiKey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication header found")
	}

	apiKey := strings.Split(val, " ")
	// apiKey := strings.TrimPrefix(val, "ApiKey ")

	if len(apiKey) != 2 || apiKey[0] != "ApiKey" {
		return "", errors.New("malformed auth header")
	}

	return apiKey[1], nil
}
