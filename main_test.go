// main_test.go
package main

import (
	"context"
	"testing"

	"github.com/trayanus1026/gambling-platform/wallet"
	"google.golang.org/grpc"
)

// TestGetBalance tests the GetBalance gRPC endpoint
func TestGetBalance(t *testing.T) {
	// Set up a gRPC connection
	conn, err := grpc.Dial(grpcServerAddr, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := wallet.NewWalletServiceClient(conn)

	// Simulate a user ID
	userID := "3"

	// Make a GetBalance request
	request := &wallet.BalanceRequest{UserId: userID}
	response, err := client.GetBalance(context.Background(), request)
	if err != nil {
		t.Fatalf("Error calling GetBalance: %v", err)
	} else {
		t.Fatalf("Balance: %v", response.Balance)
	}

	// Check the expected balance (modify this based on your test scenario)
	// expectedBalance := int32(100) // Assuming the balance should be 100
	// if response.Balance != expectedBalance {
	// 	t.Errorf("Expected balance %d, but got %d", expectedBalance, response.Balance)
	// }
}
