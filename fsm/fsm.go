package fsm

import "GO_RELOADED/tokenizer"

type FSM struct {
	state    State
	tokens   []tokenizer.Token
	position int
	result   []tokenizer.Token
	errorMsg string
}

func New(tokens []tokenizer.Token) *FSM {
	return &FSM{
		state:    StateReading,
		tokens:   tokens,
		position: 0,
		result:   make([]tokenizer.Token, 0, len(tokens)),
	}
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

	if token.Type == tokenizer.TAG {
		f.state = StateEditing
	} else {
		f.result = append(f.result, token)
		f.position++
		f.state = StateReading
	}
}

func (f *FSM) handleEditing() {
	f.position++
	f.state = StateReading
}
