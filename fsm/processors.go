package fsm

import (
	"GO_RELOADED/tokenizer"
	"strconv"
	"strings"
	"unicode"
)

// ConversionProcessor handles (hex) and (bin) tags
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

// CaseProcessor handles (up), (low), (cap) tags
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

// Helper functions

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

func startsWithVowelOrH(s string) bool {
	if len(s) == 0 {
		return false
	}
	first := unicode.ToLower(rune(s[0]))
	if first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u' || first == 'h' {
		return true
	}
	if first == '8' || (len(s) >= 2 && s[0] == '1' && (s[1] == '1' || s[1] == '8')) {
		return true
	}
	return false
}

// CorrectArticles is a post-processor for a/an correction
func CorrectArticles(tokens []tokenizer.Token) []tokenizer.Token {
	result := make([]tokenizer.Token, len(tokens))
	copy(result, tokens)

	for i := 0; i < len(result); i++ {
		if result[i].Type != tokenizer.WORD {
			continue
		}

		word := result[i].Value
		if strings.ToLower(word) != "a" {
			continue
		}

		nextWordIdx := -1
		for j := i + 1; j < len(result); j++ {
			if result[j].Type == tokenizer.WORD {
				nextWordIdx = j
				break
			}
		}

		if nextWordIdx == -1 {
			continue
		}

		if startsWithVowelOrH(result[nextWordIdx].Value) {
			if word == "a" {
				result[i].Value = "an"
			} else if word == "A" {
				result[i].Value = "An"
			}
		}
	}

	return result
}
