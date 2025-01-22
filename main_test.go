package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetName will fail if the names service is failing:
func TestGetName(t *testing.T) {
	resp, err := http.Get("https://names.mcquay.me/api/v0/")
	if err != nil {
		t.Errorf("failed to fetch name: %s", err)
	}
	// Check the response status code
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, status)
	}
}

// TestGetName will fail if the joke service is failing:
func TestGetJoke(t *testing.T) {
	url := fmt.Sprintf("http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=%s&lastName=%s", "John", "Doe")
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("failed to fetch name: %s", err)
	}
	// Check the response status code
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, status)
	}
}

// TestRootHandler calls the actual external APIs and verifies the status code and basic response.
func TestRootHandler(t *testing.T) {
	// Create a new request to hit the root handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response using httptest.NewRecorder
	rr := httptest.NewRecorder()

	// Call the rootHandler
	handler := http.HandlerFunc(rootHandler)
	handler.ServeHTTP(rr, req)

	// Check if the status code is OK (200)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, status)
	}

	if rr.Body.String() == "" {
		t.Error("Unexcpected: Response body is empty")
	}
	fmt.Println(rr.Body.String())
}
