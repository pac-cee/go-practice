package handlers

import (
    "net/http"
    "wallet-app/internal/services"
    "wallet-app/pkg/utils"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type WalletHandler struct {
    service *services.WalletService
}

func NewWalletHandler(service *services.WalletService) *WalletHandler {
    return &WalletHandler{service: service}
}

