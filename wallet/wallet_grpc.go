package wallet

import (
	"context"
)

// WalletGRPCService is the gRPC service implementation
type WalletGRPCService struct {
	walletStore *WalletHandler
}

// mustEmbedUnimplementedWalletServiceServer implements WalletServiceServer.
func (*WalletGRPCService) mustEmbedUnimplementedWalletServiceServer() {
	panic("unimplemented")
}

// NewWalletGRPCService creates a new instance of WalletGRPCService
func NewWalletGRPCService(wh *WalletHandler) *WalletGRPCService {
	return &WalletGRPCService{
		walletStore: wh,
	}
}

// GetBalance implements the gRPC GetBalance RPC method
func (s *WalletGRPCService) GetBalance(ctx context.Context, req *BalanceRequest) (*BalanceResponse, error) {

	userID := req.GetUserId()
	balance := s.walletStore.GetBalanceByUserId(userID)

	return &BalanceResponse{Balance: int32(balance)}, nil
}
