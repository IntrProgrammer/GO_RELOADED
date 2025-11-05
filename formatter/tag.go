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
