package fsm

import (
	"GO_RELOADED/tokenizer"
	"testing"
)

func TestFSMStateTransitions(t *testing.T) {
	tests := []struct {
		name           string
		tokens         []tokenizer.Token
		wantFinalState State
	}{
		{
			name: "simple word processing",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
			},
			wantFinalState: StateDone,
		},
		{
			name: "tag processing",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.TAG, "(up)"},
			},
			wantFinalState: StateDone,
		},
		{
			name: "multiple tokens",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "world"},
			},
			wantFinalState: StateDone,
		},
		{
			name:           "empty input",
			tokens:         []tokenizer.Token{},
			wantFinalState: StateDone,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsm := New(tt.tokens)
			fsm.Run()
			if fsm.CurrentState() != tt.wantFinalState {
				t.Errorf("Final state = %v, want %v", fsm.CurrentState(), tt.wantFinalState)
			}
		})
	}
}

func TestFSMResult(t *testing.T) {
	tokens := []tokenizer.Token{
		{tokenizer.WORD, "hello"},
		{tokenizer.WHITESPACE, " "},
		{tokenizer.WORD, "world"},
	}

	fsm := New(tokens)
	fsm.Run()

	result := fsm.Result()
	if len(result) != 3 {
		t.Errorf("Result length = %d, want 3", len(result))
	}
}
