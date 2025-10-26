# Task 15 — Error Handling and Recovery

## Objective
Make the FSM robust against malformed tags or bad conversions.

## Prerequisites
- Task 12 completed (FSM Core)
- Task 14 completed (File I/O and CLI)

## Deliverables
- [ ] ERROR state implementation
- [ ] Error recovery strategies
- [ ] Configuration for fail-fast vs continue
- [ ] Comprehensive error tests

## Error Handling Modes
- **Fail-fast**: Stop on first error
- **Continue**: Log error and continue processing
- **Strict**: Treat warnings as errors

## Implementation Steps

### Step 1: Define Error Types
File: `formatter/errors.go`
```go
package formatter

import "fmt"

type ErrorSeverity int

const (
    SeverityWarning ErrorSeverity = iota
    SeverityError
    SeverityCritical
)

type FormatterError struct {
    Severity ErrorSeverity
    Position int
    Message  string
    Token    string
}

func (e *FormatterError) Error() string {
    return fmt.Sprintf("[%s] at position %d: %s (token: %q)", 
        e.Severity, e.Position, e.Message, e.Token)
}

type ErrorPolicy int

const (
    PolicyFailFast ErrorPolicy = iota
    PolicyContinue
    PolicyStrict
)
```

### Step 2: Write Error Handling Tests
File: `formatter/error_test.go`
```go
package formatter

import (
    "testing"
    "yourmodule/tokenizer"
)

func TestErrorHandling(t *testing.T) {
    tests := []struct {
        name       string
        tokens     []tokenizer.Token
        policy     ErrorPolicy
        wantErr    bool
        wantResult bool
    }{
        {
            name: "invalid hex fail-fast",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "XYZ"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(hex)"},
            },
            policy:     PolicyFailFast,
            wantErr:    true,
            wantResult: false,
        },
        {
            name: "invalid hex continue",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "XYZ"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(hex)"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.WORD, "hello"},
            },
            policy:     PolicyContinue,
            wantErr:    false,
            wantResult: true,
        },
        {
            name: "unknown tag continue",
            tokens: []tokenizer.Token{
                {tokenizer.WORD, "hello"},
                {tokenizer.WHITESPACE, " "},
                {tokenizer.TAG, "(unknown)"},
            },
            policy:     PolicyContinue,
            wantErr:    false,
            wantResult: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            f := NewWithPolicy(tt.policy)
            result, err := f.ProcessTokens(tt.tokens)
            
            if (err != nil) != tt.wantErr {
                t.Errorf("ProcessTokens() error = %v, wantErr %v", err, tt.wantErr)
            }
            
            if tt.wantResult && len(result) == 0 {
                t.Error("Expected result but got empty")
            }
        })
    }
}
```

### Step 3: Implement Error Handling
File: `formatter/formatter.go` (update)
```go
package formatter

import "yourmodule/tokenizer"

type Formatter struct {
    policy ErrorPolicy
    errors []*FormatterError
}

func New() *Formatter {
    return &Formatter{
        policy: PolicyContinue,
        errors: make([]*FormatterError, 0),
    }
}

func NewWithPolicy(policy ErrorPolicy) *Formatter {
    return &Formatter{
        policy: policy,
        errors: make([]*FormatterError, 0),
    }
}

func (f *Formatter) ProcessTokens(tokens []tokenizer.Token) ([]tokenizer.Token, error) {
    result := make([]tokenizer.Token, 0, len(tokens))
    
    for i := 0; i < len(tokens); i++ {
        token := tokens[i]
        
        if token.Type == tokenizer.TAG {
            processed, err := f.processTag(token, result)
            if err != nil {
                formErr := &FormatterError{
                    Severity: SeverityError,
                    Position: i,
                    Message:  err.Error(),
                    Token:    token.Value,
                }
                
                f.errors = append(f.errors, formErr)
                
                if f.policy == PolicyFailFast {
                    return nil, formErr
                }
                // Continue processing
                continue
            }
            result = processed
        } else {
            result = append(result, token)
        }
    }
    
    return result, nil
}

func (f *Formatter) processTag(tag tokenizer.Token, current []tokenizer.Token) ([]tokenizer.Token, error) {
    // Tag processing logic with error returns
    return current, nil
}

func (f *Formatter) Errors() []*FormatterError {
    return f.errors
}

func (f *Formatter) HasErrors() bool {
    return len(f.errors) > 0
}
```

### Step 4: Update FSM Error State
File: `fsm/fsm.go` (update)
```go
func (f *FSM) handleError(err error) {
    f.state = StateError
    f.errorMsg = err.Error()
}
```

## Verification Commands
```bash
go test ./formatter/... -v -run TestErrorHandling
go test ./fsm/... -v
go test ./...
```

## Success Criteria
- Errors caught and logged
- Fail-fast mode stops immediately
- Continue mode processes remaining tokens
- No panics on invalid input

## TDD Workflow
1. ✅ RED: Write failing error handling tests
2. ✅ GREEN: Implement error policies
3. ✅ REFACTOR: Clean up error propagation

## Git Commit Message
```
feat: implement robust error handling and recovery

- Define FormatterError with severity levels
- Add ErrorPolicy (fail-fast, continue, strict)
- Implement error collection and reporting
- Update FSM with ERROR state handling
```
