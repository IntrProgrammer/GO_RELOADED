# Task 3 — Single-Quote Inner-Spacing Cleaner

## Objective
Remove extra spaces inside single quotes while preserving content.

## Prerequisites
- Task 1 completed (Tokenizer)
- Task 2 completed (Punctuation Normalization)

## Deliverables
- [ ] Quote spacing cleaner function
- [ ] Tests for quote boundary handling
- [ ] Support for nested/multiple quotes

## Spacing Rules
- Remove space after opening quote `'`
- Remove space before closing quote `'`
- Preserve all content between quotes
- Handle unbalanced quotes gracefully

## Implementation Steps

### Step 1: Write Quote Spacing Tests
File: `formatter/quote_test.go`
```go
package formatter

import (
    "reflect"
    "testing"
    "yourmodule/tokenizer"
)

func TestCleanQuoteSpacing(t *testing.T) {
    tests := []struct {
        name   string
        tokens []tokenizer.Token
        want   []tokenizer.Token
    }{
        {
            name: "remove space after opening quote",
            tokens: []tokenizer.Token{
                {tokenizer.QUOTE, "'"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "hello"},
                {tokenizer.QUOTE, "'"},
            },
            want: []tokenizer.Token{
                {tokenizer.QUOTE, "'"},
                {tokenizer.WORD, "hello"},
                {tokenizer.QUOTE, "'"},
            },
        },
        {
            name: "remove space before closing quote",
            tokens: []tokenizer.Token{
                {tokenizer.QUOTE, "'"},
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.QUOTE, "'"},
            },
            want: []tokenizer.Token{
                {tokenizer.QUOTE, "'"},
                {tokenizer.WORD, "hello"},
                {tokenizer.QUOTE, "'"},
            },
        },
        {
            name: "preserve internal spacing",
            tokens: []tokenizer.Token{
                {tokenizer.QUOTE, "'"},
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "world"},
                {tokenizer.QUOTE, "'"},
            },
            want: []tokenizer.Token{
                {tokenizer.QUOTE, "'"},
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "world"},
                {tokenizer.QUOTE, "'"},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := CleanQuoteSpacing(tt.tokens)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("CleanQuoteSpacing() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Step 2: Implement Quote Spacing Cleaner
File: `formatter/quote.go`
```go
package formatter

import "yourmodule/tokenizer"

func CleanQuoteSpacing(tokens []tokenizer.Token) []tokenizer.Token {
    // Implementation:
    // 1. Track quote state (inside/outside)
    // 2. Remove whitespace immediately after opening quote
    // 3. Remove whitespace immediately before closing quote
    var result []tokenizer.Token
    // Logic here
    return result
}
```

## Verification Commands
```bash
go test ./formatter/... -v -run TestCleanQuoteSpacing
go test ./...
```

## Success Criteria
- No extra space after opening quote
- No extra space before closing quote
- Internal spacing preserved
- Handles multiple quote pairs
- All tests pass

## TDD Workflow
1. ✅ RED: Write failing tests for quote spacing
2. ✅ GREEN: Implement CleanQuoteSpacing()
3. ✅ REFACTOR: Optimize quote state tracking

## Git Commit Message
```
feat: implement quote spacing cleaner

- Add CleanQuoteSpacing function
- Remove boundary spaces in quotes
- Preserve internal content spacing
- Handle multiple quote pairs
```
