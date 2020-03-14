package main

import (
	"consul-service/internal/api"
	"consul-service/internal/config"
	"log"
	"net/http"
)

var ()

func main() {
	// logger init
	logger, err := config.BuildLogger()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	service := &api.Service{Logger: logger}
	logger.Info("logger initialised")

	// mux init
	logger.Info("attempt setup server multiplexer")
	mux, err := ServerMux(service)
	if err != nil {
		logger.Fatalw("Failed to setup multiplexer",
			"mux", mux)
	}
	logger.Infow("setup multiplexer",
		"mux", mux)
	logger.Info("starting server")
	//server startup
	logger.Fatal(http.ListenAndServe(":8080", mux))
}
