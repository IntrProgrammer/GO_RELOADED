# Task 1 — Tokenizer

## Objective
Convert raw text into structured tokens for downstream processing.

## Prerequisites
- Task 0 completed

## Deliverables
- [ ] Token type definitions
- [ ] Tokenizer function
- [ ] Table-driven tests for tokenization
- [ ] Documentation of token categories

## Token Categories
- **WORD**: Alphanumeric sequences
- **PUNCTUATION**: `.`, `,`, `!`, `?`, `;`, `:`
- **TAG**: Parenthesized commands like `(up)`, `(hex, 2)`
- **QUOTE**: Single quote `'`
- **WHITESPACE**: Spaces (may be implicit)

## Implementation Steps

### Step 1: Define Token Type
File: `tokenizer/token.go`
```go
package tokenizer

type TokenType int

const (
    WORD TokenType = iota
    PUNCTUATION
    TAG
    QUOTE
    WHITESPACE
)

type Token struct {
    Type  TokenType
    Value string
}
```

### Step 2: Write Tokenizer Tests
File: `tokenizer/tokenizer_test.go`
```go
package tokenizer

import (
    "reflect"
    "testing"
)

func TestTokenize(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  []Token
    }{
        {
            name:  "simple words",
            input: "hello world",
            want: []Token{
                {WORD, "hello"},
                {WHITESPACE, " "},
                {WORD, "world"},
            },
        },
        {
            name:  "word with punctuation",
            input: "hello, world!",
            want: []Token{
                {WORD, "hello"},
                {PUNCTUATION, ","},
                {WHITESPACE, " "},
                {WORD, "world"},
                {PUNCTUATION, "!"},
            },
        },
        {
            name:  "with tag",
            input: "word (up)",
            want: []Token{
                {WORD, "word"},
                {WHITESPACE, " "},
                {TAG, "(up)"},
            },
        },
        {
            name:  "with quotes",
            input: "'hello world'",
            want: []Token{
                {QUOTE, "'"},
                {WORD, "hello"},
                {WHITESPACE, " "},
                {WORD, "world"},
                {QUOTE, "'"},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Tokenize(tt.input)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Tokenize() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Step 3: Implement Tokenizer
File: `tokenizer/tokenizer.go`
```go
package tokenizer

import (
    "regexp"
    "strings"
)

func Tokenize(input string) []Token {
    // Implementation here
    var tokens []Token
    // Parse input and create tokens
    return tokens
}
```

## Verification Commands
```bash
go test ./tokenizer/... -v
go test ./... 
```

## Success Criteria
- All tokenizer tests pass
- Handles words, punctuation, tags, quotes
- Preserves token order
- No data loss during tokenization

## TDD Workflow
1. ✅ RED: Write table-driven tests (all fail initially)
2. ✅ GREEN: Implement Tokenize() to pass tests
3. ✅ REFACTOR: Optimize regex/parsing logic

## Git Commit Message
```
feat: implement tokenizer with token type definitions

- Define Token and TokenType structures
- Implement Tokenize() function
- Add comprehensive table-driven tests
- Support words, punctuation, tags, quotes
```
