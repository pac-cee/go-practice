package repository

import (
    "context"
    "wallet-app/internal/domain/models"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type WalletRepository struct {
    db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
    return &WalletRepository{db: db}
}

func (r *WalletRepository) Create(ctx context.Context, wallet *models.Wallet) error {
