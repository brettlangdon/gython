package scanner

import (
	"github.com/brettlangdon/gython/token"
)

type Position struct {
	Line   int
	Column int
	Char   rune
}

type Positions struct {
	positions []*Position
}

func NewPositions() *Positions {
	return &Positions{
		positions: make([]*Position, 0),
	}
}

func (positions *Positions) Append(pos *Position) {
	positions.positions = append(positions.positions, pos)
}

func (positions *Positions) StartingLine() int {
	return positions.positions[0].Line
}

func (positions *Positions) EndingLine() int {
	last := len(positions.positions) - 1
	return positions.positions[last].Line
}

func (positions *Positions) StartingColumn() int {
	return positions.positions[0].Column
}

func (positions *Positions) EndingColumn() int {
	last := len(positions.positions) - 1
	return positions.positions[last].Column
}

func (positions *Positions) String() string {
	literal := ""
	for _, pos := range positions.positions {
		literal += string(pos.Char)
	}

	return literal
}

func (positions *Positions) AsToken(id token.TokenID) *token.Token {
	return &token.Token{
		ID:          id,
		LineStart:   positions.StartingLine(),
		ColumnStart: positions.StartingColumn(),
		LineEnd:     positions.EndingLine(),
		ColumnEnd:   positions.EndingColumn(),
		Literal:     positions.String(),
	}
}
