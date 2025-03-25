# Project 4: Expense Tracker API (with PostgreSQL)

## Overview
This project demonstrates how to build a Go backend API that connects to a PostgreSQL database. You'll learn how to use Go's `database/sql` package to perform queries and return data as JSON.

## Learning Goals
- Use `database/sql` to interact with relational databases
- Connect to PostgreSQL from Go
- Map SQL query results to Go structs
- Serve data over HTTP as JSON

## Endpoints
- `GET /expenses` — List all expenses from the database

## How to Run
1. Make sure PostgreSQL is running and accessible.
2. Create a database named `expenses` and a table:
   ```sql
   CREATE TABLE expenses (
