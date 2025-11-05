package formatter

import (
	"GO_RELOADED/tokenizer"
	"fmt"
	"strconv"
)

func convertHex(s string) (string, error) {
	val, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return "", fmt.Errorf("invalid hex: %w", err)
	}
	return strconv.FormatInt(val, 10), nil
}

func convertBin(s string) (string, error) {
	val, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return "", fmt.Errorf("invalid binary: %w", err)
	}
	return strconv.FormatInt(val, 10), nil
}

func ApplyNumberConversion(tokens []tokenizer.Token) ([]tokenizer.Token, error) {
	result := make([]tokenizer.Token, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type != tokenizer.TAG {
			result = append(result, tokens[i])
			continue
		}

		tag, err := ParseTag(tokens[i].Value)
		if err != nil {
			result = append(result, tokens[i])
			continue
		}

		if tag.Command != TagHex && tag.Command != TagBin {
			result = append(result, tokens[i])
			continue
		}

		// Find previous word token
		wordIdx := -1
		for j := len(result) - 1; j >= 0; j-- {
			if result[j].Type == tokenizer.WORD {
				wordIdx = j
				break
			}
		}

		if wordIdx == -1 {
			return nil, fmt.Errorf("no word before conversion tag")
		}

		// Convert
		var converted string
		if tag.Command == TagHex {
			converted, err = convertHex(result[wordIdx].Value)
		} else {
			converted, err = convertBin(result[wordIdx].Value)
		}

		if err != nil {
			return nil, err
		}

		result[wordIdx].Value = converted
		// Skip tag (don't append)
	}

	return result, nil
}
