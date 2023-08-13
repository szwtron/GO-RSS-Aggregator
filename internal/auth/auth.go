package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey returns the API key from the request headers
// Authorization: ApiKey {your ApiKey here}
func GetAPIKey(headers http.Header) (string, error) {
	authorizationHeader := headers.Get("Authorization")
	if authorizationHeader == "" {
		return "", errors.New("Missing Authorization header")
	}

	authorizationHeaderParts := strings.Split(authorizationHeader, " ")
	if len(authorizationHeaderParts) != 2 {
		return "", errors.New("Invalid Authorization header")
	}

	if authorizationHeaderParts[0] != "ApiKey" {
		return "", errors.New("Invalid Authorization header")
	}
		
	return authorizationHeaderParts[1], nil
}