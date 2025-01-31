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
	m.errors = append(m.errors, message)
}

// TestShapes tests the Shape interface implementations
func TestShapes(t *testing.T) {
	t.Run("circle", func(t *testing.T) {
		c := Circle{radius: 5}
		expectedArea := math.Pi * 25
		expectedPerimeter := 2 * math.Pi * 5

		if area := c.Area(); math.Abs(area-expectedArea) > 0.001 {
			t.Errorf("Circle.Area() = %v, want %v", area, expectedArea)
		}

		if perim := c.Perimeter(); math.Abs(perim-expectedPerimeter) > 0.001 {
			t.Errorf("Circle.Perimeter() = %v, want %v", perim, expectedPerimeter)
		}

		if str := c.String(); !strings.Contains(str, "radius=5") {
			t.Errorf("Circle.String() = %v, want to contain 'radius=5'", str)
		}
	})

	t.Run("rectangle", func(t *testing.T) {
		r := Rectangle{width: 4, height: 6}
		expectedArea := 24.0
		expectedPerimeter := 20.0

		if area := r.Area(); area != expectedArea {
			t.Errorf("Rectangle.Area() = %v, want %v", area, expectedArea)
		}

		if perim := r.Perimeter(); perim != expectedPerimeter {
			t.Errorf("Rectangle.Perimeter() = %v, want %v", perim, expectedPerimeter)
		}

		if str := r.String(); !strings.Contains(str, "width=4") || 
			!strings.Contains(str, "height=6") {
