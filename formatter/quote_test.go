package formatter

import (
	"GO_RELOADED/tokenizer"
	"reflect"
	"testing"
)

func TestCleanQuoteSpacing(t *testing.T) {
	tests := []struct {
		name   string
		tokens []tokenizer.Token
		want   []tokenizer.Token
	}{
		{
			name: "remove space after opening quote",
			tokens: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "hello"},
				{tokenizer.QUOTE, "'"},
			},
			want: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "hello"},
				{tokenizer.QUOTE, "'"},
			},
		},
		{
			name: "remove space before closing quote",
			tokens: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.QUOTE, "'"},
			},
			want: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "hello"},
				{tokenizer.QUOTE, "'"},
			},
		},
		{
			name: "preserve internal spacing",
			tokens: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
				{tokenizer.QUOTE, "'"},
			},
			want: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
				{tokenizer.QUOTE, "'"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CleanQuoteSpacing(tt.tokens)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CleanQuoteSpacing() = %v, want %v", got, tt.want)
			}
		})
	}
}
