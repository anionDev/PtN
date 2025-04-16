// main.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Alert struct {
	Annotations struct {
		Summary     string `json:"summary"`
		Description string `json:"description"`
	} `json:"annotations"`
}

type AlertManagerPayload struct {
	Alerts []Alert `json:"alerts"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	topic := strings.TrimPrefix(r.URL.Path, "/")
	if topic == "" {
		http.Error(w, "missing topic in path", http.StatusBadRequest)
		return
	}

	ntfyURL := os.Getenv("NTFY_URL")
	if ntfyURL == "" {
		http.Error(w, "NTFY_URL environment variable not set", http.StatusInternalServerError)
		return
	}

	user := os.Getenv("NTFY_USER")
	password := os.Getenv("NTFY_PASSWORD")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return
	}

	var payload AlertManagerPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "invalid JSON payload", http.StatusBadRequest)
		return
	}

	for _, alert := range payload.Alerts {
		title := alert.Annotations.Summary
		message := alert.Annotations.Description
		
		if title == "" {
			title = "No title"
		}
		if message == "" {
			message = "No description"
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", ntfyURL, topic), bytes.NewBufferString(message))
		if err != nil {
			log.Printf("error creating request: %v", err)
			http.Error(w, "failed to create request", http.StatusInternalServerError)
			return
		}

		req.Header.Set("Title", title)
		if user != "" && password != "" {
			req.SetBasicAuth(user, password)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 300 {
			log.Printf("ntfy request failed: %v, status: %v", err, resp.Status)
			http.Error(w, "ntfy forwarding failed", http.StatusBadGateway)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("alerts forwarded to ntfy"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
