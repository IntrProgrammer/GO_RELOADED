package formatter

import (
	"GO_RELOADED/tokenizer"
	"reflect"
	"testing"
)

func TestApplyNumberConversion(t *testing.T) {
	tests := []struct {
		name    string
		tokens  []tokenizer.Token
		want    []tokenizer.Token
		wantErr bool
	}{
		{
			name: "hex conversion",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "1E"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(hex)"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "30"},
				{tokenizer.WHITESPACE, " "},
			},
		},
		{
			name: "bin conversion",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "101010"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(bin)"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "42"},
				{tokenizer.WHITESPACE, " "},
			},
		},
		{
			name: "invalid hex",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "XYZ"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(hex)"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ApplyNumberConversion(tt.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyNumberConversion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplyNumberConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
