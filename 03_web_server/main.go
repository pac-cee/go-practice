package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"./config"
	"./middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		jsonResponse(w, http.StatusNotFound, Response{
			Status:  "error",
			Message: "Page not found",
		})
		return
	}

	jsonResponse(w, http.StatusOK, Response{
		Status:  "success",
		Message: "Welcome to the Go Web Server API",
		Data: map[string]string{
			"version": "1.0.0",
			"status":  "running",
		},
	})
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Services page")
}
