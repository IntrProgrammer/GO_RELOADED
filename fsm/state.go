package fsm

type State int

const (
	StateReading State = iota
	StateEvaluating
	StateEditing
	StateError
	StateDone
)

func (s State) String() string {
	return [...]string{"READING", "EVALUATING", "EDITING", "ERROR", "DONE"}[s]
}
