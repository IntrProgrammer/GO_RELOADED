# Task 6 — Number Conversion Rule

## Objective
Convert hexadecimal or binary numbers to decimal when followed by `(hex)` or `(bin)` tags.

## Prerequisites
- Task 1 completed (Tokenizer)
- Task 5 completed (Tag Parser)

## Deliverables
- [ ] Number conversion functions
- [ ] Tag application for hex/bin
- [ ] Tests for valid and invalid conversions
- [ ] Graceful error handling

## Conversion Rules
- `(hex)`: Convert previous word from hexadecimal to decimal
- `(bin)`: Convert previous word from binary to decimal
- Remove tag after conversion
- Handle invalid numbers gracefully

## Implementation Steps

### Step 1: Write Conversion Tests
File: `formatter/conversion_test.go`
```go
package formatter

import (
    "testing"
    "yourmodule/tokenizer"
)

func TestApplyNumberConversion(t *testing.T) {
    tests := []struct {
        name    string
        tokens  []tokenizer.Token
        want    []tokenizer.Token
        wantErr bool
    }{
        {
            name: "hex conversion",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "1E"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(hex)"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "30"},
            },
        },
        {
            name: "bin conversion",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "101010"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(bin)"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "42"},
            },
        },
        {
            name: "invalid hex",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "XYZ"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(hex)"},
            },
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := ApplyNumberConversion(tt.tokens)
            if (err != nil) != tt.wantErr {
                t.Errorf("ApplyNumberConversion() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr && !tokensEqual(got, tt.want) {
                t.Errorf("ApplyNumberConversion() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Step 2: Implement Conversion Functions
File: `formatter/conversion.go`
```go
package formatter

import (
    "fmt"
    "strconv"
    "yourmodule/tokenizer"
)

func convertHex(s string) (string, error) {
    val, err := strconv.ParseInt(s, 16, 64)
    if err != nil {
        return "", fmt.Errorf("invalid hex: %w", err)
    }
    return strconv.FormatInt(val, 10), nil
}

func convertBin(s string) (string, error) {
    val, err := strconv.ParseInt(s, 2, 64)
    if err != nil {
        return "", fmt.Errorf("invalid binary: %w", err)
    }
    return strconv.FormatInt(val, 10), nil
}

func ApplyNumberConversion(tokens []tokenizer.Token) ([]tokenizer.Token, error) {
    result := make([]tokenizer.Token, 0, len(tokens))
    
    for i := 0; i < len(tokens); i++ {
        if tokens[i].Type != tokenizer.TAG {
            result = append(result, tokens[i])
            continue
        }
        
        tag, err := ParseTag(tokens[i].Value)
        if err != nil {
            result = append(result, tokens[i])
            continue
        }
        
        if tag.Command != TagHex && tag.Command != TagBin {
            result = append(result, tokens[i])
            continue
        }
        
        // Find previous word token
        wordIdx := -1
        for j := len(result) - 1; j >= 0; j-- {
            if result[j].Type == tokenizer.WORD {
                wordIdx = j
                break
            }
        }
        
        if wordIdx == -1 {
            return nil, fmt.Errorf("no word before conversion tag")
        }
        
        // Convert
        var converted string
        if tag.Command == TagHex {
            converted, err = convertHex(result[wordIdx].Value)
        } else {
            converted, err = convertBin(result[wordIdx].Value)
        }
        
        if err != nil {
            return nil, err
        }
        
        result[wordIdx].Value = converted
        // Skip tag (don't append)
    }
    
    return result, nil
}
```

## Verification Commands
```bash
go test ./formatter/... -v -run TestApplyNumberConversion
go test ./...
```

## Success Criteria
- Hex numbers convert correctly
- Binary numbers convert correctly
- Tags removed after conversion
- Invalid numbers return errors
- No panics

## TDD Workflow
1. ✅ RED: Write failing conversion tests
2. ✅ GREEN: Implement conversion functions
3. ✅ REFACTOR: Clean up error handling

## Git Commit Message
```
feat: implement number conversion for hex and bin tags

- Add convertHex and convertBin functions
- Implement ApplyNumberConversion for token streams
- Handle invalid conversions gracefully
- Remove tags after successful conversion
```
