package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	os.Setenv("NTFY_URL", "http://example.com")
	os.Unsetenv("NTFY_USER")
	os.Unsetenv("NTFY_PASSWORD")

	alertPayload := AlertManagerPayload{
		Alerts: []Alert{
			{
				Annotations: struct {
					Summary     string `json:"summary"`
					Description string `json:"description"`
				}{
					Summary:     "Test Title",
					Description: "Test description",
				},
			},
		},
	}
	data, _ := json.Marshal(alertPayload)
	req := httptest.NewRequest(http.MethodPost, "/testtopic", bytes.NewReader(data))
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadGateway && resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK or 502 BadGateway, got %d", resp.StatusCode)
	}
}
