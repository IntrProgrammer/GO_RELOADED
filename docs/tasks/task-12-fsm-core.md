# Task 12 — FSM Core Controller

## Objective
Structure the formatter's decision-making as a finite-state machine (FSM).

## Prerequisites
- Task 11 completed (Rules Inside Quotes)

## Deliverables
- [ ] FSM state definitions
- [ ] State transition logic
- [ ] FSM controller implementation
- [ ] Tests for state transitions

## FSM States
- **READING**: Processing input tokens
- **EVALUATING**: Analyzing current token
- **EDITING**: Applying transformations
- **ERROR**: Handling invalid input

## Implementation Steps

### Step 1: Define FSM States
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

### Step 2: Define FSM Structure
File: `fsm/fsm.go`
```go
package fsm

import "yourmodule/tokenizer"

type FSM struct {
    state       State
    tokens      []tokenizer.Token
    position    int
    result      []tokenizer.Token
    errorMsg    string
}

func New(tokens []tokenizer.Token) *FSM {
    return &FSM{
        state:    StateReading,
        tokens:   tokens,
        position: 0,
        result:   make([]tokenizer.Token, 0, len(tokens)),
    }
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

### Step 3: Write FSM Tests
File: `fsm/fsm_test.go`
```go
package fsm

import (
    "testing"
    "yourmodule/tokenizer"
)

func TestFSMStateTransitions(t *testing.T) {
    tests := []struct {
        name         string
        tokens       []tokenizer.Token
        wantFinalState State
    }{
        {
            name: "simple word processing",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
            },
            wantFinalState: StateDone,
        },
        {
            name: "tag processing",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(up)"},
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
```

### Step 4: Implement FSM Logic
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

func (f *FSM) handleReading() {
    if f.position >= len(f.tokens) {
        f.state = StateDone
        return
    }
    
    f.state = StateEvaluating
}

func (f *FSM) handleEvaluating() {
    token := f.tokens[f.position]
    
    if token.Type == tokenizer.TAG {
        f.state = StateEditing
    } else {
        f.result = append(f.result, token)
        f.position++
        f.state = StateReading
    }
}

func (f *FSM) handleEditing() {
    // Apply transformation based on tag
    // For now, simplified - just skip tag
    f.position++
    f.state = StateReading
}
```

## Verification Commands
```bash
go test ./fsm/... -v
go test ./...
```

## Success Criteria
- FSM transitions correctly between states
- All tokens processed
- No infinite loops
- Error state reachable

## TDD Workflow
1. ✅ RED: Write failing FSM transition tests
2. ✅ GREEN: Implement FSM state machine
3. ✅ REFACTOR: Clean up state handlers

## Git Commit Message
```
feat: implement FSM core controller

- Define FSM states (READING, EVALUATING, EDITING, ERROR, DONE)
- Implement FSM structure and state transitions
- Add Run() method for FSM execution
- Create comprehensive state transition tests
```
