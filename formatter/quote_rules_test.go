package formatter

import (
	"GO_RELOADED/tokenizer"
	"testing"
)

func TestProcessWithQuotes(t *testing.T) {
	tests := []struct {
		name   string
		tokens []tokenizer.Token
		want   []tokenizer.Token
	}{
		{
			name: "case transform inside quotes",
			tokens: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(up)"},
				{tokenizer.QUOTE, "'"},
			},
			want: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "HELLO"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.QUOTE, "'"},
			},
		},
		{
			name: "article correction inside quotes",
			tokens: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "apple"},
				{tokenizer.QUOTE, "'"},
			},
			want: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "an"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "apple"},
				{tokenizer.QUOTE, "'"},
			},
		},
		{
			name: "hex conversion inside quotes",
			tokens: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "FF"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(hex)"},
				{tokenizer.QUOTE, "'"},
			},
			want: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "255"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.QUOTE, "'"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProcessWithQuotes(tt.tokens)
			if !tokensEqual(got, tt.want) {
				t.Errorf("ProcessWithQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
