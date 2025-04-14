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
