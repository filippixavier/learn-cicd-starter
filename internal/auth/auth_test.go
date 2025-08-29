package auth

import (
	"net/http"
	"testing"
)

func TestMissingHeader(t *testing.T) {
	header := make(http.Header)
	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatal("expected error, got nothing")
		return
	}
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("Expected %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}

func TestInvalidHeader(t *testing.T) {
	header := make(http.Header)
	header.Add("Authorization", "Toto")

	_, err := GetAPIKey(header)

	if err == nil {
		t.Fatal("Expected error, got nothing")
	}
}

func TestValidHeader(t *testing.T) {
	apiKey := "12345"
	header := make(http.Header)

	header.Add("Authorization", "ApiKey "+apiKey)

	key, err := GetAPIKey(header)

	if err != nil {
		t.Fatalf("Expected %v, got %v", apiKey, err)
		return
	}

	if key != apiKey {
		t.Fatalf("Expected %v, got %v", apiKey, err)
		return
	}
}
