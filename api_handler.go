package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/trayanus1026/gambling-platform/wallet"
)

// APIHandler handles API requests
type APIHandler struct {
	walletStore  *wallet.WalletHandler
	walletClient wallet.WalletServiceClient
}

// GetBalanceHandler handles balance requests using gRPC
func (handler *APIHandler) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	// Retrieve balance using gRPC
	req := &wallet.BalanceRequest{UserId: userID}
	resp, err := handler.walletClient.GetBalance(context.Background(), req)
	if err != nil {
		http.Error(w, "Failed to get balance", http.StatusInternalServerError)
		return
	}

	// Respond with the balance
	response := map[string]interface{}{"user_id": userID, "balance": resp.Balance}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
