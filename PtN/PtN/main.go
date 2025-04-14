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

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Extract topic from URL path (e.g., /alert → alert)
		topic := strings.TrimPrefix(r.URL.Path, "/")
		if topic == "" {
			http.Error(w, "Missing topic in path", http.StatusBadRequest)
			return
		}

		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// Get environment variables
		ntfyServer := os.Getenv("NTFY_SERVER")
		ntfyUser := os.Getenv("NTFY_USER")
		ntfyPass := os.Getenv("NTFY_PASS")

		if ntfyServer == "" {
			http.Error(w, "Missing NTFY_SERVER env-var", http.StatusInternalServerError)
			return
		}

		// Send to ntfy
		url := fmt.Sprintf("%s/%s", ntfyServer, topic)
		req, err := http.NewRequest("POST", url, bytes.NewReader(body))
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}

		if ntfyUser!="" || ntfyPass!="" {
			req.SetBasicAuth(ntfyUser, ntfyPass)
		}
		req.Header.Set("Content-Type", "text/plain")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, "Failed to send request to ntfy", http.StatusBadGateway)
			log.Printf("Error sending to ntfy: %v", err)
			return
		}
		defer resp.Body.Close()

		log.Printf("Forwarded alert to topic %s – ntfy responded with %s", topic, resp.Status)
		w.WriteHeader(http.StatusOK)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
