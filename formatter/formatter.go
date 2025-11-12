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
	tokens := tokenizer.Tokenize(input)//Pipe
	tokens = NormalizePunctuation(tokens)//Pipe
	tokens = CleanQuoteSpacing(tokens)//Pipe
	tokens = processWithFSM(tokens)//FSM
	return Render(tokens)
}

func processWithFSM(tokens []tokenizer.Token) []tokenizer.Token {
	segments := splitByQuotes(tokens)
	result := []tokenizer.Token{}

	for _, segment := range segments {
		machine := fsm.New(segment)
		machine.AddProcessor(&fsm.ConversionProcessor{})
		machine.AddProcessor(&fsm.CaseProcessor{})
		machine.Run()
		processed := machine.Result()
		processed = fsm.CorrectArticles(processed)
		result = append(result, processed...)
	}

	return result
}

func splitByQuotes(tokens []tokenizer.Token) [][]tokenizer.Token {
	segments := [][]tokenizer.Token{}
	current := []tokenizer.Token{}
	inQuote := false

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type == tokenizer.QUOTE {
			if !inQuote {
				if len(current) > 0 {
					segments = append(segments, current)
					current = []tokenizer.Token{}
				}
				current = append(current, tokens[i])
				inQuote = true
			} else {
				current = append(current, tokens[i])
				segments = append(segments, current)
				current = []tokenizer.Token{}
				inQuote = false
			}
		} else {
			current = append(current, tokens[i])
		}
	}

	if len(current) > 0 {
		segments = append(segments, current)
	}

	return segments
}
