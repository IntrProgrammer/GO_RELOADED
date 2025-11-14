# Task 10 — File I/O and CLI

## Objective
Create command-line interface for processing text files with the formatter.

## Prerequisites
- Task 8 completed (Renderer)

## Deliverables
- [ ] Main executable with CLI
- [ ] Argument validation
- [ ] File I/O operations
- [ ] Error handling

## CLI Interface
```bash
go-reloaded <input> <output>
```

## Implementation Steps

### Step 1: Create Main Entry Point
File: `cmd/go-reloaded/main.go`
```go
package main

import (
	"GO_RELOADED/formatter"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go-reloaded <input> <output>")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	content, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	f := formatter.New()
	result := f.Format(string(content))

	err = os.WriteFile(outputPath, []byte(result), 0644)
	if err != nil {
		log.Fatalf("Error writing output: %v", err)
	}

	fmt.Printf("Successfully processed %s -> %s\n", inputPath, outputPath)
}
```

### Step 2: Build and Test
Create test input file:
```bash
echo "hello (up) world" > test_input.txt
```

Build and run:
```bash
go build -o go-reloaded ./cmd/go-reloaded
./go-reloaded test_input.txt test_output.txt
cat test_output.txt
# Expected: HELLO world
```

## Verification Commands
```bash
go build -o go-reloaded ./cmd/go-reloaded
echo "hello (up) world" > test.txt
./go-reloaded test.txt output.txt
cat output.txt
```

## Success Criteria
- Validates exactly 2 arguments
- Reads input file correctly
- Processes text with formatter
- Writes output file with 0644 permissions
- Prints success message
- Handles errors with log.Fatalf
- Executable builds successfully

## Implementation Notes
- Uses `os.Args` for simple argument parsing
- Uses `os.ReadFile` and `os.WriteFile` from standard library
- Uses `log.Fatalf` for error handling (exits with error message)
- File permissions set to 0644 (readable by all, writable by owner)

## Git Commit Message
```
feat: implement CLI for file processing

- Create main.go with argument validation
- Add file I/O using os.ReadFile/WriteFile
- Integrate formatter for text processing
- Add error handling with log.Fatalf
- Build go-reloaded executable
```
