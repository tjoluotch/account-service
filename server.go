package main

import (
	"consul-service/internal/api"
	"consul-service/internal/config"
	"consul-service/internal/pb"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	GrpcServer      = "127.0.0.1:2000"
	optsGrpc        []grpc.DialOption
	rpcConnChannel  = make(chan *grpc.ClientConn, 1)
	rpcErrorChannel = make(chan error, 1)
	wg              sync.WaitGroup
)

func GrpcInit(address string, service *api.Service) (*grpc.ClientConn, error) {
	// create grpc without tls transport security
	optsGrpc = append(optsGrpc, grpc.WithInsecure())
	//optsGrpc = append(optsGrpc, grpc.WithBlock())
	conn, err := grpc.Dial(address, optsGrpc...)
	if err != nil {
		service.Logger.Errorw("failed to open grpc client connection",
			"error", err)
		return nil, err
	}
	return conn, nil
}

func main() {
	// logger init
	logger, err := config.BuildLogger()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	service := &api.Service{Logger: logger}
	logger.Info("logger initialised")

	// grpc client provisioning
	logger.Info("provisioning grpc client for connection to server:", GrpcServer)
	go func() {
		wg.Add(1)
		logger.Info("in seperate go routine, initializing grpc server")
		conn, err := GrpcInit(GrpcServer, service)
		if err != nil {
			logger.Error("in separate go routine, grpc init failed with error")
			rpcErrorChannel <- err
		}
		logger.Info("in separate go routine, grpc conn created")
		//	write connection type to chan
		rpcConnChannel <- conn
		wg.Done()
	}()
	wg.Wait()

	// check to see if error was sent error channel
	if 0 < len(rpcErrorChannel) {
		logger.Error("failed to succesfully create a grpc client connection to grpc server:", <-rpcErrorChannel)
		close(rpcErrorChannel)
		os.Exit(1)
	}
	close(rpcErrorChannel)

	conn := <-rpcConnChannel
	close(rpcConnChannel)
	defer conn.Close()
	client := pb.NewAccountRoutesClient(conn)
	logger.Info("added grpc client to service")
	service.Client = &client

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
