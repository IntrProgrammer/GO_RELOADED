package formatter

import (
	"GO_RELOADED/tokenizer"
	"strings"
	"unicode"
)

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func toLower(s string) string {
	return strings.ToLower(s)
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

func ApplyCaseTransform(tokens []tokenizer.Token) []tokenizer.Token {
	result := make([]tokenizer.Token, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type != tokenizer.TAG {
			result = append(result, tokens[i])
			continue
		}

		tag, err := ParseTag(tokens[i].Value)
		if err != nil || (tag.Command != TagUp && tag.Command != TagLow && tag.Command != TagCap) {
			result = append(result, tokens[i])
			continue
		}

		// Find previous N words
		wordIndices := []int{}
		for j := len(result) - 1; j >= 0 && len(wordIndices) < tag.Count; j-- {
			if result[j].Type == tokenizer.WORD {
				wordIndices = append(wordIndices, j)
			}
		}

		if len(wordIndices) == 0 {
			continue
		}

		// Apply transformation to all found words
		for _, idx := range wordIndices {
			switch tag.Command {
			case TagUp:
				result[idx].Value = toUpper(result[idx].Value)
			case TagLow:
				result[idx].Value = toLower(result[idx].Value)
			case TagCap:
				result[idx].Value = capitalize(result[idx].Value)
			}
		}
	}

	return result
}
