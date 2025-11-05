package formatter

import (
	"GO_RELOADED/tokenizer"
	"strings"
)

func Render(tokens []tokenizer.Token) string {
	var builder strings.Builder
	for _, token := range tokens {
		builder.WriteString(token.Value)
	}
	return builder.String()
}
