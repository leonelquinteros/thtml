package main

import (
	"net/http/httptest"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	// Init config
	_publicPath = "_example/public"
	_templatesPath = "_example/templates"

	// Init handler
	h := thtmlHandler{}

	// Test requests
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/index.html", nil)
	h.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Fatalf("Expected response code 200. Got %d", resp.Code)
	}

	resp = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/index", nil)
	h.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Fatalf("Expected response code 200. Got %d", resp.Code)
	}

	resp = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/", nil)
	h.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Fatalf("Expected response code 200. Got %d", resp.Code)
	}

	resp = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/not-found", nil)
	h.ServeHTTP(resp, req)
	if resp.Code != 404 {
		t.Fatalf("Expected response code 404. Got %d", resp.Code)
	}
}
