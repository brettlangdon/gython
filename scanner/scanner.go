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
	asyncDef            bool
	atBol               bool
	currentColumn       int
	currentLine         int
	currentPosition     *Position
	indentationAltStack []int
	indentationCurrent  int
	indentationLevel    int
	indentationPending  int
	indentationStack    []int
	positionBuffer      []*Position
	tokenBuffer         []*token.Token
	reader              *bufio.Reader
	state               errorcode.ErrorCode
	tabsize             int
	tabsizeAlt          int
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		atBol:               true,
		currentColumn:       0,
		currentLine:         1,
		indentationAltStack: make([]int, MAXINDENT),
		indentationCurrent:  0,
		indentationLevel:    0,
		indentationPending:  0,
		indentationStack:    make([]int, MAXINDENT),
		positionBuffer:      make([]*Position, 0),
		tokenBuffer:         make([]*token.Token, 0),
		reader:              bufio.NewReader(r),
		state:               errorcode.E_OK,
		tabsize:             8,
	}
}

func (scanner *Scanner) State() errorcode.ErrorCode {
	return scanner.state
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

func (scanner *Scanner) parseNumber(positions *Positions, nextChar rune) *token.Token {
	pos := scanner.nextPosition()
	switch ch := pos.Char; {
	case nextChar == '0' && (ch == 'j' || ch == 'J'):
		// Imaginary
		positions.Append(pos)
	case nextChar == '0' && (ch == 'x' || ch == 'X'):
		// Hex
		positions.Append(pos)
		pos = scanner.nextPosition()
		if !IsXDigit(pos.Char) {
			return positions.AsToken(token.ERRORTOKEN)
		}
		for IsXDigit(pos.Char) {
			positions.Append(pos)
			pos = scanner.nextPosition()
		}
		scanner.unreadPosition(pos)
	case nextChar == '0' && (ch == 'b' || ch == 'B'):
		// Binary
		positions.Append(pos)
		pos = scanner.nextPosition()
		if pos.Char != '0' && pos.Char != '1' {
			return positions.AsToken(token.ERRORTOKEN)
		}
		for pos.Char == '0' || pos.Char == '1' {
			positions.Append(pos)
			pos = scanner.nextPosition()
		}
		scanner.unreadPosition(pos)
	case nextChar == '0' && (ch == 'o' || ch == 'O'):
		// Octal
		positions.Append(pos)
		pos = scanner.nextPosition()
		if pos.Char < '0' || pos.Char >= '8' {
			return positions.AsToken(token.ERRORTOKEN)
		}
		for pos.Char >= '0' && pos.Char < '8' {
			positions.Append(pos)
			pos = scanner.nextPosition()
		}
		scanner.unreadPosition(pos)
	default:
		decimal := nextChar == '.'
		imaginary := false
		exponent := false
		for {
			if pos.Char == '.' && decimal {
				break
			} else if pos.Char == '.' && !decimal {
				decimal = true
			} else if (pos.Char == 'j' || pos.Char == 'J') && !imaginary {
				imaginary = true
			} else if (pos.Char == 'e' || pos.Char == 'E') && !exponent {
				exponent = true
				positions.Append(pos)
				pos2 := scanner.nextPosition()
				if pos2.Char == '-' || pos2.Char == '+' {
					pos3 := scanner.nextPosition()
					if !IsDigit(pos3.Char) {
						return positions.AsToken(token.ERRORTOKEN)
					}
					scanner.unreadPosition(pos3)
					positions.Append(pos2)
				} else if !IsDigit(pos2.Char) {
					return positions.AsToken(token.ERRORTOKEN)
				} else {
					scanner.unreadPosition(pos2)
				}
				pos = scanner.nextPosition()
				continue
			} else if !IsDigit(pos.Char) {
				break
			}
			positions.Append(pos)
			pos = scanner.nextPosition()
		}
		scanner.unreadPosition(pos)
	}

	return positions.AsToken(token.NUMBER)
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

func (scanner *Scanner) unreadToken(tok *token.Token) {
	scanner.tokenBuffer = append(scanner.tokenBuffer, tok)
}

func (scanner *Scanner) NextToken() *token.Token {
next_line:
	if len(scanner.tokenBuffer) > 0 {
		last := len(scanner.tokenBuffer) - 1
		nextToken := scanner.tokenBuffer[last]
		scanner.tokenBuffer = scanner.tokenBuffer[0:last]
		return nextToken
	}

	blankline := false
	positions := NewPositions()
	var pos *Position

	if scanner.atBol {
		// Get indentation level
		col := 0
		altcol := 0
		scanner.atBol = false
		pos = scanner.nextPosition()
		for {
			if pos.Char == ' ' {
				col++
				altcol++
			} else if pos.Char == '\t' {
				col = (col/scanner.tabsize + 1) * scanner.tabsize
				altcol = (altcol/scanner.tabsizeAlt + 1) * scanner.tabsizeAlt
			} else {
				break
			}
			pos = scanner.nextPosition()
		}
		scanner.unreadPosition(pos)

		if pos.Char == '#' || pos.Char == '\n' {
			// Lines with only newline or comment, shouldn't affect indentation
			// TODO: Handle prompt
			if col == 0 && pos.Char == '\n' && false {
				blankline = false
			} else {
				blankline = true
			}
		}
		if !blankline && scanner.indentationLevel == 0 {
			if col == scanner.indentationStack[scanner.indentationCurrent] {
				if altcol != scanner.indentationAltStack[scanner.indentationCurrent] {
					pos = scanner.currentPosition
					return &token.Token{
						ID:          token.ERRORTOKEN,
						LineStart:   pos.Line,
						ColumnStart: pos.Column,
						LineEnd:     pos.Line,
						ColumnEnd:   pos.Column,
						Literal:     "",
					}
				}
			} else if col > scanner.indentationStack[scanner.indentationCurrent] {
				if scanner.indentationCurrent+1 >= MAXINDENT {
					scanner.state = errorcode.E_TOODEEP
					pos = scanner.currentPosition
					return &token.Token{
						ID:          token.ERRORTOKEN,
						LineStart:   pos.Line,
						ColumnStart: pos.Column,
						LineEnd:     pos.Line,
						ColumnEnd:   pos.Column,
						Literal:     "",
					}
				}
				if altcol <= scanner.indentationAltStack[scanner.indentationCurrent] {
					pos = scanner.currentPosition
					return &token.Token{
						ID:          token.ERRORTOKEN,
						LineStart:   pos.Line,
						ColumnStart: pos.Column,
						LineEnd:     pos.Line,
						ColumnEnd:   pos.Column,
						Literal:     "",
					}
				}
				scanner.indentationPending++
				scanner.indentationCurrent++
				scanner.indentationStack[scanner.indentationCurrent] = col
				scanner.indentationAltStack[scanner.indentationCurrent] = altcol

			} else {
				for scanner.indentationCurrent > 0 && col < scanner.indentationStack[scanner.indentationCurrent] {
					scanner.indentationPending--
					scanner.indentationCurrent--
				}
				if col != scanner.indentationStack[scanner.indentationCurrent] {
					scanner.state = errorcode.E_DEDENT
					pos = scanner.currentPosition
					return &token.Token{
						ID:          token.ERRORTOKEN,
						LineStart:   pos.Line,
						ColumnStart: pos.Column,
						LineEnd:     pos.Line,
						ColumnEnd:   pos.Column,
						Literal:     "",
					}
				}
				if altcol != scanner.indentationAltStack[scanner.indentationCurrent] {
					scanner.state = errorcode.E_DEDENT
					pos = scanner.currentPosition
					return &token.Token{
						ID:          token.ERRORTOKEN,
						LineStart:   pos.Line,
						ColumnStart: pos.Column,
						LineEnd:     pos.Line,
						ColumnEnd:   pos.Column,
						Literal:     "",
					}
				}
			}
		}
	}

	if scanner.indentationPending != 0 {
		if scanner.indentationPending < 0 {
			scanner.indentationPending++
			pos = scanner.currentPosition
			return &token.Token{
				ID:          token.DEDENT,
				LineStart:   pos.Line,
				ColumnStart: pos.Column,
				LineEnd:     pos.Line,
				ColumnEnd:   pos.Column,
				Literal:     "",
			}
		} else {
			scanner.indentationPending--
			pos = scanner.currentPosition
			return &token.Token{
				ID:          token.INDENT,
				LineStart:   pos.Line,
				ColumnStart: pos.Column,
				LineEnd:     pos.Line,
				ColumnEnd:   pos.Column + 4,
				Literal:     "    ",
			}
		}
	}

	pos = scanner.nextPosition()
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

		// Check for async/await
		literal := positions.String()
		if literal == "async" || literal == "await" {
			if scanner.asyncDef {
				switch literal {
				case "async":
					return positions.AsToken(token.ASYNC)
				case "await":
					return positions.AsToken(token.AWAIT)
				}
			} else if literal == "async" {
				nextToken := scanner.NextToken()
				if nextToken.ID == token.NAME && nextToken.Literal == "def" {
					scanner.asyncDef = true
					scanner.unreadToken(nextToken)
					return positions.AsToken(token.ASYNC)
				}
				scanner.unreadToken(nextToken)
			}
		}

		return positions.AsToken(token.NAME)
	case ch == '\n':
		scanner.atBol = true
		if blankline || scanner.indentationCurrent > 0 {
			goto next_line
		}
		return positions.AsToken(token.NEWLINE)
	case ch == '.':
		pos2 := scanner.nextPosition()
		if IsDigit(pos2.Char) {
			positions.Append(pos2)
			return scanner.parseNumber(positions, pos2.Char)
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
		return scanner.parseNumber(positions, ch)
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
		scanner.indentationLevel++
		break
	case ')', ']', '}':
		// Decrement indentation level
		scanner.indentationLevel--
		break
	}

	opId := GetOneCharTokenID(pos.Char)
	return positions.AsToken(opId)
}
