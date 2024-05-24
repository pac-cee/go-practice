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
        return
    }
    defer ln.Close()
    fmt.Println("TCP server listening on port 9000")
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Printf("Received: %s", message)
    conn.Write([]byte(strings.ToUpper(message)))
