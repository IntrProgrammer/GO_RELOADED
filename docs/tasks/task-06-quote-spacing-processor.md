# Task 7 — Quote Spacing Processor (Rule 3)

## Objective
Implement QuoteSpacingProcessor to clean whitespace inside quotes: remove whitespace after opening quotes and before closing quotes.

## Prerequisites
- Task 6 completed (Formatter Integration)

## Deliverables
- [ ] QuoteSpacingProcessor implementation
- [ ] FSM quote handling updates
- [ ] Processor registration in formatter
- [ ] Tests for quote spacing

## Rule 3: Quote Spacing
- Remove whitespace immediately after opening quote (`'` or `"`)
- Remove whitespace immediately before closing quote
- Apply all other rules inside quotes

**Examples**:
```
Input:  ' hello world '
Output: 'hello world'

Input:  ' hello (up) '
Output: 'HELLO'
```

## Implementation Steps

### Step 1: Write QuoteSpacingProcessor Tests
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
			name: "closing quote removes multiple whitespaces",
			result: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
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
			name:    "non-quote token not handled",
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

### Step 2: Implement QuoteSpacingProcessor
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

### Step 3: Update FSM Quote Handling
File: `fsm/fsm.go` (update handleEvaluating)

Replace the quote handling section:
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

	// ... rest of handler remains unchanged
}
```

### Step 4: Register Processor in Formatter
File: `formatter/formatter.go` (update Format method)
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

**Important**: QuoteSpacingProcessor must be first to clean spacing before other transformations.

### Step 5: Write Integration Tests
File: `formatter/formatter_test.go` (add to existing)
```go
func TestQuoteSpacingIntegration(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "basic quote spacing",
			input: "' hello world '",
			want:  "'hello world'",
		},
		{
			name:  "double quotes",
			input: `" hello world "`,
			want:  `"hello world"`,
		},
		{
			name:  "quote spacing with transformation",
			input: "' hello (up) '",
			want:  "'HELLO'",
		},
		{
			name:  "multiple spaces",
			input: "'   hello   '",
			want:  "'hello'",
		},
		{
			name:  "no spaces",
			input: "'hello'",
			want:  "'hello'",
		},
		{
			name:  "multiple quote pairs",
			input: "' first ' and ' second '",
			want:  "'first' and 'second'",
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
go test ./formatter/... -v -run TestQuoteSpacingIntegration
go test ./...
```

## Success Criteria
- QuoteSpacingProcessor handles opening quotes
- QuoteSpacingProcessor removes trailing whitespace before closing quotes
- FSM skips whitespace after opening quotes
- Processor registered as first in formatter
- Quote spacing works with transformations
- All tests pass

## TDD Workflow
1. ✅ RED: Write failing QuoteSpacingProcessor tests
2. ✅ GREEN: Implement QuoteSpacingProcessor
3. ✅ REFACTOR: Update FSM quote handling
4. ✅ RED: Write failing integration tests
5. ✅ GREEN: Register processor in formatter
6. ✅ REFACTOR: Verify all tests pass

## Git Commit Message
```
feat: implement quote spacing processor (Rule 3)

- Add QuoteSpacingProcessor to clean whitespace inside quotes
- Remove whitespace after opening quotes
- Remove whitespace before closing quotes
- Update FSM to skip whitespace after opening quotes
- Register QuoteSpacingProcessor as first processor
- Add comprehensive quote spacing tests
```
