package fsm

import (
	"GO_RELOADED/tokenizer"
	"testing"
)

// PunctuationProcessor handles punctuation normalization
type PunctuationProcessor struct{}

func (p *PunctuationProcessor) Process(result []tokenizer.Token, current tokenizer.Token) ([]tokenizer.Token, bool) {
	return result, false // FSM doesn't use processors for this - handled in formatter
}

// QuoteProcessor handles quote spacing cleanup
type QuoteProcessor struct{}

func (q *QuoteProcessor) Process(result []tokenizer.Token, current tokenizer.Token) ([]tokenizer.Token, bool) {
	return result, false // FSM doesn't use processors for this - handled in formatter
}

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

func TestFSMPunctuationNormalization(t *testing.T) {
	// Note: FSM passes tokens through unchanged without processors
	// Actual punctuation normalization happens in formatter.NormalizePunctuation()
	// These tests verify FSM correctly processes punctuation tokens
	tests := []struct {
		name   string
		tokens []tokenizer.Token
	}{
		{
			name: "punctuation tokens pass through FSM",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "hello"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.PUNCTUATION, "!"},
			},
		},
		{
			name: "multiple punctuation marks",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "wow"},
				{tokenizer.PUNCTUATION, "!"},
				{tokenizer.PUNCTUATION, "!"},
				{tokenizer.PUNCTUATION, "!"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsm := New(tt.tokens)
			fsm.Run()
			result := fsm.Result()
			// FSM should pass all tokens through
			if len(result) != len(tt.tokens) {
				t.Errorf("Result length = %d, want %d", len(result), len(tt.tokens))
			}
			if fsm.CurrentState() != StateDone {
				t.Errorf("Final state = %v, want %v", fsm.CurrentState(), StateDone)
			}
		})
	}
}

func TestFSMQuoteSpacingCleanup(t *testing.T) {
	// Note: FSM passes tokens through unchanged without processors
	// Actual quote spacing cleanup happens in formatter.CleanQuoteSpacing()
	// These tests verify FSM correctly processes quote tokens
	tests := []struct {
		name   string
		tokens []tokenizer.Token
	}{
		{
			name: "quote tokens pass through FSM",
			tokens: []tokenizer.Token{
				{tokenizer.QUOTE, "\""},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "hello"},
				{tokenizer.QUOTE, "\""},
			},
		},
		{
			name: "single quotes with content",
			tokens: []tokenizer.Token{
				{tokenizer.QUOTE, "'"},
				{tokenizer.WORD, "test"},
				{tokenizer.QUOTE, "'"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsm := New(tt.tokens)
			fsm.Run()
			result := fsm.Result()
			// FSM should pass all tokens through
			if len(result) != len(tt.tokens) {
				t.Errorf("Result length = %d, want %d", len(result), len(tt.tokens))
			}
			if fsm.CurrentState() != StateDone {
				t.Errorf("Final state = %v, want %v", fsm.CurrentState(), StateDone)
			}
		})
	}
}
