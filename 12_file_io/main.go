package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// createDirectory creates a new directory and returns an error if it fails
func createDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}

// writeFileBuffered writes data to a file using buffered I/O
func writeFileBuffered(filename string, data []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range data {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing to buffer: %w", err)
		}
	}

	return writer.Flush()
}

// readFileBuffered reads a file line by line using buffered I/O
func readFileBuffered(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
