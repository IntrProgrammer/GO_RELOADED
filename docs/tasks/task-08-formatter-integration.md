# Task 6 — Quote Spacing & Formatter Integration

## Objective
Implement quote spacing cleanup (Rule 3) as FSM processor and integrate FSM into formatter for single-pass text processing.

## Prerequisites
- Task 5 completed (Article Correction)

## Deliverables
- [ ] QuoteSpacingProcessor implementation
- [ ] FSM quote handling updates
- [ ] Formatter.Format() method
- [ ] Quote spacing tests

## Rule 3: Quote Spacing
- Remove whitespace after opening quote (`'` or `"`)
- Remove whitespace before closing quote
- Apply all other rules inside quotes

**Example**: `' hello world '` → `'hello world'`

## Implementation Steps

### Step 1: Implement QuoteSpacingProcessor
File: `fsm/processors.go` (add before startsWithVowelOrH)
```go
// QuoteSpacingProcessor handles quote spacing cleanup (Rule 3)
type QuoteSpacingProcessor struct{}

func (p *QuoteSpacingProcessor) Process(result []tokenizer.Token, currentToken tokenizer.Token) ([]tokenizer.Token, bool) {
	if currentToken.Type != tokenizer.QUOTE {
		return result, false
	}

	// Check if this is a closing quote (we have tokens in result)
	if len(result) == 0 {
		// Opening quote - just append
		return append(result, currentToken), true
	}

	// Check if last token in result is also a quote (opening quote case)
	if result[len(result)-1].Type == tokenizer.QUOTE {
		// This is right after opening quote, just append
		return append(result, currentToken), true
	}

	// This is a closing quote - remove trailing whitespace before it
	modified := make([]tokenizer.Token, len(result))
	copy(modified, result)

	for len(modified) > 0 && modified[len(modified)-1].Type == tokenizer.WHITESPACE {
		modified = modified[:len(modified)-1]
	}

	modified = append(modified, currentToken)
	return modified, true
}
```

### Step 2: Update FSM Quote Handling
File: `fsm/fsm.go` (update handleEvaluating)
```go
func (f *FSM) handleEvaluating() {
	token := f.tokens[f.position]

	// Track quote boundaries and handle spacing
	if token.Type == tokenizer.QUOTE {
		wasInQuote := f.inQuote
		f.inQuote = !f.inQuote
		
		// Use QuoteSpacingProcessor to handle quote spacing
		for _, proc := range f.processors {
			if modified, handled := proc.Process(f.result, token); handled {
				f.result = modified
				f.position++
				
				// Skip whitespace after opening quote
				if !wasInQuote && f.position < len(f.tokens) && f.tokens[f.position].Type == tokenizer.WHITESPACE {
					f.position++
				}
				
				f.state = StateReading
				return
			}
		}
		
		// Fallback if no processor handled it
		f.result = append(f.result, token)
		f.position++
		f.state = StateReading
		return
	}

	if token.Type == tokenizer.TAG {
		f.state = StateEditing
		return
	}

	// Try processors for non-tag tokens
	for _, proc := range f.processors {
		if modified, handled := proc.Process(f.result, token); handled {
			f.result = modified
			f.position++
			f.state = StateReading
			return
		}
	}

	// Default: append token
	f.result = append(f.result, token)
	f.position++
	f.state = StateReading
}
```

