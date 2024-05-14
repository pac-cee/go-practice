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
var history []Operation

func main() {
	fmt.Println("Welcome to the Go Calculator!")
	fmt.Println("Enter calculations in the format: number operator number")
	fmt.Println("Example: 5 + 3")
	fmt.Println("Available operators: +, -, *, /")
	fmt.Println("Type 'exit' to quit or 'history' to see past calculations")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nEnter calculation: ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		if input == "history" {
			showHistory()
			continue
		}

		result, err := calculate(input)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		fmt.Printf("Result: %.2f\n", result)
	}
}

func calculate(input string) (float64, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid format: please use 'number operator number'")
