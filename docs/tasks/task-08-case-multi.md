# Task 8 — Multi-Word Case Tags

## Objective
Expand case rules to handle multiple previous words with count parameter.

## Prerequisites
- Task 7 completed (Single-Word Case Tags)

## Deliverables
- [ ] Multi-word case transformation
- [ ] Range safety (handle fewer words than count)
- [ ] Tests for edge cases
- [ ] Documentation of behavior

## Case Rules
- `(up, 3)`: Uppercase 3 previous words
- `(low, 2)`: Lowercase 2 previous words
- `(cap, 4)`: Capitalize 4 previous words
- If fewer words exist, transform all available

## Implementation Steps

### Step 1: Write Multi-Word Tests
File: `formatter/case_test.go` (add to existing)
```go
func TestApplyCaseTransformMulti(t *testing.T) {
    tests := []struct {
        name   string
        tokens []tokenizer.Token
        want   []tokenizer.Token
    }{
        {
            name: "uppercase multiple words",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "world"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(up, 2)"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "HELLO"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "WORLD"},
            },
        },
        {
            name: "capitalize with count larger than words",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(cap, 5)"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "Hello"},
            },
        },
        {
            name: "lowercase three words",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "HELLO"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "BEAUTIFUL"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "WORLD"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(low, 3)"},
            },
            want: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "beautiful"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "world"},
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

### Step 2: Update Case Transform Function
File: `formatter/case.go` (modify existing)
```go
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
        
        // Find previous N words
        wordIndices := []int{}
        for j := len(result) - 1; j >= 0 && len(wordIndices) < tag.Count; j-- {
            if result[j].Type == tokenizer.WORD {
                wordIndices = append(wordIndices, j)
            }
        }
        
        if len(wordIndices) == 0 {
            continue
        }
        
        // Apply transformation to all found words
        for _, idx := range wordIndices {
            switch tag.Command {
            case TagUp:
                result[idx].Value = toUpper(result[idx].Value)
            case TagLow:
                result[idx].Value = toLower(result[idx].Value)
            case TagCap:
                result[idx].Value = capitalize(result[idx].Value)
            }
        }
    }
    
    return result
}
```

## Verification Commands
```bash
go test ./formatter/... -v -run TestApplyCaseTransformMulti
go test ./...
```

## Success Criteria
- Multi-word transformations work correctly
- Safe when count exceeds available words
- No panics or index errors
- All tests pass

## TDD Workflow
1. ✅ RED: Write failing multi-word tests
2. ✅ GREEN: Update ApplyCaseTransform for counts
3. ✅ REFACTOR: Optimize word collection logic

## Git Commit Message
```
feat: add multi-word case transformation support

- Extend ApplyCaseTransform to handle count parameter
- Add range safety for fewer words than count
- Support (up, N), (low, N), (cap, N) tags
- Add comprehensive edge case tests
```
