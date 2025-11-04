package tokenizer

type TokenType int

const (
	WORD TokenType = iota
	PUNCTUATION
	TAG
	QUOTE
	WHITESPACE
)

type Token struct {
	Type  TokenType
	Value string
}
