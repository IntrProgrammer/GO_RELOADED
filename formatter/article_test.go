package formatter

import (
	"GO_RELOADED/tokenizer"
	"testing"
)

func TestCorrectArticles(t *testing.T) {
	tests := []struct {
		name   string
		tokens []tokenizer.Token
		want   []tokenizer.Token
	}{
		{
			name: "a before vowel",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "apple"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "an"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "apple"},
			},
		},
		{
			name: "a before consonant unchanged",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "book"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "book"},
			},
		},
		{
			name: "capital A before vowel",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "A"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "elephant"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "An"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "elephant"},
			},
		},
		{
			name: "a before h",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "hour"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "an"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "hour"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CorrectArticles(tt.tokens)
			if !tokensEqual(got, tt.want) {
				t.Errorf("CorrectArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}
