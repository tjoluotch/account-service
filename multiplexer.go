package main

import (
	. "consul-service/internal/api"
	"github.com/gorilla/mux"
	"net/http"
)

func ServerMux(service *Service) (*mux.Router, error) {
	router := mux.NewRouter()
	router.Use(service.LoggingMiddleware)
	router.HandleFunc("/api/health", service.HealthHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/payment", service.PaymentHandler).Methods(http.MethodPut)
	return router, nil
}
