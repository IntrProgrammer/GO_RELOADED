# Task 18 — Integration and Continuous Testing (CI)

## Objective
Ensure everything runs end-to-end automatically and reliably.

## Prerequisites
- All previous tasks completed

## Deliverables
- [ ] GitHub Actions workflow
- [ ] Test data files
- [ ] Coverage reporting
- [ ] Static analysis integration

## CI Pipeline
- Run all tests
- Check code coverage
- Run static analysis
- Build executable

## Implementation Steps

### Step 1: Create Test Data
File: `testdata/input1.txt`
```
hello world (up) this is a test.
```

File: `testdata/expected1.txt`
```
HELLO WORLD this is a test.
```

File: `testdata/input2.txt`
```
a apple and a orange (cap, 2)
```

File: `testdata/expected2.txt`
```
an apple And An orange
```

### Step 2: Write Integration Tests
File: `integration_test.go`
```go
package main

import (
    "os"
    "path/filepath"
    "testing"
    "yourmodule/formatter"
)

func TestIntegrationGoldenFiles(t *testing.T) {
    testCases := []struct {
        input    string
        expected string
    }{
        {"testdata/input1.txt", "testdata/expected1.txt"},
        {"testdata/input2.txt", "testdata/expected2.txt"},
    }
    
    for _, tc := range testCases {
        t.Run(tc.input, func(t *testing.T) {
            input, err := os.ReadFile(tc.input)
            if err != nil {
                t.Fatalf("Failed to read input: %v", err)
            }
            
            expected, err := os.ReadFile(tc.expected)
            if err != nil {
                t.Fatalf("Failed to read expected: %v", err)
            }
            
            f := formatter.New()
            result := f.Format(string(input))
            
            if result != string(expected) {
                t.Errorf("Result mismatch:\nGot:      %q\nExpected: %q", result, string(expected))
            }
        })
    }
}
```

### Step 3: Create GitHub Actions Workflow
File: `.github/workflows/ci.yml`
```yaml
name: CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    name: Test
    runs-on: ubuntu-latest
    
    steps:
    - name: Checkout code
      uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Download dependencies
      run: go mod download
    
    - name: Run tests
      run: go test -v -race -coverprofile=coverage.out ./...
    
    - name: Check coverage
      run: |
        go tool cover -func=coverage.out
        COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
        echo "Total coverage: $COVERAGE%"
        if (( $(echo "$COVERAGE < 70" | bc -l) )); then
          echo "Coverage below 70%"
          exit 1
        fi
    
    - name: Run go vet
      run: go vet ./...
    
    - name: Run go fmt
      run: |
        if [ -n "$(gofmt -l .)" ]; then
          echo "Code is not formatted:"
          gofmt -l .
          exit 1
        fi
    
    - name: Run staticcheck
      run: |
        go install honnef.co/go/tools/cmd/staticcheck@latest
        staticcheck ./...
    
    - name: Build
      run: go build -v ./cmd/go-reloaded

  integration:
    name: Integration Tests
    runs-on: ubuntu-latest
    needs: test
    
    steps:
    - name: Checkout code
      uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Build
      run: go build -o go-reloaded ./cmd/go-reloaded
    
    - name: Run integration tests
      run: |
        ./go-reloaded testdata/input1.txt testdata/output1.txt
        diff testdata/output1.txt testdata/expected1.txt
```

### Step 4: Create Makefile
File: `Makefile`
```makefile
.PHONY: test build clean coverage lint

test:
	go test -v -race ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	go vet ./...
	gofmt -l .
	staticcheck ./...

build:
	go build -o bin/go-reloaded ./cmd/go-reloaded

clean:
	rm -rf bin/ coverage.out coverage.html

install:
	go install ./cmd/go-reloaded

all: lint test build
```

### Step 5: Update README with CI Badge
File: `README.md` (add)
```markdown
# GO_RELOADED

![CI](https://github.com/yourusername/go-reloaded/workflows/CI/badge.svg)
[![Coverage](https://img.shields.io/badge/coverage-80%25-green.svg)]()

Text formatter with tag-based transformations.

## CI/CD

This project uses GitHub Actions for continuous integration:
- Automated testing on every push
- Code coverage reporting (minimum 70%)
- Static analysis with go vet and staticcheck
- Integration tests with golden files

## Running Tests Locally

```bash
make test          # Run all tests
make coverage      # Generate coverage report
make lint          # Run static analysis
make all           # Run everything
```
```

## Verification Commands
```bash
make test
make coverage
make lint
make build
```

## Success Criteria
- CI workflow runs successfully
- All tests pass in CI
- Coverage meets threshold (70%+)
- Static analysis passes
- Integration tests pass

## TDD Workflow
1. ✅ RED: Create CI workflow (will fail initially)
2. ✅ GREEN: Fix any issues found by CI
3. ✅ REFACTOR: Optimize CI pipeline

## Git Commit Message
```
ci: add GitHub Actions workflow and integration tests

- Create CI workflow with test, coverage, and lint jobs
- Add golden file integration tests
- Create Makefile for local testing
- Set coverage threshold at 70%
- Add static analysis with staticcheck
```
