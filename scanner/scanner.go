package scanner

import (
	"bufio"
	"io"

	"github.com/brettlangdon/gython/errorcode"
	"github.com/brettlangdon/gython/token"
)

var EOF rune = 0
var MAXINDENT int = 100

type Scanner struct {
	state           errorcode.ErrorCode
	reader          *bufio.Reader
	currentPosition *Position
	positionBuffer  []*Position

	currentLine   int
	currentColumn int
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		state:          errorcode.E_OK,
		reader:         bufio.NewReader(r),
		positionBuffer: make([]*Position, 0),
		currentLine:    1,
		currentColumn:  0,
	}
}

func (scanner *Scanner) nextPosition() *Position {
	if len(scanner.positionBuffer) > 0 {
		last = len(scanner.positionBuffer) - 1
		scanner.currentPosition = scanner.positionBuffer[last]
		scanner.positionBuffer = scanner.positionBuffer[0:last]
		return scanner.currentPosition
	}

	next, _, err := scanner.reader.ReadRune()
	if err != nil {
		scanner.state = errorcode.E_EOF
		next = EOF
	}

	scanner.currentColumn++
	if next == '\n' {
		scanner.currentLine++
		scanner.currentColumn = 0
	}

	return &Position{
		Char:   next,
		Line:   scanner.currentLine,
		Column: scanner.currentColumn,
	}
}

func (scanner *Scanner) unreadPosition(pos *Position) {
	scanner.positionBuffer = append(scanner.positionBuffer, pos)
}

func (scanner *Scanner) NextToken() *token.Token {
	return &Token{
		ID: token.ENDMARKER,
	}
}
