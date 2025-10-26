# Task A — AI-Assisted Test Generation (Optional)

## Objective
Enrich test coverage through diverse auto-generated test cases.

## Prerequisites
- Task 17 completed (Property Tests)
- Task 18 completed (CI Integration)

## Deliverables
- [ ] AI-generated test dataset
- [ ] Test data validation
- [ ] Integration with test suite
- [ ] Documentation of generation process

## Generation Strategy
- Use AI to generate edge cases
- Create diverse input scenarios
- Generate expected outputs
- Validate against formatter

## Implementation Steps

### Step 1: Create Test Generator Script
File: `scripts/generate_tests.go`
```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type TestCase struct {
    Name     string `json:"name"`
    Input    string `json:"input"`
    Expected string `json:"expected"`
    Tags     []string `json:"tags"`
}

func main() {
    testCases := []TestCase{
        {
            Name:     "complex nested tags",
            Input:    "hello world (up) test (cap) value",
            Expected: "hello WORLD Test value",
            Tags:     []string{"case", "multiple"},
        },
        {
            Name:     "mixed conversions and case",
            Input:    "FF (hex) value (up)",
            Expected: "255 VALUE",
            Tags:     []string{"conversion", "case"},
        },
        {
            Name:     "article with multiple words",
            Input:    "a elephant and a apple",
            Expected: "an elephant and an apple",
            Tags:     []string{"grammar"},
        },
        {
            Name:     "quotes with transformations",
            Input:    "'a apple (up)'",
            Expected: "'an APPLE'",
            Tags:     []string{"quotes", "grammar", "case"},
        },
        {
            Name:     "edge case empty tag",
            Input:    "hello () world",
            Expected: "hello () world",
            Tags:     []string{"edge"},
        },
    }
    
    data, err := json.MarshalIndent(testCases, "", "  ")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
    
    err = os.WriteFile("testdata/ai_generated_tests.json", data, 0644)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Println("Generated", len(testCases), "test cases")
}
```

### Step 2: Create Test Loader
File: `formatter/ai_test.go`
```go
package formatter

import (
    "encoding/json"
    "os"
    "testing"
)

type AITestCase struct {
    Name     string   `json:"name"`
    Input    string   `json:"input"`
    Expected string   `json:"expected"`
    Tags     []string `json:"tags"`
}

func TestAIGeneratedCases(t *testing.T) {
    data, err := os.ReadFile("../testdata/ai_generated_tests.json")
    if err != nil {
        t.Skip("AI test data not found")
        return
    }
    
    var testCases []AITestCase
    err = json.Unmarshal(data, &testCases)
    if err != nil {
        t.Fatalf("Failed to parse test data: %v", err)
    }
    
    for _, tc := range testCases {
        t.Run(tc.Name, func(t *testing.T) {
            f := New()
            result := f.Format(tc.Input)
            
            if result != tc.Expected {
                t.Errorf("Input: %q\nGot:      %q\nExpected: %q\nTags: %v",
                    tc.Input, result, tc.Expected, tc.Tags)
            }
        })
    }
}
```

### Step 3: Create Validation Script
File: `scripts/validate_ai_tests.go`
```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
    "yourmodule/formatter"
)

type TestCase struct {
    Name     string   `json:"name"`
    Input    string   `json:"input"`
    Expected string   `json:"expected"`
    Tags     []string `json:"tags"`
}

func main() {
    data, err := os.ReadFile("testdata/ai_generated_tests.json")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
        os.Exit(1)
    }
    
    var testCases []TestCase
    err = json.Unmarshal(data, &testCases)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
        os.Exit(1)
    }
    
    f := formatter.New()
    passed := 0
    failed := 0
    
    for _, tc := range testCases {
        result := f.Format(tc.Input)
        if result == tc.Expected {
            passed++
            fmt.Printf("✓ %s\n", tc.Name)
        } else {
            failed++
            fmt.Printf("✗ %s\n", tc.Name)
            fmt.Printf("  Input:    %q\n", tc.Input)
            fmt.Printf("  Expected: %q\n", tc.Expected)
            fmt.Printf("  Got:      %q\n", result)
        }
    }
    
    fmt.Printf("\nResults: %d passed, %d failed\n", passed, failed)
    
    if failed > 0 {
        os.Exit(1)
    }
}
```

### Step 4: Create AI Prompt Template
File: `docs/AI_TEST_GENERATION.md`
```markdown
# AI Test Generation Guide

## Prompt Template

Use this prompt with AI tools to generate test cases:

```
Generate 20 diverse test cases for a text formatter with these features:
- Case transformations: (up), (low), (cap) with optional count
- Number conversions: (hex), (bin)
- Grammar: a/an correction
- Quote processing: rules apply inside quotes
- Punctuation normalization

For each test case provide:
1. Descriptive name
2. Input text
3. Expected output
4. Tags (categories)

Format as JSON array with fields: name, input, expected, tags

Focus on:
- Edge cases
- Complex combinations
- Boundary conditions
- Error scenarios
- Unicode and special characters
```

## Example Output

```json
[
  {
    "name": "multiple tags in sequence",
    "input": "hello (up) world (low) test (cap)",
    "expected": "HELLO world Test",
    "tags": ["case", "multiple"]
  }
]
```

## Validation Process

1. Generate tests with AI
2. Save to `testdata/ai_generated_tests.json`
3. Run validation: `go run scripts/validate_ai_tests.go`
4. Fix any failing tests
5. Integrate into test suite
```

### Step 5: Update Makefile
File: `Makefile` (add)
```makefile
.PHONY: generate-tests validate-ai-tests

generate-tests:
	go run scripts/generate_tests.go

validate-ai-tests:
	go run scripts/validate_ai_tests.go

ai-tests: generate-tests validate-ai-tests
	go test ./formatter/... -v -run TestAIGeneratedCases
```

### Step 6: Create Sample AI Dataset
File: `testdata/ai_generated_tests.json`
```json
[
  {
    "name": "unicode with transformations",
    "input": "héllo wörld (up)",
    "expected": "héllo WÖRLD",
    "tags": ["unicode", "case"]
  },
  {
    "name": "multiple articles",
    "input": "a apple a elephant a orange",
    "expected": "an apple an elephant an orange",
    "tags": ["grammar", "multiple"]
  },
  {
    "name": "nested quotes with tags",
    "input": "'hello (up) 'nested' world'",
    "expected": "'HELLO 'nested' world'",
    "tags": ["quotes", "nested", "case"]
  },
  {
    "name": "large hex number",
    "input": "FFFF (hex)",
    "expected": "65535",
    "tags": ["conversion", "boundary"]
  },
  {
    "name": "binary with leading zeros",
    "input": "00001010 (bin)",
    "expected": "10",
    "tags": ["conversion", "edge"]
  }
]
```

## Verification Commands
```bash
make generate-tests
make validate-ai-tests
make ai-tests
go test ./formatter/... -v -run TestAIGeneratedCases
```

## Success Criteria
- AI test generator works
- Validation script runs
- Generated tests integrate with suite
- Documentation complete
- Tests find new edge cases

## TDD Workflow
1. ✅ RED: Generate tests that may fail
2. ✅ GREEN: Fix formatter to pass tests
3. ✅ REFACTOR: Improve test generation

## Git Commit Message
```
test: add AI-assisted test generation

- Create test generator script
- Add test validation tool
- Integrate AI-generated tests into suite
- Document generation process
- Add sample AI-generated test dataset
```
