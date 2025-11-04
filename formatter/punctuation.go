/*
	The following part is for handling paunctioations when occured in the programme
	For example:
		Hello,world ! -> Hello, world!
*/

package formatter

import "GO_RELOADED/tokenizer"

func NormalizePunctuation(tokens []tokenizer.Token) []tokenizer.Token {
	var result []tokenizer.Token
	for i := 0; i <= len(tokens)-1; i++ {
		//Remove whitespace before punctuation
		if tokens[i].Type == tokenizer.WHITESPACE &&
			i+1 < len(tokens) &&
			tokens[i+1].Type == tokenizer.PUNCTUATION {
			continue
		}
		result = append(result, tokens[i])
		// Ensure whitespace after punctuation (unless at the end or already whitespace)
		if tokens[i].Type == tokenizer.PUNCTUATION &&
			(i+1 == len(tokens) || tokens[i+1].Type != tokenizer.WHITESPACE) {
			if i+1 != len(tokens) {
				result = append(result, tokenizer.Token{Type: tokenizer.WHITESPACE, Value: " "})
			}
		}
	}
	return result
}
