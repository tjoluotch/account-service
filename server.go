package main

import (
	"consul-service/internal/api"
	"consul-service/internal/config"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

var (
	GRPC_STUB = "localhost:3000"
	optsGrpc  []grpc.DialOption
)

func GrpcInit(address string, service *api.Service) *grpc.ClientConn {
	// create grpc without tls transport security
	optsGrpc = append(optsGrpc, grpc.WithInsecure())
	optsGrpc = append(optsGrpc, grpc.WithBlock())
	conn, err := grpc.Dial(address, optsGrpc...)
	if err != nil {
		service.Logger.Fatalw("failed to open grpc client connection",
			"error", err)
	}
	return conn
}

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
