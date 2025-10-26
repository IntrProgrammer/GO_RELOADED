# Task 9 — Tag Application Orchestration

## Objective
Coordinate multiple tag applications over a full token stream in correct order.

## Prerequisites
- Task 6 completed (Number Conversion)
- Task 8 completed (Multi-Word Case Tags)

## Deliverables
- [ ] Orchestration function
- [ ] Deterministic tag processing order
- [ ] Safe index management
- [ ] Tests with mixed tags

## Processing Order
1. Number conversions (hex, bin)
2. Case transformations (up, low, cap)
3. Left-to-right processing

## Implementation Steps

### Step 1: Write Orchestration Tests
File: `formatter/orchestrator_test.go`
```go
package formatter

import (
    "testing"
    "yourmodule/tokenizer"
)

func TestApplyAllTags(t *testing.T) {
    tests := []struct {
        name   string
        tokens []tokenizer.Token
        want   []tokenizer.Token
    }{
        {
            name: "mixed hex and case",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "1a"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(hex)"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(up)"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "26"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "HELLO"},
            },
        },
        {
            name: "multiple case tags",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "world"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(up)"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "test"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(cap)"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "WORLD"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "Test"},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := ApplyAllTags(tt.tokens)
            if !tokensEqual(got, tt.want) {
                t.Errorf("ApplyAllTags() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Step 2: Implement Orchestrator
File: `formatter/orchestrator.go`
```go
package formatter

import "yourmodule/tokenizer"

func ApplyAllTags(tokens []tokenizer.Token) []tokenizer.Token {
    result := tokens
    
    // Pass 1: Number conversions
    result = applyNumberConversions(result)
    
    // Pass 2: Case transformations
    result = applyCaseTransformations(result)
    
    return result
}

func applyNumberConversions(tokens []tokenizer.Token) []tokenizer.Token {
    result := make([]tokenizer.Token, 0, len(tokens))
    
    for i := 0; i < len(tokens); i++ {
        if tokens[i].Type != tokenizer.TAG {
            result = append(result, tokens[i])
            continue
        }
        
        tag, err := ParseTag(tokens[i].Value)
        if err != nil || (tag.Command != TagHex && tag.Command != TagBin) {
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
        
        // Convert
        var converted string
        var convErr error
        if tag.Command == TagHex {
            converted, convErr = convertHex(result[wordIdx].Value)
        } else {
            converted, convErr = convertBin(result[wordIdx].Value)
        }
        
        if convErr == nil {
            result[wordIdx].Value = converted
        }
        // Skip tag
    }
    
    return result
}

func applyCaseTransformations(tokens []tokenizer.Token) []tokenizer.Token {
    return ApplyCaseTransform(tokens)
}
```

## Verification Commands
```bash
go test ./formatter/... -v -run TestApplyAllTags
go test ./...
```

## Success Criteria
- Tags processed in correct order
- No interference between tag types
- Index management is safe
- Mixed tag scenarios work correctly

## TDD Workflow
1. ✅ RED: Write failing orchestration tests
2. ✅ GREEN: Implement ApplyAllTags
3. ✅ REFACTOR: Optimize multi-pass logic

## Git Commit Message
```
feat: implement tag application orchestration

- Add ApplyAllTags coordinator function
- Process number conversions before case changes
- Ensure deterministic left-to-right processing
- Handle mixed tag scenarios correctly
```
