package token

import "fmt"

type Token struct {
	End     int
	ID      TokenID
	Literal string
	Start   int
}

func (token *Token) String() string {
	return TokenNames[token.ID]
}

func (token *Token) Repr() string {
	return fmt.Sprintf(
		"Token{ID: %#v, Literal: %#v, Start: %#v, End: %#v}",
		token.ID,
		token.Literal,
		token.Start,
		token.End,
	)
}
