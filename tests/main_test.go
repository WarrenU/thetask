package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getName() (string, error) {
	// Simulate a successful fetch
	return "John Doe", nil
}

func getJoke() (string, error) {
	// Simulate a successful joke fetch
	return "Why did the chicken cross the road?", nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Simple handler that calls getName and getJoke
	name, _ := getName()
	joke, _ := getJoke()
	w.Write([]byte(name + ": " + joke))
}

func TestGetName(t *testing.T) {
	// Test getName success
	name, err := getName()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if name != "John Doe" {
		t.Fatalf("Expected name 'John Doe', got %v", name)
	}
}

func TestGetJoke(t *testing.T) {
	// Test getJoke success
	joke, err := getJoke()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if joke != "Why did the chicken cross the road?" {
		t.Fatalf("Expected joke 'Why did the chicken cross the road?', got %v", joke)
	}
}

func TestHandler(t *testing.T) {
	// Test handler to check combined response
	req, err := http.NewRequest("GET", "/api", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, status)
	}

	// Check response body
	expected := "John Doe: Why did the chicken cross the road?"
	if rr.Body.String() != expected {
		t.Errorf("Expected body %v, got %v", expected, rr.Body.String())
	}
}

func TestGetNameError(t *testing.T) {
	// Test getName error (simulate)
	getName := func() (string, error) {
		return "", errors.New("fetch name error")
	}

	_, err := getName()
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestGetJokeError(t *testing.T) {
	// Test getJoke error (simulate)
	getJoke := func() (string, error) {
		return "", errors.New("fetch joke error")
	}

	_, err := getJoke()
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
