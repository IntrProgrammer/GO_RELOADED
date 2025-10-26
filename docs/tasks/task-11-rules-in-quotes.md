# Task 11 — Rules Inside Quotes

## Objective
Ensure all transformations apply equivalently within quoted content.

## Prerequisites
- Task 3 completed (Quote Spacing)
- Task 9 completed (Tag Orchestration)
- Task 10 completed (A/An Rule)

## Deliverables
- [ ] Quote-aware processing
- [ ] Apply all rules inside quotes
- [ ] Preserve quote boundaries
- [ ] Tests for quoted transformations

## Processing Rules
- Detect quoted sections
- Apply all transformations inside quotes
- Keep quote marks intact
- Handle nested/multiple quote pairs

## Implementation Steps

### Step 1: Write Quote Processing Tests
File: `formatter/quote_rules_test.go`
```go
package formatter

import (
    "testing"
    "yourmodule/tokenizer"
)

func TestProcessQuotedContent(t *testing.T) {
    tests := []struct {
        name   string
        tokens []tokenizer.Token
        want   []tokenizer.Token
    }{
        {
            name: "case transform inside quotes",
            tokens: []tokenizer.Token{
                {tokenizer.QUOTE, "'"},
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(up)"},
                {tokenizer.QUOTE, "'"},
            },
            want: []tokenizer.Token{
                {tokenizer.QUOTE, "'"},
                {tokenizer.WORD, "HELLO"},
                {tokenizer.QUOTE, "'"},
            },
        },
        {
            name: "article correction inside quotes",
            tokens: []tokenizer.Token{
                {tokenizer.QUOTE, "'"},
                {tokenizer.WORD, "a"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "apple"},
                {tokenizer.QUOTE, "'"},
            },
            want: []tokenizer.Token{
                {tokenizer.QUOTE, "'"},
                {tokenizer.WORD, "an"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "apple"},
                {tokenizer.QUOTE, "'"},
            },
        },
        {
            name: "hex conversion inside quotes",
            tokens: []tokenizer.Token{
                {tokenizer.QUOTE, "'"},
                {tokenizer.WORD, "FF"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(hex)"},
                {tokenizer.QUOTE, "'"},
            },
            want: []tokenizer.Token{
                {tokenizer.QUOTE, "'"},
                {tokenizer.WORD, "255"},
                {tokenizer.QUOTE, "'"},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := ProcessWithQuotes(tt.tokens)
            if !tokensEqual(got, tt.want) {
                t.Errorf("ProcessWithQuotes() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Step 2: Implement Quote-Aware Processing
File: `formatter/quote_rules.go`
```go
package formatter

import "yourmodule/tokenizer"

func ProcessWithQuotes(tokens []tokenizer.Token) []tokenizer.Token {
    result := make([]tokenizer.Token, 0, len(tokens))
    inQuote := false
    quoteStart := 0
    
    for i := 0; i < len(tokens); i++ {
        if tokens[i].Type == tokenizer.QUOTE {
            if !inQuote {
                // Opening quote
                inQuote = true
                quoteStart = len(result)
                result = append(result, tokens[i])
            } else {
                // Closing quote - process content between quotes
                quotedContent := result[quoteStart+1:]
                
                // Apply all transformations
                quotedContent = ApplyAllTags(quotedContent)
                quotedContent = CorrectArticles(quotedContent)
                
                // Rebuild result with processed content
                result = result[:quoteStart+1]
                result = append(result, quotedContent...)
                result = append(result, tokens[i])
                
                inQuote = false
            }
        } else {
            result = append(result, tokens[i])
        }
    }
    
    // If not in quotes, apply transformations to entire stream
    if !inQuote {
        result = ApplyAllTags(result)
        result = CorrectArticles(result)
    }
    
    return result
}
```

### Step 3: Update Main Processing Pipeline
File: `formatter/formatter.go`
```go
package formatter

import "yourmodule/tokenizer"

func (f *Formatter) Format(input string) string {
    tokens := tokenizer.Tokenize(input)
    tokens = NormalizePunctuation(tokens)
    tokens = CleanQuoteSpacing(tokens)
    tokens = ProcessWithQuotes(tokens)
    return Render(tokens)
}
```

## Verification Commands
```bash
go test ./formatter/... -v -run TestProcessQuotedContent
go test ./...
```

## Success Criteria
- All rules apply inside quotes
- Quote marks preserved
- Transformations work correctly
- No quote balance issues

## TDD Workflow
1. ✅ RED: Write failing quote processing tests
2. ✅ GREEN: Implement ProcessWithQuotes
3. ✅ REFACTOR: Optimize quote detection

## Git Commit Message
```
feat: apply all rules inside quoted content

- Add ProcessWithQuotes function
- Apply tags and article corrections in quotes
- Preserve quote boundaries
- Integrate with main formatting pipeline
```
