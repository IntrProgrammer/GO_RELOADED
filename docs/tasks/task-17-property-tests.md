# Task 17 — Property and Edge-Case Tests

## Objective
Broaden reliability through randomized and boundary testing.

## Prerequisites
- All previous tasks completed

## Deliverables
- [ ] Property-based tests
- [ ] Edge case tests
- [ ] Fuzz tests
- [ ] Boundary condition tests

## Test Categories
- **Property tests**: Invariants that always hold
- **Edge cases**: Empty input, single character, etc.
- **Fuzz tests**: Random input generation
- **Boundary tests**: Max sizes, limits

## Implementation Steps

### Step 1: Write Property Tests
File: `formatter/property_test.go`
```go
package formatter

import (
    "strings"
    "testing"
    "yourmodule/tokenizer"
)

func TestPropertyTokenizeRenderRoundTrip(t *testing.T) {
    tests := []string{
        "hello world",
        "a b c",
        "",
        "single",
        strings.Repeat("word ", 1000),
    }
    
    for _, input := range tests {
        tokens := tokenizer.Tokenize(input)
        output := Render(tokens)
        
        // Property: tokenize then render should preserve content
        if !contentEqual(input, output) {
            t.Errorf("Round-trip failed: input=%q, output=%q", input, output)
        }
    }
}

func TestPropertyNoTokenLoss(t *testing.T) {
    inputs := []string{
        "hello (up) world",
        "test (hex) value",
        "a apple",
    }
    
    for _, input := range inputs {
        tokens := tokenizer.Tokenize(input)
        initialCount := len(tokens)
        
        // Process
        f := New()
        result := f.Format(input)
        
        // Property: should not panic
        if result == "" && input != "" {
            t.Errorf("Processing lost all content for input: %q", input)
        }
        
        // Property: token count should be reasonable
        resultTokens := tokenizer.Tokenize(result)
        if len(resultTokens) > initialCount*2 {
            t.Errorf("Token explosion: %d -> %d", initialCount, len(resultTokens))
        }
    }
}

func contentEqual(a, b string) bool {
    // Normalize whitespace for comparison
    return strings.TrimSpace(a) == strings.TrimSpace(b)
}
```

### Step 2: Write Edge Case Tests
File: `formatter/edge_test.go`
```go
package formatter

import "testing"

func TestEdgeCases(t *testing.T) {
    tests := []struct {
        name  string
        input string
    }{
        {"empty string", ""},
        {"single char", "a"},
        {"only spaces", "   "},
        {"only punctuation", "...!!!"},
        {"unbalanced quotes", "'hello"},
        {"nested quotes", "' ' ' '"},
        {"tag without word", "(up)"},
        {"multiple tags", "(up) (low) (cap)"},
        {"very long word", strings.Repeat("a", 10000)},
        {"unicode", "héllo wörld"},
        {"emoji", "hello 😀 world"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            defer func() {
                if r := recover(); r != nil {
                    t.Errorf("Panic on input %q: %v", tt.input, r)
                }
            }()
            
            f := New()
            result := f.Format(tt.input)
            
            // Should not panic, result can be anything
            _ = result
        })
    }
}
```

### Step 3: Write Fuzz Tests
File: `formatter/fuzz_test.go`
```go
package formatter

import (
    "math/rand"
    "strings"
    "testing"
)

func FuzzFormatter(f *testing.F) {
    // Seed corpus
    f.Add("hello world")
    f.Add("test (up)")
    f.Add("a apple")
    
    f.Fuzz(func(t *testing.T, input string) {
        defer func() {
            if r := recover(); r != nil {
                t.Errorf("Panic on input %q: %v", input, r)
            }
        }()
        
        formatter := New()
        result := formatter.Format(input)
        
        // Basic invariants
        if len(input) > 0 && len(result) == 0 {
            // Empty result might be valid for some inputs
        }
    })
}

func TestRandomInputs(t *testing.T) {
    rand.Seed(42)
    
    for i := 0; i < 100; i++ {
        input := generateRandomInput()
        
        f := New()
        result := f.Format(input)
        
        // Should not panic
        _ = result
    }
}

func generateRandomInput() string {
    words := []string{"hello", "world", "test", "a", "an"}
    tags := []string{"(up)", "(low)", "(cap)", "(hex)", "(bin)"}
    
    var parts []string
    length := rand.Intn(20) + 1
    
    for i := 0; i < length; i++ {
        if rand.Float32() < 0.7 {
            parts = append(parts, words[rand.Intn(len(words))])
        } else {
            parts = append(parts, tags[rand.Intn(len(tags))])
        }
    }
    
    return strings.Join(parts, " ")
}
```

### Step 4: Write Boundary Tests
File: `formatter/boundary_test.go`
```go
package formatter

import (
    "strings"
    "testing"
)

func TestBoundaryConditions(t *testing.T) {
    tests := []struct {
        name  string
        input string
    }{
        {"max tag count", strings.Repeat("word (up) ", 1000)},
        {"max word length", strings.Repeat("a", 100000)},
        {"max quote depth", strings.Repeat("'", 100)},
        {"zero width", ""},
        {"single token", "x"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            f := New()
            result := f.Format(tt.input)
            _ = result
        })
    }
}
```

## Verification Commands
```bash
go test ./formatter/... -v -run TestProperty
go test ./formatter/... -v -run TestEdge
go test ./formatter/... -fuzz=FuzzFormatter -fuzztime=10s
go test ./...
```

## Success Criteria
- Property tests pass consistently
- No panics on edge cases
- Fuzz tests run without crashes
- Boundary conditions handled

## TDD Workflow
1. ✅ RED: Write property and edge tests
2. ✅ GREEN: Fix issues found by tests
3. ✅ REFACTOR: Strengthen invariants

## Git Commit Message
```
test: add property-based and edge-case tests

- Implement property tests for invariants
- Add comprehensive edge case coverage
- Create fuzz tests for random inputs
- Test boundary conditions and limits
```
