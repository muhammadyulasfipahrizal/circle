package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message  string `json:"message"`
	Hostname string `json:"hostname"`
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte(`{"status":"ok"}` + "\n"))
	if err != nil {
		log.Printf("failed to write health response: %v", err)
	}
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("failed to get hostname: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message:  "Hello",
			Hostname: hostname,
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("failed to encode response: %v", err)
		}
	})

	http.HandleFunc("/health", health)

	log.Printf("Starting app on the port 8080 and hostname %s", hostname)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
