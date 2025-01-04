# Go Messaging System

This module demonstrates how to implement a real-time messaging system in Go using WebSockets and a pub/sub pattern.

## Code Explanation

The `main.go` file implements a complete messaging system with the following components:

### 1. Data Models
```go
type Message struct {
    ID        string      `json:"id"`
    Topic     string      `json:"topic"`
    Data      interface{} `json:"data"`
    Timestamp time.Time   `json:"timestamp"`
}

type Subscription struct {
    Topic    string
    Channel  chan Message
    ClientID string
}
```
- Message structure
- Subscription management
- Topic organization
- Client tracking

### 2. Message Broker
```go
type MessageBroker struct {
    sync.RWMutex
    subscriptions map[string][]Subscription
    topics        map[string][]Message
}
```
- Thread-safe implementation
- Subscription management
- Message history
- Topic organization

### 3. WebSocket Client
```go
type WebSocketClient struct {
    ID       string
    Conn     *websocket.Conn
    Broker   *MessageBroker
    Topics   []string
    SendChan chan Message
}
```
- Client identification
- Connection management
- Topic subscriptions
- Message buffering

### 4. Pub/Sub Operations

#### Subscribe
```go
func (mb *MessageBroker) Subscribe(topic, clientID string) chan Message
```
- Topic subscription
- Channel creation
- Client tracking
- Message buffering

#### Publish
```go
func (mb *MessageBroker) Publish(msg Message)
```
- Message broadcasting
- History management
- Concurrent delivery
- Error handling

#### Unsubscribe
```go
func (mb *MessageBroker) Unsubscribe(topic, clientID string)
```
- Subscription cleanup
- Channel closing
- Resource management

### 5. WebSocket Handlers

#### Connection Handler
```go
func (mb *MessageBroker) HandleWebSocket(w http.ResponseWriter, r *http.Request)
```
- Connection upgrade
- Client initialization
- Error handling
- Origin checking

#### Message Pumps
```go
func (wc *WebSocketClient) ReadPump()
func (wc *WebSocketClient) WritePump()
```
- Bidirectional communication
- Message processing
- Connection maintenance
- Error handling

## Messaging Patterns

### 1. Pub/Sub
- Topic-based messaging
- Multiple subscribers
- Message history
- Real-time delivery

### 2. WebSocket
- Persistent connections
- Bidirectional communication
- Connection management
- Heartbeat mechanism

### 3. Message Queuing
- Message buffering
- Ordered delivery
- Channel management
- Backpressure handling

## API Endpoints

### HTTP Endpoints
- `POST /publish`: Publish a message
  ```json
  {
    "topic": "notifications",
    "data": {
      "message": "Hello World"
    }
  }
  ```

- `GET /topics/{topic}/history`: Get topic history

### WebSocket Endpoint
- `WS /ws`: WebSocket connection
  - Subscribe message:
    ```json
    {
      "action": "subscribe",
      "topic": "notifications"
    }
    ```
  - Unsubscribe message:
    ```json
    {
      "action": "unsubscribe",
      "topic": "notifications"
    }
    ```

## Best Practices Demonstrated
1. Concurrent message handling
2. Resource cleanup
3. Error management
4. Connection maintenance
5. Message buffering
6. Thread safety

## Running the Server
```bash
# Install dependencies
go get github.com/gorilla/mux
go get github.com/gorilla/websocket

