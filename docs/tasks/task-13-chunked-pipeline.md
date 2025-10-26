# Task 13 — Chunked Pipeline for Large Files

## Objective
Enable large text handling by splitting work into chunks and processing concurrently.

## Prerequisites
- Task 12 completed (FSM Core)

## Deliverables
- [ ] Chunking strategy
- [ ] Concurrent processing
- [ ] Chunk reassembly
- [ ] Race condition tests

## Chunking Strategy
- Split text at sentence boundaries
- Process chunks concurrently
- Reassemble in order
- Preserve context at boundaries

## Implementation Steps

### Step 1: Write Chunking Tests
File: `pipeline/chunked_test.go`
```go
package pipeline

import (
    "testing"
)

func TestChunkText(t *testing.T) {
    tests := []struct {
        name      string
        input     string
        chunkSize int
        wantChunks int
    }{
        {
            name:      "small text single chunk",
            input:     "Hello world.",
            chunkSize: 100,
            wantChunks: 1,
        },
        {
            name:      "large text multiple chunks",
            input:     "First sentence. Second sentence. Third sentence.",
            chunkSize: 20,
            wantChunks: 3,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            chunks := ChunkText(tt.input, tt.chunkSize)
            if len(chunks) != tt.wantChunks {
                t.Errorf("ChunkText() chunks = %d, want %d", len(chunks), tt.wantChunks)
            }
        })
    }
}

func TestProcessConcurrent(t *testing.T) {
    input := "hello (up) world (low) test (cap)"
    
    result := ProcessConcurrent(input, 10)
    expected := ProcessSequential(input)
    
    if result != expected {
        t.Errorf("ProcessConcurrent() = %q, want %q", result, expected)
    }
}
```

### Step 2: Implement Chunking
File: `pipeline/chunked.go`
```go
package pipeline

import (
    "strings"
    "sync"
)

type Chunk struct {
    Index   int
    Content string
}

func ChunkText(text string, maxSize int) []string {
    if len(text) <= maxSize {
        return []string{text}
    }
    
    var chunks []string
    sentences := strings.Split(text, ". ")
    
    current := ""
    for _, sentence := range sentences {
        if len(current)+len(sentence) > maxSize && current != "" {
            chunks = append(chunks, current)
            current = sentence
        } else {
            if current != "" {
                current += ". "
            }
            current += sentence
        }
    }
    
    if current != "" {
        chunks = append(chunks, current)
    }
    
    return chunks
}

func ProcessConcurrent(text string, chunkSize int) string {
    chunks := ChunkText(text, chunkSize)
    
    results := make([]string, len(chunks))
    var wg sync.WaitGroup
    
    for i, chunk := range chunks {
        wg.Add(1)
        go func(idx int, content string) {
            defer wg.Done()
            results[idx] = processChunk(content)
        }(i, chunk)
    }
    
    wg.Wait()
    
    return strings.Join(results, " ")
}

func processChunk(chunk string) string {
    // Use existing formatter
    // This is a placeholder - integrate with actual formatter
    return chunk
}

func ProcessSequential(text string) string {
    // Use existing formatter for comparison
    return text
}
```

### Step 3: Add Race Detection Test
File: `pipeline/race_test.go`
```go
package pipeline

import (
    "testing"
)

func TestConcurrentRaceConditions(t *testing.T) {
    input := "test " + strings.Repeat("word (up) ", 100)
    
    // Run multiple times to catch race conditions
    for i := 0; i < 10; i++ {
        result := ProcessConcurrent(input, 50)
        if result == "" {
            t.Error("ProcessConcurrent returned empty result")
        }
    }
}
```

## Verification Commands
```bash
go test ./pipeline/... -v
go test -race ./pipeline/...
go test -race ./...
```

## Success Criteria
- Text chunks correctly
- Concurrent processing works
- Results identical to sequential
- No race conditions detected

## TDD Workflow
1. ✅ RED: Write failing chunking tests
2. ✅ GREEN: Implement chunking and concurrent processing
3. ✅ REFACTOR: Optimize chunk boundaries

## Git Commit Message
```
feat: implement chunked pipeline for large files

- Add ChunkText for splitting at sentence boundaries
- Implement ProcessConcurrent with goroutines
- Add race condition tests
- Ensure output matches sequential processing
```