### Step 3: Write Quote Spacing Tests
File: `fsm/processors_test.go` (add to existing)
```go
func TestQuoteSpacingProcessor(t *testing.T) {
	tests := []struct {
		name    string
		result  []tokenizer.Token
		token   tokenizer.Token
		want    []tokenizer.Token
		handled bool
	}{
		{
			name:   "opening quote",
			result: []tokenizer.Token{},
			token:  tokenizer.Token{tokenizer.QUOTE, "'"},
			want: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
			},
			handled: true,
		},
		{
			name: "closing quote removes trailing whitespace",
			result: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
			},
			token: tokenizer.Token{tokenizer.QUOTE, "'"},
			want: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "hello"},
				{tokenizer.QUOTE, "'"},
			},
			handled: true,
		},
		{
			name:    "non-quote token",
			result:  []tokenizer.Token{},
			token:   tokenizer.Token{tokenizer.WORD, "hello"},
			want:    []tokenizer.Token{},
			handled: false,
		},
	}

	proc := &QuoteSpacingProcessor{}
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

### Step 4: Create Formatter Structure
File: `formatter/formatter.go`
```go
package formatter

import (
	"GO_RELOADED/fsm"
	"GO_RELOADED/tokenizer"
)

type Formatter struct{}

func New() *Formatter {
	return &Formatter{}
}
```

### Step 2: Implement Format Method
File: `formatter/formatter.go` (continued)
```go
func (f *Formatter) Format(input string) string {
	tokens := tokenizer.Tokenize(input)
	machine := fsm.New(tokens)
	machine.AddProcessor(&fsm.QuoteSpacingProcessor{})  // FIRST: Clean quote spacing
	machine.AddProcessor(&fsm.ConversionProcessor{})
	machine.AddProcessor(&fsm.CaseProcessor{})
	machine.Run()
	return Render(machine.Result())
}
```

### Step 5: Write Integration Tests
File: `formatter/formatter_test.go`
```go
package formatter

import (
	"testing"
)

func TestNew(t *testing.T) {
	f := New()
	if f == nil {
		t.Fatal("New() returned nil")
	}
}

func TestFormatWithQuoteBoundaries(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "quote spacing basic",
			input: `' hello world '`,
			want:  `'hello world'`,
		},
		{
			name:  "quote spacing with transformation",
			input: `' hello (up) '`,
			want:  `'HELLO'`,
		},
		{
			name:  "hex inside quotes",
			input: `"1E (hex) files"`,
			want:  `"30 files"`,
		},
		{
			name:  "hex outside quotes",
			input: `"hello" 1E (hex)`,
			want:  `"hello" 30`,
		},
		{
			name:  "case inside quotes",
			input: `'hello (up) world'`,
			want:  `'HELLO world'`,
		},
		{
			name:  "case outside quotes",
			input: `"hello" world (up)`,
			want:  `"hello" WORLD`,
		},
		{
			name:  "multi-word case respects quotes",
			input: `"one two" three four (up, 2)`,
			want:  `"one two" THREE FOUR`,
		},
		{
			name:  "article correction",
			input: "a apple",
			want:  "an apple",
		},
		{
			name:  "article correction in quotes",
			input: `'a elephant'`,
			want:  `'an elephant'`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := New()
			got := f.Format(tt.input)
			if got != tt.want {
				t.Errorf("Format() = %q, want %q", got, tt.want)
			}
		})
	}
}
```

## Verification Commands
```bash
go test ./fsm/... -v -run TestQuoteSpacingProcessor
go test ./formatter/... -v
go test ./...
```

## Success Criteria
- QuoteSpacingProcessor removes whitespace correctly
- FSM skips whitespace after opening quote
- FSM removes whitespace before closing quote
- Formatter integrates all processors
- Quote spacing works with transformations
- Single-pass processing maintained
- All tests pass

## TDD Workflow
1. ✅ RED: Write failing integration tests
2. ✅ GREEN: Implement Format() method
3. ✅ REFACTOR: Clean up integration

## Git Commit Message
```
feat: implement quote spacing and integrate FSM into formatter

- Add QuoteSpacingProcessor for Rule 3 (quote spacing cleanup)
- Update FSM to handle quote spacing with processor
- Skip whitespace after opening quotes
- Remove whitespace before closing quotes
- Create Formatter struct with Format() method
- Integrate all processors (QuoteSpacing, Conversion, Case)
- Add comprehensive quote spacing and integration tests
```
