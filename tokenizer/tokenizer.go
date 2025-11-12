package tokenizer

import (
	"regexp"
)

const (
	// Group 1: TAG - Matches commands like (up), (hex, 2)
	TagPattern = `(\([^)]+\))`
	// Group 2: QUOTE - Matches single or double quote characters
	QuotePattern = `(['"])`
	// Group 3: PUNCTUATION - Matches any defined punctuation mark
	PunctuationPattern = `([.,!?;:])`
	// Group 4: WHITESPACE - Matches one or more spaces
	WhitespacePattern = `(\s+)`
	// Group 5: WORD - The catch-all: matches one or more non-whitespace characters
	WordPattern = `([^\s.,!?;:()'"]+)`
)

// The master regex combines all patterns using OR (|).
// The order is important: specific patterns first, generic patterns last.
var masterRegex = regexp.MustCompile(
	TagPattern + `|` +
		QuotePattern + `|` +
		PunctuationPattern + `|` +
		WhitespacePattern + `|` +
		WordPattern,
)

func determineTokenType(match []string) TokenType {
	// The groups from masterRegex are numbered starting from 1.

	// match[1] corresponds to the 1st the TagPattern
	if match[1] != "" {
		return TAG
	}
	// match[2] corresponds to the 2nd parenthesis: QuotePattern
	if match[2] != "" {
		return QUOTE
	}
	// match[3] corresponds to the 3rd parenthesis: PunctuationPattern
	if match[3] != "" {
		return PUNCTUATION
	}
	// match[4] corresponds to the 4th parenthesis: WhitespacePattern
	if match[4] != "" {
		return WHITESPACE
	}
	// If all specific groups are empty, it must have matched the 5th group (WordPattern).
	return WORD
}

func Tokenize(input string) []Token {
	var tokens []Token
	//Grouping all the matches together in order to order them to the apropriete tokenTypes
	matches := masterRegex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		tokenType := determineTokenType(match)
		tokens = append(tokens, Token{tokenType, match[0]})
	}
	return tokens
}
