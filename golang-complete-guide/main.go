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
