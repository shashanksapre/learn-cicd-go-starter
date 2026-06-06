package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	t.Run("no authorization header should return error", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("hello", "there")
		_, err := GetAPIKey(headers)

		if err == nil {
			t.Errorf("Expected an error for no authorization header")
		}
	})

	t.Run("invalid authorization header should return error", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("Authorization", "bearer token")
		_, err := GetAPIKey(headers)

		if err == nil {
			t.Errorf("Expected an error for invalid authorization header")
		}
	})

	t.Run("correct authorization header should not return error", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("Authorization", "ApiKey 123")
		apiKey, err := GetAPIKey(headers)

		if err != nil {
			t.Errorf("Expected to pass with correct authorization header")
		}

		if apiKey != "123" {
			t.Errorf("Expected ApiKey to be 123, but received %s", apiKey)
		}
	})
}
