# Real-time Chat Application

A web-based chat application demonstrating real-time communication using WebSocket protocol in Go.

## Concepts Covered

- WebSocket protocol
- Goroutines and channels
- Concurrent programming
- JSON encoding/decoding
- Web server creation
- Client-server communication
- Event handling
- Frontend development

## Features

- Real-time messaging
- User join/leave notifications
- Message timestamps
- Clean and responsive UI
- Username support
- System messages
- Auto-scroll to latest messages

## Prerequisites

1. Go installed on your system
2. Install required Go package:
   ```bash
   go get github.com/gorilla/websocket
   ```

## How to Run

1. Navigate to the project directory:
   ```bash
   cd level-2/05-chat-app
   ```

2. Run the program:
   ```bash
   go run main.go
   ```

3. Open your browser and visit:
   ```
   http://localhost:8080
   ```

4. Open multiple browser windows to simulate different users.

## Project Structure

```
05-chat-app/
├── main.go          # Main server file with WebSocket handling
