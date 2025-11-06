package formatter

import "GO_RELOADED/tokenizer"

func ApplyAllTags(tokens []tokenizer.Token) []tokenizer.Token {
	result := tokens

	// Pass 1: Number conversions
	result, _ = ApplyNumberConversion(result)

	// Pass 2: Case transformations
	result = ApplyCaseTransform(result)

	return result
}
