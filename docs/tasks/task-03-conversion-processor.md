# Task 3 — Conversion Processor (hex/bin)

## Objective
Implement processor to convert hexadecimal and binary numbers to decimal when followed by `(hex)` or `(bin)` tags.

## Prerequisites
- Task 2 completed (FSM Core)

## Deliverables
- [ ] ConversionProcessor implementation
- [ ] Tag parsing helper
- [ ] Hex/bin conversion logic
- [ ] Tests for valid and invalid conversions

## Conversion Rules
- `(hex)`: Convert previous WORD token from hexadecimal to decimal
- `(bin)`: Convert previous WORD token from binary to decimal
- Remove tag and trailing whitespace after conversion
- Handle invalid numbers gracefully (skip conversion)

## Implementation Steps

### Step 1: Write Conversion Tests
File: `fsm/processors_test.go`
```go
package fsm

import (
	"GO_RELOADED/tokenizer"
	"reflect"
	"testing"
)

func TestConversionProcessor(t *testing.T) {
	tests := []struct {
		name   string
		result []tokenizer.Token
		token  tokenizer.Token
		want   []tokenizer.Token
		handled bool
	}{
		{
			name: "hex conversion",
			result: []tokenizer.Token{
				{tokenizer.WORD, "1E"},
				{tokenizer.WHITESPACE, " "},
			},
			token: tokenizer.Token{tokenizer.TAG, "(hex)"},
			want: []tokenizer.Token{
				{tokenizer.WORD, "30"},
			},
			handled: true,
		},
		{
			name: "bin conversion",
			result: []tokenizer.Token{
				{tokenizer.WORD, "101010"},
				{tokenizer.WHITESPACE, " "},
			},
			token: tokenizer.Token{tokenizer.TAG, "(bin)"},
			want: []tokenizer.Token{
				{tokenizer.WORD, "42"},
			},
			handled: true,
		},
		{
			name: "invalid hex",
			result: []tokenizer.Token{
				{tokenizer.WORD, "XYZ"},
				{tokenizer.WHITESPACE, " "},
			},
			token: tokenizer.Token{tokenizer.TAG, "(hex)"},
			want: []tokenizer.Token{
				{tokenizer.WORD, "XYZ"},
				{tokenizer.WHITESPACE, " "},
			},
			handled: true,
		},
		{
			name: "non-conversion tag",
			result: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
			},
			token: tokenizer.Token{tokenizer.TAG, "(up)"},
			want: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
			},
			handled: false,
		},
	}

	proc := &ConversionProcessor{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, handled := proc.Process(tt.result, tt.token)
			if handled != tt.handled {
				t.Errorf("handled = %v, want %v", handled, tt.handled)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

### Step 2: Implement Tag Parser Helper
File: `fsm/processors.go`
```go
package fsm

import (
	"GO_RELOADED/tokenizer"
	"strconv"
	"strings"
)

type tag struct {
	command string
	count   int
}

func parseTag(input string) *tag {
	if !strings.HasPrefix(input, "(") || !strings.HasSuffix(input, ")") {
		return nil
	}

	content := strings.Trim(input, "()")
	parts := strings.Split(content, ",")

	command := strings.TrimSpace(parts[0])
	count := 1

	if len(parts) == 2 {
		if c, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil && c > 0 {
			count = c
		}
	}

	return &tag{command: command, count: count}
}
```

### Step 3: Implement ConversionProcessor
File: `fsm/processors.go` (continued)
```go
type ConversionProcessor struct{}

func (p *ConversionProcessor) Process(result []tokenizer.Token, currentToken tokenizer.Token) ([]tokenizer.Token, bool) {
	if currentToken.Type != tokenizer.TAG {
		return result, false
	}

	tag := parseTag(currentToken.Value)
	if tag == nil || (tag.command != "hex" && tag.command != "bin") {
		return result, false
	}

	// Find previous WORD token (stop at quote boundaries)
	wordIdx := -1
	for j := len(result) - 1; j >= 0; j-- {
		if result[j].Type == tokenizer.QUOTE {
			break
		}
		if result[j].Type == tokenizer.WORD {
			wordIdx = j
			break
		}
	}

	if wordIdx == -1 {
		return result, true
	}

	// Convert
	var converted string
	var err error
	if tag.command == "hex" {
		val, e := strconv.ParseInt(result[wordIdx].Value, 16, 64)
		if e == nil {
			converted = strconv.FormatInt(val, 10)
		}
		err = e
	} else {
		val, e := strconv.ParseInt(result[wordIdx].Value, 2, 64)
		if e == nil {
			converted = strconv.FormatInt(val, 10)
		}
		err = e
	}

	if err != nil {
		return result, true
	}

	// Modify result
	modified := make([]tokenizer.Token, len(result))
	copy(modified, result)
	modified[wordIdx].Value = converted

	// Remove trailing whitespace
	if len(modified) > 0 && modified[len(modified)-1].Type == tokenizer.WHITESPACE {
		modified = modified[:len(modified)-1]
	}

	return modified, true
}
```

### Step 4: Integration Test
File: `fsm/fsm_test.go` (add to existing)
```go
func TestFSMWithConversionProcessor(t *testing.T) {
	tests := []struct {
		name   string
		tokens []tokenizer.Token
		want   []tokenizer.Token
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsm := New(tt.tokens)
			fsm.AddProcessor(&ConversionProcessor{})
			fsm.Run()
			
			got := fsm.Result()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Result = %v, want %v", got, tt.want)
			}
		})
	}
}
```

## Verification Commands
```bash
go test ./fsm/... -v -run TestConversionProcessor
go test ./fsm/... -v -run TestFSMWithConversionProcessor
go test ./...
```

## Success Criteria
- Hex numbers convert correctly (1E → 30, FF → 255)
- Binary numbers convert correctly (101010 → 42)
- Tags removed after conversion
- Trailing whitespace removed
- Invalid numbers handled gracefully
- All tests pass

## TDD Workflow
1. ✅ RED: Write failing conversion tests
2. ✅ GREEN: Implement ConversionProcessor
3. ✅ REFACTOR: Clean up conversion logic

## Git Commit Message
```
feat: implement conversion processor for hex/bin tags

- Add ConversionProcessor with hex and bin support
- Implement parseTag helper for tag parsing
- Handle invalid conversions gracefully
- Remove tags and trailing whitespace after conversion
```
