package main

import (
	"log"
	"net"
	"net/http"
	//	"github.com/trayanus1026/gambling-platform/grpc"
)

func main() {
	// Initialize your in-memory data store

	// Initialize HTTP server and router
	walletHandler := http.NewWalletHandler()
	router := http.NewRouter(walletHandler)

	// Initialize WebSocket server
	// wsServer := &websocket.WebSocketServer{}

	// // Initialize gRPC server
	// grpcService := &grpc.WalletService{}
	// go startGRPCServer(grpcService)

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func startGRPCServer(service *grpc.WalletService) {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpc.RegisterWalletServiceServer(grpcServer, service)

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
