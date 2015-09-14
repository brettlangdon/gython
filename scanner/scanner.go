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
		last := len(scanner.positionBuffer) - 1
		scanner.currentPosition = scanner.positionBuffer[last]
		scanner.positionBuffer = scanner.positionBuffer[0:last]
		return scanner.currentPosition
	}

	next, _, err := scanner.reader.ReadRune()
	if err != nil {
		scanner.state = errorcode.E_EOF
		next = EOF
		scanner.currentLine++
		scanner.currentColumn = 0
	}

	if next == '\n' || next == EOF {
		scanner.currentLine++
		scanner.currentColumn = 0
	}

	pos := &Position{
		Char:   next,
		Line:   scanner.currentLine,
		Column: scanner.currentColumn,
	}
	scanner.currentColumn++
	return pos
}

func (scanner *Scanner) unreadPosition(pos *Position) {
	scanner.positionBuffer = append(scanner.positionBuffer, pos)
}

func (scanner *Scanner) parseQuoted(positions *Positions, quote rune) *token.Token {
	// Determine quote size, 1 or 3 (e.g. 'string',  '''string''')
	quoteSize := 1
	endQuoteSize := 0
	pos := scanner.nextPosition()
	if pos.Char == quote {
		pos2 := scanner.nextPosition()
		if pos2.Char == quote {
			positions.Append(pos)
			positions.Append(pos2)
			quoteSize = 3
		} else {
			scanner.unreadPosition(pos2)
			endQuoteSize = 1
		}
	} else {
		scanner.unreadPosition(pos)
	}

	for {
		if endQuoteSize == quoteSize {
			break
		}
		pos = scanner.nextPosition()
		positions.Append(pos)
		if pos.Char == EOF {
			return positions.AsToken(token.ERRORTOKEN)
		}
		if quoteSize == 1 && pos.Char == '\n' {
			return positions.AsToken(token.ERRORTOKEN)
		}
		if pos.Char == quote {
			endQuoteSize += 1
		} else {
			endQuoteSize = 0
			if pos.Char == '\\' {
				pos = scanner.nextPosition()
			}
		}
	}
	return positions.AsToken(token.STRING)
}

func (scanner *Scanner) NextToken() *token.Token {
	positions := NewPositions()

	pos := scanner.nextPosition()
	// skip spaces
	for {
		if pos.Char != ' ' && pos.Char != '\t' {
			break
		}
		pos = scanner.nextPosition()
	}

	// skip comments
	if pos.Char == '#' {
		for {
			pos = scanner.nextPosition()
			if pos.Char == EOF || pos.Char == '\n' {
				break
			}
		}
	}

	positions.Append(pos)
	switch ch := pos.Char; {
	case ch == EOF:
		id := token.ENDMARKER
		if scanner.state != errorcode.E_EOF {
			id = token.ERRORTOKEN
		}
		return positions.AsToken(id)
	case IsIdentifierStart(ch):
		// Parse Identifier
		saw_b, saw_r, saw_u := false, false, false
		for {
			if !(saw_b || saw_u) && (ch == 'b' || ch == 'B') {
				saw_b = true
			} else if !(saw_b || saw_u || saw_r) && (ch == 'u' || ch == 'U') {
				saw_u = true
			} else if !(saw_r || saw_u) && (ch == 'r' || ch == 'R') {
				saw_r = true
			} else {
				break
			}
			pos = scanner.nextPosition()
			if IsQuote(pos.Char) {
				positions.Append(pos)
				return scanner.parseQuoted(positions, pos.Char)
			}
		}
		pos = scanner.nextPosition()
		for IsIdentifierChar(pos.Char) {
			positions.Append(pos)
			pos = scanner.nextPosition()
		}
		scanner.unreadPosition(pos)
		return positions.AsToken(token.NAME)
	case ch == '\n':
		return positions.AsToken(token.NEWLINE)
	case ch == '.':
		pos2 := scanner.nextPosition()
		if IsDigit(pos2.Char) {
			// Parse Number
		} else if pos2.Char == '.' {
			positions.Append(pos2)
			pos3 := scanner.nextPosition()
			if pos3.Char == '.' {
				positions.Append(pos3)
				return positions.AsToken(token.ELLIPSIS)
			}
			scanner.unreadPosition(pos3)
		}
		scanner.unreadPosition(pos2)

		return positions.AsToken(token.DOT)
	case IsDigit(ch):
		// Parse Number
	case IsQuote(ch):
		// Parse String
		return scanner.parseQuoted(positions, ch)
	case ch == '\\':
		// Parse Continuation
	default:
		// Two and Three character operators
		pos2 := scanner.nextPosition()
		op2Id := GetTwoCharTokenID(pos.Char, pos2.Char)
		if op2Id != token.OP {
			positions.Append(pos2)
			pos3 := scanner.nextPosition()
			op3Id := GetThreeCharTokenID(pos.Char, pos2.Char, pos3.Char)
			if op3Id != token.OP {
				positions.Append(pos3)
				return positions.AsToken(op3Id)
			}
			scanner.unreadPosition(pos3)
			return positions.AsToken(op2Id)
		}
		scanner.unreadPosition(pos2)
	}
	switch pos.Char {
	case '(', '[', '{':
		// Increment indentation level
		// scanner.indentationLevel++
		break
	case ')', ']', '}':
		// Decrement indentation level
		// scanner.indentationLevel--
		break
	}

	opId := GetOneCharTokenID(pos.Char)
	return positions.AsToken(opId)
}
