package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Alert struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type AlertmanagerPayload struct {
	Alerts []Alert `json:"alerts"`
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Extract topic from URL path (e.g., /alert → alert)
		topic := strings.TrimPrefix(r.URL.Path, "/")
		if topic == "" {
			http.Error(w, "Missing topic in path", http.StatusBadRequest)
			return
		}

		if ntfyServer == "" {
			http.Error(w, "Missing NTFY_SERVER env-var", http.StatusInternalServerError)
			return
		}

		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		var payload AlertmanagerPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			log.Println("Error parsing JSON:", err)
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		// Get environment variables
		ntfyServer := os.Getenv("NTFY_SERVER")
		ntfyUser := os.Getenv("NTFY_USER")
		ntfyPass := os.Getenv("NTFY_PASS")

		// Compose message text
		var messages []string
		for _, alert := range payload.Alerts {
			title := alert.Labels["summary"]
			message := alert.Annotations["description"]
	
			if title == "" {
				title = "Alert"
			}
			if message == "" {
				message = "No description provided"
			}
	
			req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", ntfyBase, topic), bytes.NewBufferString(message))
			if err != nil {
				log.Println("Failed to create request:", err)
				continue
			}
			req.Header.Set("Title", title)
			req.SetBasicAuth(ntfyUser, ntfyPass)
	
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("Error sending to ntfy: %v\n", err)
				continue
			}

			bodyBytes, err := io.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil || (399 < resp.StatusCode) {
				http.Error(w, "Failed to send request to ntfy.", http.StatusBadGateway)
				log.Printf("Failed to send request to ntfy. Response-body: %s",  string(bodyBytes))
				if err != nil {
					log.Printf("Error sending to ntfy: %v", err)
				}
			}
		}
		log.Printf("Forwarded alert to topic %s – ntfy responded with %s", topic, resp.StatusCode)
		w.WriteHeader(http.StatusOK)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
