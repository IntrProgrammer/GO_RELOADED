package fsm

type State int

const (
	StateReading State = iota
	StateEditing
	StateDone
)

func (s State) String() string {
	return [...]string{"READING", "EDITING", "DONE"}[s]
}
