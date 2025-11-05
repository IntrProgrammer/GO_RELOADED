package formatter

import (
	"testing"
)

func TestParseTag(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Tag
		wantErr bool
	}{
		{
			name:  "simple up tag",
			input: "(up)",
			want:  Tag{Command: TagUp, Count: 1},
		},
		{
			name:  "cap with count",
			input: "(cap, 3)",
			want:  Tag{Command: TagCap, Count: 3},
		},
		{
			name:  "hex tag",
			input: "(hex)",
			want:  Tag{Command: TagHex, Count: 1},
		},
		{
			name:    "invalid format",
			input:   "up",
			wantErr: true,
		},
		{
			name:    "unknown command",
			input:   "(unknown)",
			wantErr: true,
		},
		{
			name:    "invalid count",
			input:   "(up, abc)",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTag(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ParseTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
