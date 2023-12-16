package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/trayanus1026/gambling-platform/wallet"
	"github.com/trayanus1026/gambling-platform/ws"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcServerAddr = "localhost:50051"

func setupGRPCConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(grpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial gRPC server: %v", err)
	}
	return conn, nil
}

func main() {
	webSocket := ws.NewWebSocketServer()
	walletHandler := wallet.NewWalletHandler(webSocket)
	grpcService := wallet.NewWalletGRPCService(walletHandler)

	router := mux.NewRouter()
	router.HandleFunc("/api/wallet/deposit", walletHandler.Deposit).Methods("POST")
	router.HandleFunc("/api/wallet/withdraw", walletHandler.Withdraw).Methods("POST")
	router.HandleFunc("/api/wallet/balance/{user_id}", walletHandler.GetBalance).Methods("GET")
	router.HandleFunc("/ws", webSocket.HandleConnections)
	http.Handle("/", router)

	grpcServer := grpc.NewServer()
	wallet.RegisterWalletServiceServer(grpcServer, grpcService)
	// Set up gRPC connection for API
	grpcConn, err := setupGRPCConnection()
	if err != nil {
		fmt.Println("Error setting up gRPC connection:", err)
		return
	}
	defer grpcConn.Close()

	// Start WebSocket listener in a separate goroutine
	fmt.Printf("Start WebSocket listener\n")
	go func() {
		for {
			// Simulate events (replace with actual game outcomes or leaderboard changes)
			walletHandler.BroadcastGameOutcome("Game outcome: You won!")
			walletHandler.BroadcastLeaderboardChange("Leaderboard changed")

			<-time.After(10 * time.Second)
		}
	}()

	grpcListener, err := net.Listen("tcp", grpcServerAddr)
	if err != nil {
		log.Fatalf("Error setting up gRPC connection: %v", err)
	}

	// Start gRPC server
	go func() {
		fmt.Println("Starting gRPC server on", grpcServerAddr)
		if err := grpcServer.Serve(grpcListener); err != nil {
			fmt.Println("Error starting gRPC server:", err)
		}
	}()

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}
