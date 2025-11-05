package formatter

import (
	"GO_RELOADED/tokenizer"
)

func CleanQuoteSpacing(tokens []tokenizer.Token) []tokenizer.Token {
	// Implementation:
	// 1. Track quote state (inside/outside)
	// 2. Remove whitespace immediately after opening quote
	// 3. Remove whitespace immediately before closing quote
	var result []tokenizer.Token
	insideQuote := false

	for i := 0; i < len(tokens); i++ {
		//Tells the programme when to trim spaces when inside Quote
		current := tokens[i]
		if current.Type == tokenizer.QUOTE && !insideQuote {
			insideQuote = true
			result = append(result, current)
			// Skip any following whitespace
			if i+1 < len(tokens) && tokens[i+1].Type == tokenizer.WHITESPACE {
				i++ // skip the whitespace
			}
			continue
		}
		//When you find a closing Quote and its len is higher than 1 (i>1)
		if insideQuote && current.Type == tokenizer.QUOTE {
			// Keep removing last token while it's whitespace
			for len(result) > 0 && result[len(result)-1].Type == tokenizer.WHITESPACE {
				result = result[:len(result)-1]
			}
			result = append(result, current)
			insideQuote = false
			continue
		}
		//Add other tokens to the result
		result = append(result, current)
	}
	return result
}
