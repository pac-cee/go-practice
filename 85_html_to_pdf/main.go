package main

import (
	"fmt"
	"os"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input.html> [output.pdf]")
		return
	}
	input := os.Args[1]
	output := "output.pdf"
	if len(os.Args) > 2 {
		output = os.Args[2]
	}
	f, err := os.Open(input)
