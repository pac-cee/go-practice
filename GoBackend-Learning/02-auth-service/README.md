# Project 2: Auth Service (JWT-based)

## Overview
This project demonstrates how to build a simple authentication service in Go using JWT (JSON Web Tokens) for stateless authentication and bcrypt for secure password storage.

## Learning Goals
- Build and use HTTP middleware in Go
- Securely hash and check passwords with bcrypt
- Generate and validate JWT tokens for authentication
- Understand stateless session management

## Endpoints
- `POST /signup` — Register a new user. Body: `{ "username": "user", "password": "pass" }`
- `POST /login` — Authenticate and receive a JWT. Body: `{ "username": "user", "password": "pass" }`
- `GET /protected` — Example protected route. Requires `Authorization` header with JWT token.

## How to Run
```sh
# In the project directory
