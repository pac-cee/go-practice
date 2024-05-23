package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {
    go startServer()
    startClient()
}

func startServer() {
    ln, err := net.Listen("tcp", ":9000")
    if err != nil {
        fmt.Println("Error starting server:", err)
