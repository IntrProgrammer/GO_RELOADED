# Task 2 — Punctuation Normalization

## Objective
Clean up punctuation spacing and attachment to words.

## Prerequisites
- Task 1 completed (Tokenizer)

## Deliverables
- [ ] Punctuation normalization function
- [ ] Tests for spacing rules
- [ ] Proper punctuation-to-word binding

## Normalization Rules
- Punctuation attaches to preceding word (no space before)
- Single space after punctuation (except end of text)
- Multiple punctuation marks stay together

## Implementation Steps

### Step 1: Write Normalization Tests
File: `formatter/punctuation_test.go`
```go
package formatter

import (
    "reflect"
    "testing"
    "yourmodule/tokenizer"
)

func TestNormalizePunctuation(t *testing.T) {
    tests := []struct {
        name   string
        tokens []tokenizer.Token
        want   []tokenizer.Token
    }{
        {
            name: "remove space before comma",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.PUNCTUATION, ","},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "world"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.PUNCTUATION, ","},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "world"},
            },
        },
        {
            name: "ensure space after punctuation",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.PUNCTUATION, ","},
                {tokenizer.WORD, "world"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.PUNCTUATION, ","},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "world"},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := NormalizePunctuation(tt.tokens)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NormalizePunctuation() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Step 2: Implement Normalization
File: `formatter/punctuation.go`
```go
package formatter

import "yourmodule/tokenizer"

func NormalizePunctuation(tokens []tokenizer.Token) []tokenizer.Token {
    // Implementation:
    // 1. Remove whitespace before punctuation
    // 2. Ensure whitespace after punctuation (if not end)
    var result []tokenizer.Token
    // Logic here
    return result
}
```

## Verification Commands
```bash
go test ./formatter/... -v
go test ./...
```

## Success Criteria
- Punctuation binds to preceding words
- Correct spacing after punctuation
- No double spaces
- All tests pass

## TDD Workflow
1. ✅ RED: Write failing tests for spacing rules
2. ✅ GREEN: Implement NormalizePunctuation()
3. ✅ REFACTOR: Simplify token manipulation logic

## Git Commit Message
```
feat: implement punctuation normalization

- Add NormalizePunctuation function
- Remove spaces before punctuation
- Ensure spaces after punctuation
- Add comprehensive test coverage
```
