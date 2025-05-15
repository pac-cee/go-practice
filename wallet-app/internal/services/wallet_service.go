package services

import (
    "context"
    "wallet-app/internal/domain/models"
    "wallet-app/internal/domain/repository"

    "github.com/google/uuid"
)

type WalletService struct {
    repo *repository.WalletRepository
}

func NewWalletService(repo *repository.WalletRepository) *WalletService {
    return &WalletService{repo: repo}
}

func (s *WalletService) CreateWallet(ctx context.Context, userID uuid.UUID, currency string) (*models.Wallet, error) {
