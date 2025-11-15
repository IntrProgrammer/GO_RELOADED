package formatter

import (
	"GO_RELOADED/tokenizer"
	"strings"
)

func Render(tokens []tokenizer.Token) string {
	var builder strings.Builder
	for i, token := range tokens {
		builder.WriteString(token.Value)
		if token.Type == tokenizer.PUNCTUATION && i < len(tokens)-1 && tokens[i+1].Type != tokenizer.WHITESPACE {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}
