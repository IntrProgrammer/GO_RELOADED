# Task 10 — "a" → "an" Rule

## Objective
Correct grammatical article usage: "a" before consonants, "an" before vowels/h.

## Prerequisites
- Task 1 completed (Tokenizer)

## Deliverables
- [ ] Article correction function
- [ ] Vowel/h detection logic
- [ ] Tests for edge cases
- [ ] Skip punctuation-only tokens

## Correction Rules
- "a" → "an" before vowels (a, e, i, o, u)
- "a" → "an" before silent h
- "A" → "An" (preserve capitalization)
- Skip punctuation when looking ahead

## Implementation Steps

### Step 1: Write Article Correction Tests
File: `formatter/article_test.go`
```go
package formatter

import (
    "testing"
    "yourmodule/tokenizer"
)

func TestCorrectArticles(t *testing.T) {
    tests := []struct {
        name   string
        tokens []tokenizer.Token
        want   []tokenizer.Token
    }{
        {
            name: "a before vowel",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "a"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "apple"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "an"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "apple"},
            },
        },
        {
            name: "a before consonant unchanged",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "a"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "book"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "a"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "book"},
            },
        },
        {
            name: "capital A before vowel",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "A"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "elephant"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "An"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "elephant"},
            },
        },
        {
            name: "a before h",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "a"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "hour"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "an"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "hour"},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := CorrectArticles(tt.tokens)
            if !tokensEqual(got, tt.want) {
                t.Errorf("CorrectArticles() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Step 2: Implement Article Correction
File: `formatter/article.go`
```go
package formatter

import (
    "strings"
    "unicode"
    "yourmodule/tokenizer"
)

func startsWithVowelOrH(s string) bool {
    if len(s) == 0 {
        return false
    }
    first := unicode.ToLower(rune(s[0]))
    return first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u' || first == 'h'
}

func CorrectArticles(tokens []tokenizer.Token) []tokenizer.Token {
    result := make([]tokenizer.Token, len(tokens))
    copy(result, tokens)
    
    for i := 0; i < len(result); i++ {
        if result[i].Type != tokenizer.WORD {
            continue
        }
        
        word := result[i].Value
        if strings.ToLower(word) != "a" {
            continue
        }
        
        // Find next word token
        nextWordIdx := -1
        for j := i + 1; j < len(result); j++ {
            if result[j].Type == tokenizer.WORD {
                nextWordIdx = j
                break
            }
        }
        
        if nextWordIdx == -1 {
            continue
        }
        
        // Check if next word starts with vowel or h
        if startsWithVowelOrH(result[nextWordIdx].Value) {
            if word == "a" {
                result[i].Value = "an"
            } else if word == "A" {
                result[i].Value = "An"
            }
        }
    }
    
    return result
}
```

## Verification Commands
```bash
go test ./formatter/... -v -run TestCorrectArticles
go test ./...
```

## Success Criteria
- "a" → "an" before vowels
- "a" → "an" before h
- Capitalization preserved
- Consonants unchanged
- Punctuation skipped correctly

## TDD Workflow
1. ✅ RED: Write failing article tests
2. ✅ GREEN: Implement CorrectArticles
3. ✅ REFACTOR: Optimize vowel detection

## Git Commit Message
```
feat: implement a/an article correction

- Add startsWithVowelOrH detection function
- Implement CorrectArticles for grammatical correctness
- Preserve capitalization (A → An)
- Skip punctuation when looking ahead
```
