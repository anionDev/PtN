package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestAlertForwarding(t *testing.T) {
	// Set environment variables for test
	os.Setenv("NTFY_SERVER", "http://example.com")
	os.Setenv("NTFY_USER", "testuser")
	os.Setenv("NTFY_PASS", "testpass")

	// Start a fake ntfy server to capture requests
	ntfyCalled := false
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/testtopic" {
			t.Errorf("Expected topic path '/testtopic', got '%s'", r.URL.Path)
		}
		body, _ := io.ReadAll(r.Body)
		if string(body) != "test alert" {
			t.Errorf("Expected body 'test alert', got '%s'", string(body))
		}
		ntfyCalled = true
		w.WriteHeader(http.StatusOK)
	}))
	defer testServer.Close()

	// Replace NTFY_SERVER with test server URL
	os.Setenv("NTFY_SERVER", testServer.URL)

	// Create request to our Go handler
	req := httptest.NewRequest("POST", "/testtopic", strings.NewReader("test alert"))
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		topic := strings.TrimPrefix(r.URL.Path, "/")
		if topic == "" {
			http.Error(w, "Missing topic in path", http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		ntfyServer := os.Getenv("NTFY_SERVER")
		ntfyUser := os.Getenv("NTFY_USER")
		ntfyPass := os.Getenv("NTFY_PASS")

		if ntfyServer == "" {
			http.Error(w, "Missing NTFY_SERVER env-var", http.StatusInternalServerError)
			return
		}

		url := ntfyServer + "/" + topic
		reqToNtfy, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}

		if ntfyUser!="" || ntfyPass!="" {
			req.SetBasicAuth(ntfyUser, ntfyPass)
		}
		reqToNtfy.Header.Set("Content-Type", "text/plain")

		resp, err := http.DefaultClient.Do(reqToNtfy)
		if err != nil {
			http.Error(w, "Failed to send request to ntfy", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(http.StatusOK)
	})

	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if !ntfyCalled {
		t.Error("Expected ntfy server to be called, but it wasn't")
	}
}
