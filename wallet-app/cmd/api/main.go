package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wallet-app/internal/config"
	"wallet-app/internal/handlers"
	"wallet-app/internal/middleware"
	"wallet-app/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
