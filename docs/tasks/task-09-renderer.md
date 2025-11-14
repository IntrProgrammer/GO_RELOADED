# Task 9 — Simple Renderer

## Objective
Convert normalized token streams back into readable text.

## Prerequisites
- Task 1 completed (Tokenizer)

## Deliverables
- [ ] Render function
- [ ] Round-trip tests (tokenize → render)
- [ ] Integration tests with normalization

## Rendering Rules
- Concatenate tokens in order
- Respect whitespace tokens
- Produce clean, readable output

## Implementation Steps

### Step 1: Write Renderer Tests
File: `formatter/renderer_test.go`
```go
package formatter

import (
    "testing"
    "yourmodule/tokenizer"
)

func TestRender(t *testing.T) {
    tests := []struct {
        name   string
        tokens []tokenizer.Token
        want   string
    }{
        {
            name: "simple words",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "world"},
            },
            want: "hello world",
        },
        {
            name: "with punctuation",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.PUNCTUATION, ","},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "world"},
                {tokenizer.PUNCTUATION, "!"},
            },
            want: "hello, world!",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Render(tt.tokens)
            if got != tt.want {
                t.Errorf("Render() = %q, want %q", got, tt.want)
            }
        })
    }
}

func TestRoundTrip(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  string
    }{
        {
            name:  "normalize punctuation spacing",
            input: "hello , world !",
            want:  "hello, world!",
        },
        {
            name:  "clean quote spacing",
            input: "' hello world '",
            want:  "'hello world'",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tokens := tokenizer.Tokenize(tt.input)
            tokens = processWithFSM(tokens)
            got := Render(tokens)
            if got != tt.want {
                t.Errorf("Round-trip = %q, want %q", got, tt.want)
            }
        })
    }
}
```

### Step 2: Implement Renderer
File: `formatter/renderer.go`
```go
package formatter

import (
    "strings"
    "yourmodule/tokenizer"
)

func Render(tokens []tokenizer.Token) string {
    var builder strings.Builder
    for _, token := range tokens {
        builder.WriteString(token.Value)
    }
    return builder.String()
}
```

## Verification Commands
```bash
go test ./formatter/... -v -run TestRender
go test ./formatter/... -v -run TestRoundTrip
go test ./...
```

## Success Criteria
- Tokens render to correct text
- Round-trip tests pass
- Integration with normalization works
- No data loss

## TDD Workflow
1. ✅ RED: Write failing render and round-trip tests
2. ✅ GREEN: Implement Render()
3. ✅ REFACTOR: Optimize string building

## Git Commit Message
```
feat: implement token renderer

- Add Render function for token-to-text conversion
- Implement round-trip integration tests
- Verify normalization pipeline works end-to-end
```
