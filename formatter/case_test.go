package formatter

import (
	"GO_RELOADED/tokenizer"
	"reflect"
	"testing"
)

func tokensEqual(a, b []tokenizer.Token) bool {
	return reflect.DeepEqual(a, b)
}

func TestApplyCaseTransform(t *testing.T) {
	tests := []struct {
		name   string
		tokens []tokenizer.Token
		want   []tokenizer.Token
	}{
		{
			name: "uppercase single word",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(up)"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "HELLO"},
				{tokenizer.WHITESPACE, " "},
			},
		},
		{
			name: "lowercase single word",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "HELLO"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(low)"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
			},
		},
		{
			name: "capitalize single word",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(cap)"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "Hello"},
				{tokenizer.WHITESPACE, " "},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ApplyCaseTransform(tt.tokens)
			if !tokensEqual(got, tt.want) {
				t.Errorf("ApplyCaseTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplyCaseTransformMulti(t *testing.T) {
	tests := []struct {
		name   string
		tokens []tokenizer.Token
		want   []tokenizer.Token
	}{
		{
			name: "uppercase multiple words",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(up, 2)"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "HELLO"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "WORLD"},
				{tokenizer.WHITESPACE, " "},
			},
		},
		{
			name: "capitalize with count larger than words",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(cap, 5)"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "Hello"},
				{tokenizer.WHITESPACE, " "},
			},
		},
		{
			name: "lowercase three words",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "HELLO"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "BEAUTIFUL"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "WORLD"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(low, 3)"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "beautiful"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
				{tokenizer.WHITESPACE, " "},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ApplyCaseTransform(tt.tokens)
			if !tokensEqual(got, tt.want) {
				t.Errorf("ApplyCaseTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
