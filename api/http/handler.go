package http

import (
	"encoding/json"
	"net/http"
	"sync"
)

type DepositRequest struct {
	UserID string `json:"user_id"`
	Amount int    `json:"amount"`
}

type WalletHandler struct {
	mu     sync.Mutex
	wallet map[string]int
}

func NewWalletHandler() *WalletHandler {
	return &WalletHandler{
		wallet: make(map[string]int),
	}
}

func (wh *WalletHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	// Decode the JSON payload
	var request DepositRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the input (optional)

	// Update the user's wallet balance
	wh.wallet[request.UserID] += request.Amount

	// Respond with the updated balance
	response := map[string]int{"balance": wh.wallet[request.UserID]}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (wh *WalletHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	// Implement withdraw endpoint logic
}

func (wh *WalletHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	// Implement get balance endpoint logic
}
