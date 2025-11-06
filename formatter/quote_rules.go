package formatter

import "GO_RELOADED/tokenizer"

func ProcessWithQuotes(tokens []tokenizer.Token) []tokenizer.Token {
	result := make([]tokenizer.Token, 0, len(tokens))
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

	for _, segment := range segments {
		processed := ApplyAllTags(segment)
		processed = CorrectArticles(processed)
		result = append(result, processed...)
	}

	return result
}
