// FSM — concise overview
// Purpose: orchestrates token processing: reads tokens, delegates token-specific work to processors,
// collects output, and reports errors.
// Flow: tokens -> Run() loop -> step() -> handleReading / handleEvaluating / handleEditing -> result

package fsm

import "GO_RELOADED/tokenizer"

type Processor interface {
	Process(result []tokenizer.Token, currentToken tokenizer.Token) (modified []tokenizer.Token, handled bool)
}

type FSM struct {
	state      State             // Current state of the FSM (Reading, Evaluating, Editing, Done, Error)
	tokens     []tokenizer.Token // Input tokens to process
	position   int               // Current position in the tokens slice
	result     []tokenizer.Token // Processed tokens output
	errorMsg   string            // Error message if FSM enters error state
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

func (f *FSM) Error() string {
	return f.errorMsg
}

func (f *FSM) Result() []tokenizer.Token {
	return f.result
}

func (f *FSM) Run() {
	for f.state != StateDone && f.state != StateError {
		f.step()
	}
	// Post-process: Apply article correction
	f.result = CorrectArticles(f.result)
}

func (f *FSM) step() {
	switch f.state {
	case StateReading:
		f.handleReading()
	case StateEvaluating:
		f.handleEvaluating()
	case StateEditing:
		f.handleEditing()
	case StateError:
		return
	}
}

func (f *FSM) handleReading() {
	if f.position >= len(f.tokens) {
		f.state = StateDone
		return
	}

	f.state = StateEvaluating
}

func (f *FSM) handleEvaluating() {
	token := f.tokens[f.position]

	// Track quote boundaries and handle spacing
	if token.Type == tokenizer.QUOTE {
		wasInQuote := f.inQuote
		f.inQuote = !f.inQuote
		
		// Use QuoteSpacingProcessor to handle quote spacing
		for _, proc := range f.processors {
			if modified, handled := proc.Process(f.result, token); handled {
				f.result = modified
				f.position++
				
				// Skip whitespace after opening quote
				if !wasInQuote && f.position < len(f.tokens) && f.tokens[f.position].Type == tokenizer.WHITESPACE {
					f.position++
				}
				
				f.state = StateReading
				return
			}
		}
		
		// Fallback if no processor handled it
		f.result = append(f.result, token)
		f.position++
		f.state = StateReading
		return
	}

	if token.Type == tokenizer.TAG {
		f.state = StateEditing
		return
	}

	// Try processors for non-tag tokens
	for _, proc := range f.processors {
		if modified, handled := proc.Process(f.result, token); handled {
			f.result = modified
			f.position++
			f.state = StateReading
			return
		}
	}

	// Default: append token
	f.result = append(f.result, token)
	f.position++
	f.state = StateReading
}

func (f *FSM) handleEditing() {
	tag := f.tokens[f.position]

	// Try processors for tokens
	for _, proc := range f.processors {
		if modified, handled := proc.Process(f.result, tag); handled {
			f.result = modified
			f.position++
			f.state = StateReading
			return
		}
	}

	// Default: skip
	f.position++
	f.state = StateReading
}
