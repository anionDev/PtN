package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandler_SendsAlertToNtfy(t *testing.T) {
	// Set fake ntfy URL (mock server)
	mockNtfy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Optional: Check content
		if r.Method != "POST" {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "text/plain" {
			t.Errorf("Expected Content-Type text/plain, got %s", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer mockNtfy.Close()

	// Set environment variable for ntfy URL
	os.Setenv("NTFY_URL", mockNtfy.URL)

	// Fake Alertmanager JSON payload
	alertJSON := `{
		"status": "firing",
		"alerts": [{
			"labels": {
				"alertname": "WebsiteDown",
				"severity": "critical"
			},
			"annotations": {
				"summary": "The website is not reachable"
			}
		}]
	}`

	req := httptest.NewRequest("POST", "/alert", bytes.NewBuffer([]byte(alertJSON)))
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}
