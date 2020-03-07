package main

import (
	"go.uber.org/zap"
	"net/http"
)

var (
	logger *zap.SugaredLogger
)

func main() {
	logger = setLogger()
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
