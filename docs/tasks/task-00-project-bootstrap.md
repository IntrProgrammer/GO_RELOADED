# Task 0 — Project Bootstrap

## Objective
Establish a runnable, testable Go project skeleton.

## Prerequisites
- None (first task)

## Deliverables
- [ ] Initialized Go module
- [ ] Package `formatter` with minimal constructor
- [ ] Basic passing test confirming setup

## Implementation Steps

### Step 1: Initialize Go Module
```bash
go mod init github.com/yourusername/go-reloaded
```

### Step 2: Create Package Structure
Create directory: `formatter/`

### Step 3: Create Formatter Constructor
File: `formatter/formatter.go`
```go
package formatter

type Formatter struct{}

func New() *Formatter {
    return &Formatter{}
}
```

### Step 4: Create Basic Test
File: `formatter/formatter_test.go`
```go
package formatter

import "testing"

func TestNew(t *testing.T) {
    f := New()
    if f == nil {
        t.Fatal("New() returned nil")
    }
}
```

## Verification Commands
```bash
go test ./...
go vet ./...
go fmt ./...
go build ./...
```

## Success Criteria
- All tests pass
- No vet warnings
- Code builds without errors
- Module initialized correctly

## TDD Workflow
1. ✅ RED: Write failing test (TestNew expects non-nil)
2. ✅ GREEN: Implement New() to return &Formatter{}
3. ✅ REFACTOR: Clean up if needed

## Git Commit Message
```
feat: initialize project with formatter package

- Initialize Go module
- Create formatter package with New() constructor
- Add basic test coverage
```
