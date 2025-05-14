package models

import (
    "time"

    "github.com/google/uuid"
)

type TransactionType string

const (
    Deposit  TransactionType = "deposit"
    Withdraw TransactionType = "withdraw"
)

type Wallet struct {
    ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
    UserID    uuid.UUID `json:"user_id" gorm:"type:uuid"`
    Balance   float64   `json:"balance"`
