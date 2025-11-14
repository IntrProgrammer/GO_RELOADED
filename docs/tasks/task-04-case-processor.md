# Task 4 — Case Processor (up/low/cap)

## Objective
Implement processor to handle case transformations with support for single and multiple words.

## Prerequisites
- Task 3 completed (Conversion Processor)

## Deliverables
- [ ] CaseProcessor implementation
- [ ] Single-word case transformations
- [ ] Multi-word case transformations with count
- [ ] Tests for all case types

## Case Rules
- `(up)` or `(up, N)`: Convert N previous words to UPPERCASE
- `(low)` or `(low, N)`: Convert N previous words to lowercase
- `(cap)` or `(cap, N)`: Capitalize N previous words (first letter upper, rest lower)
- Default count is 1 if not specified
- If fewer words exist than count, transform all available

## Implementation Steps

### Step 1: Write Case Transformation Tests
File: `fsm/processors_test.go` (add to existing)
```go
func TestCaseProcessor(t *testing.T) {
	tests := []struct {
		name    string
		result  []tokenizer.Token
		token   tokenizer.Token
		want    []tokenizer.Token
		handled bool
	}{
		{
			name: "uppercase single word",
			result: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
			},
			token: tokenizer.Token{tokenizer.TAG, "(up)"},
			want: []tokenizer.Token{
				{tokenizer.WORD, "HELLO"},
			},
			handled: true,
		},
		{
			name: "lowercase single word",
			result: []tokenizer.Token{
				{tokenizer.WORD, "HELLO"},
				{tokenizer.WHITESPACE, " "},
			},
			token: tokenizer.Token{tokenizer.TAG, "(low)"},
			want: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
			},
			handled: true,
		},
		{
			name: "capitalize single word",
			result: []tokenizer.Token{
				{tokenizer.WORD, "hELLo"},
				{tokenizer.WHITESPACE, " "},
			},
			token: tokenizer.Token{tokenizer.TAG, "(cap)"},
			want: []tokenizer.Token{
				{tokenizer.WORD, "Hello"},
			},
			handled: true,
		},
		{
			name: "uppercase multiple words",
			result: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
				{tokenizer.WHITESPACE, " "},
			},
			token: tokenizer.Token{tokenizer.TAG, "(up, 2)"},
			want: []tokenizer.Token{
				{tokenizer.WORD, "HELLO"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "WORLD"},
			},
			handled: true,
		},
		{
			name: "capitalize with count larger than words",
			result: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
			},
			token: tokenizer.Token{tokenizer.TAG, "(cap, 5)"},
			want: []tokenizer.Token{
				{tokenizer.WORD, "Hello"},
			},
			handled: true,
		},
		{
			name: "non-case tag",
			result: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
			},
			token: tokenizer.Token{tokenizer.TAG, "(hex)"},
			want: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
			},
			handled: false,
		},
	}

	proc := &CaseProcessor{}
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

### Step 2: Implement Capitalize Helper
File: `fsm/processors.go` (add to existing)
```go
import (
	"unicode"
)

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
```

### Step 3: Implement CaseProcessor
File: `fsm/processors.go` (add to existing)
```go
type CaseProcessor struct{}

func (p *CaseProcessor) Process(result []tokenizer.Token, currentToken tokenizer.Token) ([]tokenizer.Token, bool) {
	if currentToken.Type != tokenizer.TAG {
		return result, false
	}

	tag := parseTag(currentToken.Value)
	if tag == nil || (tag.command != "up" && tag.command != "low" && tag.command != "cap") {
		return result, false
	}

	// Find previous N WORD tokens (stop at quote boundaries)
	wordIndices := []int{}
	for j := len(result) - 1; j >= 0 && len(wordIndices) < tag.count; j-- {
		if result[j].Type == tokenizer.QUOTE {
			break
		}
		if result[j].Type == tokenizer.WORD {
			wordIndices = append(wordIndices, j)
		}
	}

	if len(wordIndices) == 0 {
		return result, true
	}

	// Apply transformation
	modified := make([]tokenizer.Token, len(result))
	copy(modified, result)

	for _, idx := range wordIndices {
		switch tag.command {
		case "up":
			modified[idx].Value = strings.ToUpper(modified[idx].Value)
		case "low":
			modified[idx].Value = strings.ToLower(modified[idx].Value)
		case "cap":
			modified[idx].Value = capitalize(modified[idx].Value)
		}
	}

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
func TestFSMWithCaseProcessor(t *testing.T) {
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
			name: "capitalize multiple words",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(cap, 2)"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "Hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "World"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsm := New(tt.tokens)
			fsm.AddProcessor(&CaseProcessor{})
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
go test ./fsm/... -v -run TestCaseProcessor
go test ./fsm/... -v -run TestFSMWithCaseProcessor
go test ./...
```

## Success Criteria
- Single-word transformations work (up, low, cap)
- Multi-word transformations work with count
- Count parameter parsed correctly
- Safe when count exceeds available words
- Trailing whitespace removed
- All tests pass

## TDD Workflow
1. ✅ RED: Write failing case transformation tests
2. ✅ GREEN: Implement CaseProcessor
3. ✅ REFACTOR: Optimize word collection logic

## Git Commit Message
```
feat: implement case processor for up/low/cap tags

- Add CaseProcessor with single and multi-word support
- Implement capitalize helper function
- Support count parameter for multi-word transformations
- Handle edge cases (count > available words)
```
