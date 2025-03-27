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
