package tokenizer

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Token
	}{
		{
			name:  "simple words",
			input: "hello world",
			want: []Token{
				{WORD, "hello"},
				{WHITESPACE, " "},
				{WORD, "world"},
			},
		},
		{
			name:  "word with punctuation",
			input: "hello, world!",
			want: []Token{
				{WORD, "hello"},
				{PUNCTUATION, ","},
				{WHITESPACE, " "},
				{WORD, "world"},
				{PUNCTUATION, "!"},
			},
		},
		{
			name:  "with tag",
			input: "word (up)",
			want: []Token{
				{WORD, "word"},
				{WHITESPACE, " "},
				{TAG, "(up)"},
			},
		},
		{
			name:  "with quotes",
			input: "'hello world'",
			want: []Token{
				{QUOTE, "'"},
				{WORD, "hello"},
				{WHITESPACE, " "},
				{WORD, "world"},
				{QUOTE, "'"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Tokenize(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
