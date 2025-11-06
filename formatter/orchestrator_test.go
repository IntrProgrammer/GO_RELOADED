package formatter

import (
	"GO_RELOADED/tokenizer"
	"testing"
)

func TestApplyAllTags(t *testing.T) {
	tests := []struct {
		name   string
		tokens []tokenizer.Token
		want   []tokenizer.Token
	}{
		{
			name: "mixed hex and case",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "1a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(hex)"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(up)"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "26"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "HELLO"},
				{tokenizer.WHITESPACE, " "},
			},
		},
		{
			name: "multiple case tags",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(up)"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "test"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(cap)"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "WORLD"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "Test"},
				{tokenizer.WHITESPACE, " "},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ApplyAllTags(tt.tokens)
			if !tokensEqual(got, tt.want) {
				t.Errorf("ApplyAllTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
