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
   - Efficient for large files
   - Line-by-line processing

3. **Directory Operations**
   - Create directories
   - List contents
   - Get file information

4. **File Information**
   - Size
   - Permissions
   - Last modified time
   - Directory status

## Running the Code

```bash
go run main.go
```

The program will:

1. Create a data directory
2. Demonstrate various file operations
3. Show file contents and information
4. List directory contents

## Best Practices Demonstrated

- ✅ Resource cleanup with `defer`
- ✅ Proper error handling
- ✅ Cross-platform compatibility
- ✅ Both buffered and unbuffered I/O
- ✅ Modern Go package usage

## Original Tasks (Completed)

- ✅ Append to a file
