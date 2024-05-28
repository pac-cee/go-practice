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
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return lines, nil
}

// getFileInfo retrieves and displays file information
func getFileInfo(filename string) error {
	info, err := os.Stat(filename)
	if err != nil {
		return fmt.Errorf("error getting file info: %w", err)
	}

	fmt.Printf("File Information for %s:\n", filename)
	fmt.Printf("Size: %d bytes\n", info.Size())
	fmt.Printf("Permissions: %s\n", info.Mode())
	fmt.Printf("Last Modified: %s\n", info.ModTime())
	fmt.Printf("Is Directory: %t\n", info.IsDir())

	return nil
}

func main() {
	// Create a directory for our files
	dataDir := "data"
	if err := createDirectory(dataDir); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
