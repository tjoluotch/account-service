package main

import (
	"consul-service/internal/config"
	"net/http"
)

var (
	logger = config.Logger{}
)

func main() {
	logger = config.Logger{}
	logger.Set()
	logger.Infow("attempt setup server multiplexer")
	mux, err := ServerMux()
	if err != nil {
		logger.Fatalw("Failed to setup multiplexer",
			"mux", mux)
	}
	logger.Infow("setup multiplexer",
		"mux", mux)
	logger.Info("starting server")
	logger.Fatal(http.ListenAndServe(":8080", mux))
}
