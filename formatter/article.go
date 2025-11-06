package formatter

import (
	"GO_RELOADED/tokenizer"
	"strings"
	"unicode"
)

func startsWithVowelOrH(s string) bool {
	if len(s) == 0 {
		return false
	}
	first := unicode.ToLower(rune(s[0]))
	return first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u' || first == 'h'
}

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
