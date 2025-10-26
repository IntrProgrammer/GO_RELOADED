# Task 16 — Logging and Diagnostics

## Objective
Add transparency to the system's runtime behavior through logging.

## Prerequisites
- Task 15 completed (Error Handling)

## Deliverables
- [ ] Logger interface
- [ ] Logging implementation
- [ ] Configurable log levels
- [ ] Integration with formatter

## Log Levels
- **DEBUG**: Detailed state transitions
- **INFO**: General processing information
- **WARN**: Non-critical issues
- **ERROR**: Critical failures

## Implementation Steps

### Step 1: Define Logger Interface
File: `logger/logger.go`
```go
package logger

type Level int

const (
    LevelDebug Level = iota
    LevelInfo
    LevelWarn
    LevelError
)

func (l Level) String() string {
    return [...]string{"DEBUG", "INFO", "WARN", "ERROR"}[l]
}

type Logger interface {
    Debug(msg string, args ...interface{})
    Info(msg string, args ...interface{})
    Warn(msg string, args ...interface{})
    Error(msg string, args ...interface{})
    SetLevel(level Level)
}
```

### Step 2: Implement Default Logger
File: `logger/default.go`
```go
package logger

import (
    "fmt"
    "log"
    "os"
)

type DefaultLogger struct {
    level  Level
    logger *log.Logger
}

func New() *DefaultLogger {
    return &DefaultLogger{
        level:  LevelInfo,
        logger: log.New(os.Stdout, "", log.LstdFlags),
    }
}

func (l *DefaultLogger) SetLevel(level Level) {
    l.level = level
}

func (l *DefaultLogger) Debug(msg string, args ...interface{}) {
    if l.level <= LevelDebug {
        l.log(LevelDebug, msg, args...)
    }
}

func (l *DefaultLogger) Info(msg string, args ...interface{}) {
    if l.level <= LevelInfo {
        l.log(LevelInfo, msg, args...)
    }
}

func (l *DefaultLogger) Warn(msg string, args ...interface{}) {
    if l.level <= LevelWarn {
        l.log(LevelWarn, msg, args...)
    }
}

func (l *DefaultLogger) Error(msg string, args ...interface{}) {
    if l.level <= LevelError {
        l.log(LevelError, msg, args...)
    }
}

func (l *DefaultLogger) log(level Level, msg string, args ...interface{}) {
    formatted := fmt.Sprintf(msg, args...)
    l.logger.Printf("[%s] %s", level, formatted)
}
```

### Step 3: Write Logger Tests
File: `logger/logger_test.go`
```go
package logger

import (
    "bytes"
    "log"
    "strings"
    "testing"
)

func TestLoggerLevels(t *testing.T) {
    var buf bytes.Buffer
    logger := &DefaultLogger{
        level:  LevelInfo,
        logger: log.New(&buf, "", 0),
    }
    
    logger.Debug("debug message")
    logger.Info("info message")
    logger.Warn("warn message")
    logger.Error("error message")
    
    output := buf.String()
    
    if strings.Contains(output, "debug message") {
        t.Error("Debug message should not appear at INFO level")
    }
    
    if !strings.Contains(output, "info message") {
        t.Error("Info message should appear at INFO level")
    }
    
    if !strings.Contains(output, "warn message") {
        t.Error("Warn message should appear at INFO level")
    }
    
    if !strings.Contains(output, "error message") {
        t.Error("Error message should appear at INFO level")
    }
}
```

### Step 4: Integrate with Formatter
File: `formatter/formatter.go` (update)
```go
package formatter

import (
    "yourmodule/logger"
    "yourmodule/tokenizer"
)

type Formatter struct {
    policy ErrorPolicy
    errors []*FormatterError
    logger logger.Logger
}

func New() *Formatter {
    return &Formatter{
        policy: PolicyContinue,
        errors: make([]*FormatterError, 0),
        logger: logger.New(),
    }
}

func (f *Formatter) SetLogger(l logger.Logger) {
    f.logger = l
}

func (f *Formatter) Format(input string) string {
    f.logger.Info("Starting format operation")
    f.logger.Debug("Input length: %d", len(input))
    
    tokens := tokenizer.Tokenize(input)
    f.logger.Debug("Tokenized into %d tokens", len(tokens))
    
    tokens = NormalizePunctuation(tokens)
    f.logger.Debug("Normalized punctuation")
    
    tokens = CleanQuoteSpacing(tokens)
    f.logger.Debug("Cleaned quote spacing")
    
    tokens = ProcessWithQuotes(tokens)
    f.logger.Debug("Processed quotes and tags")
    
    result := Render(tokens)
    f.logger.Info("Format operation complete")
    
    return result
}
```

### Step 5: Update FSM with Logging
File: `fsm/fsm.go` (update)
```go
func (f *FSM) step() {
    f.logger.Debug("FSM state: %s, position: %d", f.state, f.position)
    
    switch f.state {
    case StateReading:
        f.handleReading()
    case StateEvaluating:
        f.handleEvaluating()
    case StateEditing:
        f.handleEditing()
    case StateError:
        f.logger.Error("FSM in error state: %s", f.errorMsg)
        return
    }
}
```

## Verification Commands
```bash
go test ./logger/... -v
go test ./...
```

## Success Criteria
- Logger interface defined
- Log levels work correctly
- Integration with formatter complete
- Tests verify logging behavior

## TDD Workflow
1. ✅ RED: Write failing logger tests
2. ✅ GREEN: Implement logger interface and default
3. ✅ REFACTOR: Integrate with existing components

## Git Commit Message
```
feat: implement logging and diagnostics

- Define Logger interface with levels
- Implement DefaultLogger with configurable output
- Integrate logging with Formatter and FSM
- Add comprehensive logging tests
```
