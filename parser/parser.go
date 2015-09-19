package parser

import (
	"io"

	"github.com/brettlangdon/gython/ast"
	"github.com/brettlangdon/gython/errorcode"
	"github.com/brettlangdon/gython/scanner"
	"github.com/brettlangdon/gython/token"
)

type Parser struct {
	Errors      []*Error
	tokenizer   *scanner.Scanner
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

func (parser *Parser) addError(msg string) {
	parser.Errors = append(parser.Errors, &Error{
		Message: msg,
	})
}

func (parser *Parser) expect(tokID token.TokenID) {
	next := parser.nextToken()
	if next.ID != tokID {
		msg := "Unexpected token \"" + next.ID.String() + "\" expected \"" + tokID.String() + "\""
		parser.addError(msg)
	}
}

// compound_stmt: if_stmt | while_stmt | for_stmt | try_stmt | with_stmt | funcdef | classdef | decorated | async_stmt
func (parser *Parser) parseCompoundStatement() *ast.CompoundStatement {
	compoundStmt := ast.NewCompoundStatement()
	return compoundStmt
}

// or_test: and_test ('or' and_test)*
func (parser *Parser) parseOrTest() *ast.OrTest {
	orTest := ast.NewOrTest()
	return orTest
}

// test: or_test ['if' or_test 'else' test] | lambdef
func (parser *Parser) parseTest() *ast.Test {
	test := ast.NewTest()
	return test
}

// testlist_star_expr: (test|star_expr) (',' (test|star_expr))* [',']
func (parser *Parser) parseTestlistStarExpression() *ast.TestlistStarExpression {
	testlistStarExpression := ast.NewTestListStarExpression()

	var expr ast.TestlistStarExpressionChildNode
	expr = parser.parseTest()
	if expr == nil {
		return nil
	}
	testlistStarExpression.SetChild(expr)
	return testlistStarExpression
}

// expr_stmt: testlist_star_expr (augassign (yield_expr|testlist) |
//                      ('=' (yield_expr|testlist_star_expr))*)
func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	exprStmt := ast.NewExpressionStatement()
	testlistStarExpression := parser.parseTestlistStarExpression()
	if testlistStarExpression == nil {
		return nil
	}
	exprStmt.SetChild(testlistStarExpression)
	return exprStmt
}

// small_stmt: (expr_stmt | del_stmt | pass_stmt | flow_stmt |
//              import_stmt | global_stmt | nonlocal_stmt | assert_stmt)
func (parser *Parser) parseSmallStatment() *ast.SmallStatement {
	smallStmt := ast.NewSmallStatement()

	var stmt ast.SmallStatementChildNode
	stmt = parser.parseExpressionStatement()
	if stmt != nil {
		smallStmt.SetChild(stmt)
	}

	if stmt == nil {
		return nil
	}
	return smallStmt
}

// simple_stmt: small_stmt (';' small_stmt)* [';'] NEWLINE
func (parser *Parser) parseSimpleStatement() *ast.SimpleStatement {
	simpleStmt := ast.NewSimpleStatement()
	for {
		smallStmt := parser.parseSmallStatment()
		if smallStmt == nil {
			break
		}
		simpleStmt.Append(smallStmt)
		next := parser.nextToken()
		if next.ID != token.SEMI {
			parser.unreadToken(next)
			break
		}
	}
	parser.expect(token.NEWLINE)

	// no small statements found
	if simpleStmt.Length() == 0 {
		return nil
	}
	return simpleStmt
}

// stmt: simple_stmt | compound_stmt
func (parser *Parser) parseStatement() *ast.Statement {
	var next ast.StatementChildNode
	next = parser.parseSimpleStatement()
	if next == nil {
		next = parser.parseCompoundStatement()
	}

	if next == nil {
		return nil
	}

	stmt := ast.NewStatement()
	stmt.SetChild(next)
	return stmt
}

// file_input: (NEWLINE | stmt)* ENDMARKER
func (parser *Parser) parseFileInput() *ast.FileInput {
	root := ast.NewFileInput()
	for parser.tokenizer.State() == errorcode.E_OK {
		next := parser.nextToken()
		if next.ID == token.NEWLINE {
			// root.Append(ast.NewTokenNode(next))
			continue
		} else if next.ID == token.ENDMARKER {
			// Unread, so we can read in the expected value later
			parser.unreadToken(next)
			break
		} else {
			parser.unreadToken(next)
			stmt := parser.parseStatement()
			if stmt == nil {
				break
			}
			root.Append(stmt)
			break
		}
	}

	parser.expect(token.ENDMARKER)

	return root
}

func ParseReader(r io.Reader) (*ast.FileInput, *Parser) {
	parser := &Parser{
		tokenizer:   scanner.NewScanner(r),
		tokenBuffer: make([]*token.Token, 0),
	}

	return parser.parseFileInput(), parser
}
