package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	homeHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %v", rec.Code)
	}

	expected := "Welcome to the Go API Server!\n"
	if rec.Body.String() != expected {
		t.Errorf("Expected body %q; got %q", expected, rec.Body.String())
	}
}

func TestHomeHandler_NotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/notfound", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	homeHandler(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status NotFound; go %v", rec.Code)
	}
}
