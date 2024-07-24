package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestPing tests the /ping endpoint
func TestPing(t *testing.T) {
	// Create a new server instance
	srv := NewServer(":8081")
	srv.HandleEndpoints()

	// Create a new HTTP request to the /ping endpoint
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Serve the HTTP request
	srv.serverMux.ServeHTTP(rr, req)

	// Check the status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Define the expected response struct
	expected := PingResponse{
		Id:  0,
		Msg: "Pong",
	}

	// Decode the actual response body
	var actual PingResponse
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Compare the actual and expected responses
	if actual != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", actual, expected)
	}
}
