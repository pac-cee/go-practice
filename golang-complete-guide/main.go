package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"
	"net/http"
	"github.com/redis/go-redis/v9"
)

// BASIC CONCEPTS
// 1. Variables and Data Types
func basicTypes() {
	// Integer types
	var i int = 42
	var i8 int8 = 127
	var ui uint = 123

	// Float types
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793

	// Boolean
	var isTrue bool = true

	// String
	var str string = "Hello, Go!"

	// Type inference
	auto := "Type inferred automatically"

	fmt.Printf("Types: %T %T %T %T %T %T %T %T\n", i, i8, ui, f32, f64, isTrue, str, auto)
}

// 2. Constants
const (
	Pi       = 3.14159
	Username = "admin"
	// iota example
	Sunday = iota
	Monday
	Tuesday
)

// 3. Arrays and Slices
func arraysAndSlices() {
	// Array
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	// Slice
	slice := []int{1, 2, 3}
	slice = append(slice, 4)

	// Slice operations
	subSlice := slice[1:3]
	
	fmt.Println(arr, slice, subSlice)
}

// 4. Maps
func maps() {
	// Declaration and initialization
	scores := map[string]int{
		"Alice": 95,
		"Bob":   89,
	}

	// Adding and accessing
	scores["Charlie"] = 85
	fmt.Println(scores["Alice"])

	// Check if key exists
	if score, exists := scores["David"]; exists {
		fmt.Println(score)
	}
}

// INTERMEDIATE CONCEPTS
// 1. Structs and Methods
type Person struct {
	Name string
	Age  int
}

// Method
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I'm %d years old", p.Name, p.Age)
}

// 2. Interfaces
type Animal interface {
	Speak() string
