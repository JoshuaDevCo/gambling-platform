// api/http/router.go
package http

import (
	"github.com/gorilla/mux"
)

func NewRouter(wh *WalletHandler) *mux.Router {
	router := mux.NewRouter()

	// Register API endpoints
	router.HandleFunc("/api/wallet/deposit", wh.Deposit).Methods("POST")

	// Add more endpoints as needed

	return router
}
