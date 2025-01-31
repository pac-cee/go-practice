package main

import (
	"math"
	"strings"
	"testing"
)

// MockLogger for testing
type MockLogger struct {
	logs   []string
	errors []string
}

func (m *MockLogger) Log(message string) {
	m.logs = append(m.logs, message)
}

func (m *MockLogger) Error(message string) {
