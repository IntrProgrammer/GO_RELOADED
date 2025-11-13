package formatter

import (
	"GO_RELOADED/tokenizer"
	"testing"
)

func TestRender(t *testing.T) {
	tests := []struct {
		name   string
		tokens []tokenizer.Token
		want   string
	}{
		{
			name: "simple words",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
			},
			want: "hello world",
		},
		{
			name: "with punctuation",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.PUNCTUATION, ","},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
				{tokenizer.PUNCTUATION, "!"},
			},
			want: "hello, world!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Render(tt.tokens)
			if got != tt.want {
				t.Errorf("Render() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "normalize punctuation spacing",
			input: "hello , world !",
			want:  "hello, world!",
		},
		{
			name:  "clean quote spacing",
			input: "' hello world '",
			want:  "'hello world'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenizer.Tokenize(tt.input)
			got := Render(tokens)
			if got != tt.want {
				t.Errorf("Round-trip = %q, want %q", got, tt.want)
			}
		})
	}
}
