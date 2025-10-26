# Task 14 — File I/O and CLI

## Objective
Provide a usable command-line interface for real text input/output.

## Prerequisites
- Task 13 completed (Chunked Pipeline)

## Deliverables
- [ ] File reader/writer
- [ ] CLI argument parsing
- [ ] Main executable
- [ ] Usage documentation

## CLI Interface
```bash
go-reloaded input.txt output.txt
go-reloaded -chunk-size=1000 input.txt output.txt
```

## Implementation Steps

### Step 1: Write File I/O Tests
File: `io/fileio_test.go`
```go
package io

import (
    "os"
    "path/filepath"
    "testing"
)

func TestReadFile(t *testing.T) {
    tmpDir := t.TempDir()
    testFile := filepath.Join(tmpDir, "test.txt")
    
    content := "hello world"
    err := os.WriteFile(testFile, []byte(content), 0644)
    if err != nil {
        t.Fatal(err)
    }
    
    got, err := ReadFile(testFile)
    if err != nil {
        t.Fatalf("ReadFile() error = %v", err)
    }
    
    if got != content {
        t.Errorf("ReadFile() = %q, want %q", got, content)
    }
}

func TestWriteFile(t *testing.T) {
    tmpDir := t.TempDir()
    testFile := filepath.Join(tmpDir, "output.txt")
    
    content := "hello world"
    err := WriteFile(testFile, content)
    if err != nil {
        t.Fatalf("WriteFile() error = %v", err)
    }
    
    got, err := os.ReadFile(testFile)
    if err != nil {
        t.Fatal(err)
    }
    
    if string(got) != content {
        t.Errorf("File content = %q, want %q", string(got), content)
    }
}
```

### Step 2: Implement File I/O
File: `io/fileio.go`
```go
package io

import (
    "fmt"
    "os"
)

func ReadFile(path string) (string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return "", fmt.Errorf("failed to read file: %w", err)
    }
    return string(data), nil
}

func WriteFile(path string, content string) error {
    err := os.WriteFile(path, []byte(content), 0644)
    if err != nil {
        return fmt.Errorf("failed to write file: %w", err)
    }
    return nil
}
```

### Step 3: Create CLI
File: `cmd/go-reloaded/main.go`
```go
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "yourmodule/formatter"
    "yourmodule/io"
)

func main() {
    chunkSize := flag.Int("chunk-size", 0, "Chunk size for large files (0 = no chunking)")
    flag.Parse()
    
    args := flag.Args()
    if len(args) != 2 {
        fmt.Println("Usage: go-reloaded [options] <input> <output>")
        fmt.Println("Options:")
        flag.PrintDefaults()
        os.Exit(1)
    }
    
    inputPath := args[0]
    outputPath := args[1]
    
    // Read input
    content, err := io.ReadFile(inputPath)
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }
    
    // Process
    f := formatter.New()
    result := f.Format(content)
    
    // Write output
    err = io.WriteFile(outputPath, result)
    if err != nil {
        log.Fatalf("Error writing output: %v", err)
    }
    
    fmt.Printf("Successfully processed %s -> %s\n", inputPath, outputPath)
}
```

### Step 4: Create Test Files
File: `testdata/sample_input.txt`
```
hello world (up) this is a test.
```

File: `testdata/sample_expected.txt`
```
HELLO WORLD this is a test.
```

### Step 5: Integration Test
File: `cmd/go-reloaded/main_test.go`
```go
package main

import (
    "os"
    "os/exec"
    "path/filepath"
    "testing"
)

func TestCLIIntegration(t *testing.T) {
    tmpDir := t.TempDir()
    inputFile := filepath.Join(tmpDir, "input.txt")
    outputFile := filepath.Join(tmpDir, "output.txt")
    
    input := "hello (up)"
    err := os.WriteFile(inputFile, []byte(input), 0644)
    if err != nil {
        t.Fatal(err)
    }
    
    cmd := exec.Command("go", "run", "main.go", inputFile, outputFile)
    err = cmd.Run()
    if err != nil {
        t.Fatalf("CLI execution failed: %v", err)
    }
    
    output, err := os.ReadFile(outputFile)
    if err != nil {
        t.Fatal(err)
    }
    
    if len(output) == 0 {
        t.Error("Output file is empty")
    }
}
```

## Verification Commands
```bash
go test ./io/... -v
go test ./cmd/... -v
go build ./cmd/go-reloaded
./go-reloaded testdata/sample_input.txt testdata/output.txt
```

## Success Criteria
- Files read/write correctly
- CLI parses arguments
- Executable builds
- Integration test passes

## TDD Workflow
1. ✅ RED: Write failing I/O tests
2. ✅ GREEN: Implement file operations and CLI
3. ✅ REFACTOR: Clean up error handling

## Git Commit Message
```
feat: implement file I/O and CLI interface

- Add ReadFile and WriteFile functions
- Create CLI with flag parsing
- Build main executable
- Add integration tests and sample files
```
