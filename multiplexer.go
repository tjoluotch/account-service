package main

import (
	"consul-service/internal/api"
	"github.com/gorilla/mux"
)

func ServerMux() (*mux.Router, error) {
	router := mux.NewRouter()
	router.HandleFunc("/api/health", api.HealthHandler)
	router.HandleFunc("/api/payment", api.PaymentHandler)
	router.Use(api.LoggingMiddleware)
	return router, nil
}
