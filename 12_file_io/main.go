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
