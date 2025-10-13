package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	// Creates an instance of *http.Request to send to the route handler.
	// It simulates a client making a request to the server.
	mockRequest := httptest.NewRequest("GET", "/", nil)

	// NewRecorder returns an instance of the ResponseRecorder struct.
	// The ResponseRecorder struct is an implementation of ResponseWriter
	// that records its mutations for later inspection in tests like these.
	// Instead of writing the response over a network, it stores the status
	// codes, headers and body in its fields.
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(RootHandler)

	// Calls the handler function (RootHandler) with the test-specific implementations of the ResponseWriter and *Request structs.
	handler.ServeHTTP(responseRecorder, mockRequest)

	// Verify Same Status
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify Same Body
	expected := "you just reached the main route."
	if responseRecorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), expected)
	}
}
