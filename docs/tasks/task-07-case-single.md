# Task 7 — Case-Change Single-Word Tags

## Objective
Handle case transformation tags `(up)`, `(low)`, `(cap)` applied to one preceding word.

## Prerequisites
- Task 5 completed (Tag Parser)

## Deliverables
- [ ] Case transformation functions
- [ ] Single-word tag application
- [ ] Tests for all case types
- [ ] Punctuation preservation

## Case Rules
- `(up)`: Convert word to UPPERCASE
- `(low)`: Convert word to lowercase
- `(cap)`: Capitalize first letter only
- Preserve attached punctuation

## Implementation Steps

### Step 1: Write Case Transformation Tests
File: `formatter/case_test.go`
```go
package formatter

import (
    "testing"
    "yourmodule/tokenizer"
)

func TestApplyCaseTransform(t *testing.T) {
    tests := []struct {
        name   string
        tokens []tokenizer.Token
        want   []tokenizer.Token
    }{
        {
            name: "uppercase single word",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(up)"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "HELLO"},
            },
        },
        {
            name: "lowercase single word",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "HELLO"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(low)"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
            },
        },
        {
            name: "capitalize single word",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(cap)"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "Hello"},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := ApplyCaseTransform(tt.tokens)
            if !tokensEqual(got, tt.want) {
                t.Errorf("ApplyCaseTransform() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Step 2: Implement Case Functions
File: `formatter/case.go`
```go
package formatter

import (
    "strings"
    "unicode"
    "yourmodule/tokenizer"
)

func toUpper(s string) string {
    return strings.ToUpper(s)
}

func toLower(s string) string {
    return strings.ToLower(s)
}

func capitalize(s string) string {
    if len(s) == 0 {
        return s
    }
    runes := []rune(s)
    runes[0] = unicode.ToUpper(runes[0])
    for i := 1; i < len(runes); i++ {
        runes[i] = unicode.ToLower(runes[i])
    }
    return string(runes)
}

func ApplyCaseTransform(tokens []tokenizer.Token) []tokenizer.Token {
    result := make([]tokenizer.Token, 0, len(tokens))
    
    for i := 0; i < len(tokens); i++ {
        if tokens[i].Type != tokenizer.TAG {
            result = append(result, tokens[i])
            continue
        }
        
        tag, err := ParseTag(tokens[i].Value)
        if err != nil || (tag.Command != TagUp && tag.Command != TagLow && tag.Command != TagCap) {
            result = append(result, tokens[i])
            continue
        }
        
        // Find previous word
        wordIdx := -1
        for j := len(result) - 1; j >= 0; j-- {
            if result[j].Type == tokenizer.WORD {
                wordIdx = j
                break
            }
        }
        
        if wordIdx == -1 {
            continue
        }
        
        // Apply transformation
        switch tag.Command {
        case TagUp:
            result[wordIdx].Value = toUpper(result[wordIdx].Value)
        case TagLow:
            result[wordIdx].Value = toLower(result[wordIdx].Value)
        case TagCap:
            result[wordIdx].Value = capitalize(result[wordIdx].Value)
        }
    }
    
    return result
}
```

## Verification Commands
```bash
go test ./formatter/... -v -run TestApplyCaseTransform
go test ./...
```

## Success Criteria
- All case transformations work correctly
- Tags removed after application
- Punctuation preserved
- No panics on edge cases

## TDD Workflow
1. ✅ RED: Write failing case transformation tests
2. ✅ GREEN: Implement case functions
3. ✅ REFACTOR: Optimize string operations

## Git Commit Message
```
feat: implement single-word case transformations

- Add toUpper, toLower, capitalize functions
- Implement ApplyCaseTransform for single words
- Support (up), (low), (cap) tags
- Preserve punctuation attachment
```
