package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/trayanus1026/gambling-platform/wallet"
	"github.com/trayanus1026/gambling-platform/ws"
)

func main() {
	webSocket := ws.NewWebSocketServer()
	walletHandler := wallet.NewWalletHandler(webSocket)

	router := mux.NewRouter()
	router.HandleFunc("/api/wallet/deposit", walletHandler.Deposit).Methods("POST")
	router.HandleFunc("/api/wallet/withdraw", walletHandler.Withdraw).Methods("POST")
	router.HandleFunc("/api/wallet/balance/{user_id}", walletHandler.GetBalance).Methods("GET")
	router.HandleFunc("/ws", webSocket.HandleConnections)
	http.Handle("/", router)

	// // Initialize gRPC server
	// grpcService := &grpc.WalletService{}
	// go startGRPCServer(grpcService)
	// Start WebSocket listener in a separate goroutine
	go func() {
		for {
			// Simulate events (replace with actual game outcomes or leaderboard changes)
			walletHandler.BroadcastGameOutcome("Game outcome: You won!")
			walletHandler.BroadcastLeaderboardChange("Leaderboard changed")

			<-time.After(10 * time.Second)
		}
	}()

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}

// func startGRPCServer(service *grpc.WalletService) {
// 	listen, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	grpc.RegisterWalletServiceServer(grpcServer, service)

// 	log.Println("gRPC server listening on :50051")
// 	if err := grpcServer.Serve(listen); err != nil {
// 		log.Fatalf("Failed to serve: %v", err)
// 	}
// }
