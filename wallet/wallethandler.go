package wallet

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/trayanus1026/gambling-platform/ws"
)

type DepositRequest struct {
	UserID string `json:"user_id"`
	Amount int    `json:"amount"`
}

type WithdrawRequest struct {
	UserID string `json:"user_id"`
	Amount int    `json:"amount"`
}

type WithdrawResponse struct {
	Balance int `json:"balance"`
}

type WalletHandler struct {
	mu        sync.Mutex
	wallet    map[string]int
	webSocket *ws.WebSocketServer
}

func NewWalletHandler(webSocket *ws.WebSocketServer) *WalletHandler {
	return &WalletHandler{
		wallet:    make(map[string]int),
		webSocket: webSocket,
	}
}

func (wh *WalletHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	var request DepositRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	wh.wallet[request.UserID] += request.Amount

	response := map[string]int{"balance": wh.wallet[request.UserID]}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (wh *WalletHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	// Decode the JSON payload
	var request WithdrawRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate user ID and withdrawal amount
	if request.UserID == "" || request.Amount <= 0 {
		http.Error(w, "Invalid user ID or withdrawal amount", http.StatusBadRequest)
		return
	}

	// Check if the user has sufficient funds
	if wh.wallet[request.UserID] < request.Amount {
		http.Error(w, "Insufficient funds", http.StatusBadRequest)
		return
	}

	// Update the user's wallet balance
	wh.wallet[request.UserID] -= request.Amount

	// Respond with the updated balance
	response := WithdrawResponse{Balance: wh.wallet[request.UserID]}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (wh *WalletHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	// Extract user ID from the request parameters
	userID := mux.Vars(r)["user_id"]

	// Check if the user exists
	if balance, ok := wh.wallet[userID]; ok {
		// Respond with the user's wallet balance
		response := map[string]interface{}{"user_id": userID, "balance": balance}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
}

func (wh *WalletHandler) GetBalanceByUserId(userId string) int {
	wh.mu.Lock()
	defer wh.mu.Unlock()
	//	return wh.wallets[userID].Balance

	if balance, ok := wh.wallet[userId]; ok {
		return balance
	} else {
		return -1
	}
}

func (wh *WalletHandler) BroadcastGameOutcome(outcome string) {
	message := []byte("Game Outcome: " + outcome)
	wh.webSocket.Broadcast(websocket.TextMessage, message)
}

// BroadcastLeaderboardChange broadcasts a leaderboard change to all connected WebSocket clients.
func (wh *WalletHandler) BroadcastLeaderboardChange(leaderboard string) {
	message := []byte("Leaderboard Change: " + leaderboard)
	wh.webSocket.Broadcast(websocket.TextMessage, message)
}
