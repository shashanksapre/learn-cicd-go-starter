package auth

import (
	"net/http"
	"testing"
)

func TestAuthorizationHeader(t *testing.T) {
	// no authorization header should return error
	headers := http.Header{}
	headers.Add("hello", "there")
	_, err := GetAPIKey(headers)

	if err == nil {
		t.Errorf("Expected an error for no authorization header")
	}

	// invalid authorization header should return error
	headers.Add("Authorization", "bearer token")
	_, err = GetAPIKey(headers)

	if err == nil {
		t.Errorf("Expected an error for invalid authorization header")
	}

	// correct authorization header should not return error
	headers.Set("Authorization", "ApiKey 123")
	apiKey, err := GetAPIKey(headers)

	if err != nil {
		t.Errorf("Expected to pass with correct authorization header")
	}

	if apiKey != "123" {
		t.Errorf("Expected ApiKey to be 123, but received %s", apiKey)
	}
}
