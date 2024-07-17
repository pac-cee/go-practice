// boilerplate.go
// Common Go boilerplate code snippets for quick reference.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Minimal runnable Go program
func minimal() {
	fmt.Println("Hello, World!")
}

// Basic HTTP server
func httpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, HTTP!")
	})
	log.Println("Starting server at :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Simple CLI structure
func cliExample() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run boilerplate.go <command>")
		return
	}
	switch os.Args[1] {
	case "hello":
		fmt.Println("Hello from CLI!")
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
