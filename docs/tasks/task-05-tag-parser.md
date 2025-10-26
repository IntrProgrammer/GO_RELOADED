# Task 5 — Tag Parser

## Objective
Extract structured information from tag tokens like `(up)`, `(hex)`, `(cap, 3)`.

## Prerequisites
- Task 1 completed (Tokenizer)

## Deliverables
- [ ] Tag structure definition
- [ ] ParseTag function
- [ ] Tests for valid and invalid tags
- [ ] Error handling

## Tag Format
- Command only: `(up)`, `(low)`, `(cap)`, `(hex)`, `(bin)`
- Command with count: `(up, 2)`, `(cap, 3)`

## Implementation Steps

### Step 1: Define Tag Structure
File: `formatter/tag.go`
```go
package formatter

import (
    "errors"
    "strconv"
    "strings"
)

type TagCommand string

const (
    TagUp  TagCommand = "up"
    TagLow TagCommand = "low"
    TagCap TagCommand = "cap"
    TagHex TagCommand = "hex"
    TagBin TagCommand = "bin"
)

type Tag struct {
    Command TagCommand
    Count   int // 1 if not specified
}

var (
    ErrInvalidTag     = errors.New("invalid tag format")
    ErrUnknownCommand = errors.New("unknown tag command")
)
```

### Step 2: Write Tag Parser Tests
File: `formatter/tag_test.go`
```go
package formatter

import (
    "testing"
)

func TestParseTag(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    Tag
        wantErr bool
    }{
        {
            name:  "simple up tag",
            input: "(up)",
            want:  Tag{Command: TagUp, Count: 1},
        },
        {
            name:  "cap with count",
            input: "(cap, 3)",
            want:  Tag{Command: TagCap, Count: 3},
        },
        {
            name:  "hex tag",
            input: "(hex)",
            want:  Tag{Command: TagHex, Count: 1},
        },
        {
            name:    "invalid format",
            input:   "up",
            wantErr: true,
        },
        {
            name:    "unknown command",
            input:   "(unknown)",
            wantErr: true,
        },
        {
            name:    "invalid count",
            input:   "(up, abc)",
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := ParseTag(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("ParseTag() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr && got != tt.want {
                t.Errorf("ParseTag() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Step 3: Implement Tag Parser
File: `formatter/tag.go` (continued)
```go
func ParseTag(input string) (Tag, error) {
    // Remove parentheses
    if !strings.HasPrefix(input, "(") || !strings.HasSuffix(input, ")") {
        return Tag{}, ErrInvalidTag
    }
    
    content := strings.TrimPrefix(input, "(")
    content = strings.TrimSuffix(content, ")")
    
    // Split by comma
    parts := strings.Split(content, ",")
    if len(parts) > 2 {
        return Tag{}, ErrInvalidTag
    }
    
    command := strings.TrimSpace(parts[0])
    count := 1
    
    if len(parts) == 2 {
        var err error
        count, err = strconv.Atoi(strings.TrimSpace(parts[1]))
        if err != nil || count < 1 {
            return Tag{}, ErrInvalidTag
        }
    }
    
    // Validate command
    tagCmd := TagCommand(command)
    switch tagCmd {
    case TagUp, TagLow, TagCap, TagHex, TagBin:
        return Tag{Command: tagCmd, Count: count}, nil
    default:
        return Tag{}, ErrUnknownCommand
    }
}
```

## Verification Commands
```bash
go test ./formatter/... -v -run TestParseTag
go test ./...
```

## Success Criteria
- Valid tags parse correctly
- Invalid tags return errors
- Count defaults to 1
- No panics on malformed input

## TDD Workflow
1. ✅ RED: Write failing tests for all tag formats
2. ✅ GREEN: Implement ParseTag()
3. ✅ REFACTOR: Clean up parsing logic

## Git Commit Message
```
feat: implement tag parser

- Define Tag structure and TagCommand enum
- Implement ParseTag with validation
- Add comprehensive error handling
- Support command-only and command-with-count formats
```
