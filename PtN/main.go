package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Alert is the expected structure of an Alertmanager webhook payload.
type Alert struct {
	Status string `json:"status"`
	Alerts []struct {
		Labels      map[string]string `json:"labels"`
		Annotations map[string]string `json:"annotations"`
	} `json:"alerts"`
}

// handler receives the webhook from Alertmanager and forwards each alert to ntfy
func handler(w http.ResponseWriter, r *http.Request) {
	var alert Alert

	// Read the incoming request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse JSON body into Alert struct
	if err := json.Unmarshal(body, &alert); err != nil {
		log.Println("Error parsing JSON:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Get the ntfy URL from environment
	ntfyURL := os.Getenv("NTFY_URL")
	if ntfyURL == "" {
		log.Println("NTFY_URL not set")
		http.Error(w, "Server misconfigured", http.StatusInternalServerError)
		return
	}

	// Iterate over all alerts and send them to ntfy
	for _, a := range alert.Alerts {
		message := fmt.Sprintf("🔔 [%s] %s\n%s",
			a.Labels["severity"],
			a.Labels["alertname"],
			a.Annotations["summary"])

		// Create POST request to ntfy
		req, err := http.NewRequest("POST", ntfyURL, bytes.NewBuffer([]byte(message)))
		if err != nil {
			log.Println("Error creating ntfy request:", err)
			continue
		}

		req.Header.Set("Content-Type", "text/plain")

		// If credentials are provided, set Basic Auth header
		if user := os.Getenv("NTFY_USER"); user != "" {
			pass := os.Getenv("NTFY_PASS")
			req.SetBasicAuth(user, pass)
		}

		// Send the message
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("Error sending to ntfy:", err)
			continue
		}
		resp.Body.Close()
		log.Printf("Message sent: %s (%d)", message, resp.StatusCode)
	}

	w.WriteHeader(http.StatusOK)
}

// main starts the HTTP server to receive webhooks from Alertmanager
func main() {
	http.HandleFunc("/alert", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Proxy running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
