# Microservice-based E-commerce System

A scalable e-commerce system built with microservices architecture using Go, MongoDB, and Docker.

## Architecture Overview

The system consists of four main services:

1. **API Gateway (Port 8080)**
   - Routes requests to appropriate services
   - Handles authentication
   - Implements CORS and logging

2. **Product Service (Port 8081)**
   - Manages product catalog
   - Handles inventory
   - CRUD operations for products

3. **Order Service (Port 8082)**
   - Processes orders
   - Manages order status
   - Validates product availability

4. **User Service (Port 8083)**
   - User authentication
   - Profile management
   - JWT token generation

## Technologies Used

- Go (Golang) 1.19
- MongoDB
- Docker & Docker Compose
- JWT for authentication
- Gorilla Mux for routing
- WebSocket for real-time updates

## Prerequisites

1. Docker and Docker Compose installed
2. Go 1.19 or later (for development)
3. MongoDB (handled by Docker)

## Project Structure

```
06-ecommerce/
├── api-gateway/
│   ├── main.go
│   └── Dockerfile
├── services/
│   ├── product/
│   │   ├── main.go
│   │   ├── Dockerfile
│   │   └── go.mod
│   ├── order/
│   │   ├── main.go
