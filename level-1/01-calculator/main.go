package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Println("\nSimple Calculator")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Print("Choose operation (1-5): ")

		scanner.Scan()
		choice := scanner.Text()

		if choice == "5" {
			fmt.Println("Goodbye!")
			break
		}

		if !isValidChoice(choice) {
			fmt.Println("Invalid choice! Please choose 1-5")
			continue
		}

		fmt.Print("Enter first number: ")
		scanner.Scan()
		num1, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
