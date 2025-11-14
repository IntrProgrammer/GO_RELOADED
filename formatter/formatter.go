package formatter

import (
	"GO_RELOADED/fsm"
	"GO_RELOADED/tokenizer"
)

type Formatter struct{}

func New() *Formatter {
	return &Formatter{}
}

func (f *Formatter) Format(input string) string {
	tokens := tokenizer.Tokenize(input)
	machine := fsm.New(tokens)
	machine.AddProcessor(&fsm.PunctuationNormalization{})
	machine.AddProcessor(&fsm.QuoteSpacingProcessor{})
	machine.AddProcessor(&fsm.ConversionProcessor{})
	machine.AddProcessor(&fsm.CaseProcessor{})
	machine.Run()
	return Render(machine.Result())
}
