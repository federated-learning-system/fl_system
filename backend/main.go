package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HealthResponse defines the structure for the /healthz endpoint
type HealthResponse struct {
	Status string `json:"status"`
}

func main() {
	// 1. Health Check Endpoint
	// Returns 200 OK and {"status": "ok"}
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HealthResponse{Status: "ok"})
	})

	// 2. Metrics Endpoint
	// Placeholder for Prometheus text/plain metrics
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "# HELP http_requests_total Total number of HTTP requests.\n")
		fmt.Fprint(w, "# TYPE http_requests_total counter\n")
		fmt.Fprint(w, "http_requests_total{method=\"post\",endpoint=\"/upload\"} 0\n")
	})

	// Start the server
	port := ":8080"
	fmt.Printf("Server starting on %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
