package main

import "net/http"

func ServerMux() (*http.ServeMux, error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", HealthHandler)
	return mux, nil
}
