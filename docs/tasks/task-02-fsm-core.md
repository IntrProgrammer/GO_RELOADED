# Task 2 — FSM Core Structure

## Objective
Build the Finite State Machine (FSM) core that orchestrates token processing through state transitions.

## Prerequisites
- Task 1 completed (Tokenizer)

## Deliverables
- [ ] FSM state definitions
- [ ] FSM structure with state machine logic
- [ ] Processor interface
- [ ] State transition tests

## FSM States
- **READING**: Check if more tokens exist, transition to EVALUATING or DONE
- **EVALUATING**: Examine token type, route to processor or append
- **EDITING**: Apply transformations via processors
- **DONE**: Terminal state, return processed tokens
- **ERROR**: Error handling state

## Implementation Steps

### Step 1: Define States
File: `fsm/state.go`
```go
package fsm

type State int

const (
	StateReading State = iota
	StateEvaluating
	StateEditing
	StateError
	StateDone
)

func (s State) String() string {
	return [...]string{"READING", "EVALUATING", "EDITING", "ERROR", "DONE"}[s]
}
```

### Step 2: Define Processor Interface
File: `fsm/fsm.go`
```go
package fsm

import "GO_RELOADED/tokenizer"

type Processor interface {
	Process(result []tokenizer.Token, currentToken tokenizer.Token) (modified []tokenizer.Token, handled bool)
}
```

### Step 3: Define FSM Structure
File: `fsm/fsm.go` (continued)
```go
type FSM struct {
	state      State
	tokens     []tokenizer.Token
	position   int
	result     []tokenizer.Token
	errorMsg   string
	processors []Processor
	inQuote    bool              // Track quote boundaries
}

func New(tokens []tokenizer.Token) *FSM {
	return &FSM{
		state:      StateReading,
		tokens:     tokens,
		position:   0,
		result:     make([]tokenizer.Token, 0, len(tokens)),
		processors: []Processor{},
		inQuote:    false,
	}
}

func (f *FSM) AddProcessor(p Processor) {
	f.processors = append(f.processors, p)
}

func (f *FSM) CurrentState() State {
	return f.state
}

func (f *FSM) Error() string {
	return f.errorMsg
}

func (f *FSM) Result() []tokenizer.Token {
	return f.result
}
```

### Step 4: Implement State Machine Loop
File: `fsm/fsm.go` (continued)
```go
func (f *FSM) Run() {
	for f.state != StateDone && f.state != StateError {
		f.step()
	}
}

func (f *FSM) step() {
	switch f.state {
	case StateReading:
		f.handleReading()
	case StateEvaluating:
		f.handleEvaluating()
	case StateEditing:
		f.handleEditing()
	case StateError:
		return
	}
}
```

### Step 5: Implement State Handlers
File: `fsm/fsm.go` (continued)
```go
func (f *FSM) handleReading() {
	if f.position >= len(f.tokens) {
		f.state = StateDone
		return
	}
	f.state = StateEvaluating
}

func (f *FSM) handleEvaluating() {
	token := f.tokens[f.position]

	// Track quote boundaries
	if token.Type == tokenizer.QUOTE {
		f.inQuote = !f.inQuote
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

func (f *FSM) handleEditing() {
	tag := f.tokens[f.position]

	// Try processors for tag tokens
	for _, proc := range f.processors {
		if modified, handled := proc.Process(f.result, tag); handled {
			f.result = modified
			f.position++
			f.state = StateReading
			return
		}
	}

	// Default: skip tag
	f.position++
	f.state = StateReading
}
```

### Step 6: Write FSM Tests
File: `fsm/fsm_test.go`
```go
package fsm

import (
	"GO_RELOADED/tokenizer"
	"testing"
)

func TestFSMStateTransitions(t *testing.T) {
	tests := []struct {
		name           string
		tokens         []tokenizer.Token
		wantFinalState State
	}{
		{
			name: "empty tokens",
			tokens: []tokenizer.Token{},
			wantFinalState: StateDone,
		},
		{
			name: "simple word processing",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
			},
			wantFinalState: StateDone,
		},
		{
			name: "multiple tokens",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
			},
			wantFinalState: StateDone,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsm := New(tt.tokens)
			fsm.Run()
			if fsm.CurrentState() != tt.wantFinalState {
				t.Errorf("Final state = %v, want %v", fsm.CurrentState(), tt.wantFinalState)
			}
		})
	}
}

func TestFSMPassthrough(t *testing.T) {
	tokens := []tokenizer.Token{
		{tokenizer.WORD, "hello"},
		{tokenizer.WHITESPACE, " "},
		{tokenizer.WORD, "world"},
	}

	fsm := New(tokens)
	fsm.Run()

	result := fsm.Result()
	if len(result) != len(tokens) {
		t.Errorf("Result length = %d, want %d", len(result), len(tokens))
	}

	for i, token := range result {
		if token != tokens[i] {
			t.Errorf("Token %d = %v, want %v", i, token, tokens[i])
		}
	}
}
```

## Verification Commands
```bash
go test ./fsm/... -v
go test ./...
```

## Success Criteria
- FSM transitions correctly between states
- All tokens processed without processors
- No infinite loops
- Empty input handled correctly
- All tests pass

## TDD Workflow
1. ✅ RED: Write failing FSM tests
2. ✅ GREEN: Implement FSM state machine
3. ✅ REFACTOR: Clean up state handlers

## Git Commit Message
```
feat: implement FSM core structure

- Define FSM states (READING, EVALUATING, EDITING, ERROR, DONE)
- Implement FSM state machine with Run() loop
- Add Processor interface for pluggable transformations
- Create state transition tests
```
