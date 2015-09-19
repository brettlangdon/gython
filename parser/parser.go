package parser

import (
	"io"

	"github.com/brettlangdon/gython/ast"
	"github.com/brettlangdon/gython/errorcode"
	"github.com/brettlangdon/gython/scanner"
	"github.com/brettlangdon/gython/token"
)

type Parser struct {
	tokenizer *scanner.Scanner

	tokenBuffer []*token.Token
}

func (parser *Parser) nextToken() *token.Token {
	if len(parser.tokenBuffer) > 0 {
		last := len(parser.tokenBuffer) - 1
		next := parser.tokenBuffer[last]
		parser.tokenBuffer = parser.tokenBuffer[:last]
		return next
	}

	return parser.tokenizer.NextToken()
}

func (parser *Parser) unreadToken(tok *token.Token) {
	parser.tokenBuffer = append(parser.tokenBuffer, tok)
}

func (parser *Parser) parseFileInput() *ast.FileInput {
	root := ast.NewFileInput()
	for parser.tokenizer.State() == errorcode.E_OK {
		next := parser.nextToken()
		if next.ID == token.NEWLINE {
			root.AppendToken(next)
		} else if next.ID == token.ENDMARKER {
			break
		} else {
			parser.unreadToken(next)
			// TODO: parser.parseStatement()
			break
		}
	}

	return root
}

func ParseReader(r io.Reader) *ast.FileInput {
	parser := &Parser{
		tokenizer:   scanner.NewScanner(r),
		tokenBuffer: make([]*token.Token, 0),
	}

	return parser.parseFileInput()
}
