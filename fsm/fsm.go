// FSM — concise overview
// Purpose: orchestrates token processing: reads tokens, delegates token-specific work to processors,
// collects output, and reports errors.
// Flow: tokens -> Run() loop -> step() -> handleReading / handleEvaluating / handleEditing -> result

package fsm

import (
	"GO_RELOADED/tokenizer"
)

type Processor interface {
	Process(result []tokenizer.Token, currentToken tokenizer.Token) (modified []tokenizer.Token, handled bool)
}

type FSM struct {
	state      State             // Current state of the FSM (Reading, Evaluating, Editing, Done, Error)
	tokens     []tokenizer.Token // Input tokens to process
	position   int               // Current position in the tokens slice
	result     []tokenizer.Token // Processed tokens output
	processors []Processor       // List of processors to handle different token types
	inQuote    bool              // Track if currently inside quotes
}

func New(tokens []tokenizer.Token) *FSM {
	return &FSM{
		state:      StateReading,
		tokens:     tokens,
		position:   0,
		result:     make([]tokenizer.Token, 0, len(tokens)),
		processors: []Processor{},
		inQuote:    false,
	}
}

func (f *FSM) AddProcessor(p Processor) {
	f.processors = append(f.processors, p)
}

func (f *FSM) CurrentState() State {
	return f.state
}

func (f *FSM) Result() []tokenizer.Token {
	return f.result
}

func (f *FSM) Run() {
	for f.state != StateDone {
		f.step()
	}
	// Post-process: Apply article correction
	f.result = CorrectArticles(f.result)
}

func (f *FSM) step() {
	switch f.state {
	case StateReading:
		f.handleReading()
	case StateDone:
		return
	}
}

func (f *FSM) handleReading() {
	if f.position >= len(f.tokens) {
		f.state = StateDone
		return
	}
}
