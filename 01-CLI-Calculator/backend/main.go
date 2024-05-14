package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Operation represents a calculation operation
type Operation struct {
	FirstNumber  float64
	SecondNumber float64
	Operator     string
	Result       float64
}

// History stores the last 5 calculations
