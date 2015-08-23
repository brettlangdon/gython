package token

import "fmt"

type Token struct {
	ColumnEnd   int
	LineEnd     int
	ID          TokenID
	Literal     string
	ColumnStart int
	LineStart   int
}

func (token *Token) String() string {
	return TokenNames[token.ID]
}

func (token *Token) Start() []int {
	return []int{token.LineStart, token.ColumnStart}
}

func (token *Token) End() []int {
	return []int{token.LineEnd, token.ColumnEnd}
}

func (token *Token) Repr() string {
	return fmt.Sprintf(
		"Token{ID: %#v, Literal: %#v, LineStart: %#v, ColumnStart: %#v, LineEnd: %#v, ColumnEnd: %#v}",
		token.ID,
		token.Literal,
		token.LineStart,
		token.ColumnStart,
		token.LineEnd,
		token.ColumnEnd,
	)
}
