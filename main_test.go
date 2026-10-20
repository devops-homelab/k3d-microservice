package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string
		status int
		body   string
	}{
		{"greeting", http.MethodGet, "/", http.StatusOK, "{\"message\":\"Hello from k3d-microservice!!\"}\n"},
		{"health", http.MethodGet, "/health", http.StatusOK, "{\"status\":\"ok!!\"}\n"},
		{"not found", http.MethodGet, "/missing", http.StatusNotFound, ""},
		{"unsupported method", http.MethodPost, "/", http.StatusMethodNotAllowed, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			handler().ServeHTTP(response, httptest.NewRequest(tt.method, tt.path, nil))

			if response.Code != tt.status {
				t.Fatalf("status = %d, want %d", response.Code, tt.status)
			}
			if tt.body != "" {
				if got := response.Header().Get("Content-Type"); got != "application/json" {
					t.Errorf("Content-Type = %q, want application/json", got)
				}
				if got := response.Body.String(); got != tt.body {
					t.Errorf("body = %q, want %q", got, tt.body)
				}
			}
		})
	}
}
