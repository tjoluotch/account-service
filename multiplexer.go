package main

import (
	. "consul-service/internal/api"
	"github.com/gorilla/mux"
)

func ServerMux(service *Service) (*mux.Router, error) {
	router := mux.NewRouter()
	router.Use(service.LoggingMiddleware)
	router.HandleFunc("/api/health", service.HealthHandler)
	router.HandleFunc("/api/payment", service.PaymentHandler)
	return router, nil
}
