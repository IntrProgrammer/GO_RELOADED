package formatter

import (
	"GO_RELOADED/tokenizer"
	"reflect"
	"testing"
)

func TestNormalizePunctuation(t *testing.T) {
	tests := []struct {
		name   string
		tokens []tokenizer.Token
		want   []tokenizer.Token
	}{
		{
			name: "remove space before comma",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.PUNCTUATION, ","},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.PUNCTUATION, ","},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
			},
		},
		{
			name: "ensure space after punctuation",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.PUNCTUATION, ","},
				{tokenizer.WORD, "world"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.PUNCTUATION, ","},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizePunctuation(tt.tokens)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NormalizePunctuation() = %v, want %v", got, tt.want)
			}
		})
	}
}
