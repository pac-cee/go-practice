# File I/O in Go

A comprehensive example of file operations in Go, demonstrating various I/O techniques and best practices.

## Features

✨ **File Operations**

- Direct file reading and writing
- Buffered I/O operations
- Line-by-line file processing
- File appending capabilities

🛠️ **Directory Management**

- Directory creation and listing
- Cross-platform path handling
- File metadata retrieval

⚡ **Performance Optimizations**

- Buffered I/O for large files
- Efficient memory usage
- Resource cleanup with defer

🔒 **Error Handling**

- Comprehensive error checking
- Proper error wrapping
- Descriptive error messages

## Code Structure

The project contains two main files:

- `main.go`: Implementation of file operations
- `README.md`: Documentation and usage guide

## Implementation Details

### Core Functions

```go
createDirectory(path string) error
writeFileBuffered(filename string, data []string) error
readFileBuffered(filename string) ([]string, error)
getFileInfo(filename string) error
```

### Features Implemented

1. **Direct File Operations**
   - Using `os.WriteFile` and `os.ReadFile`
   - Unbuffered operations for small files

2. **Buffered Operations**
   - Using `bufio` package
