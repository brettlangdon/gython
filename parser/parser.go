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

func (parser *Parser) expect(tokID token.TokenID) bool {
	next := parser.nextToken()
	if next.ID != tokID {
		msg := "Unexpected token \"" + next.ID.String() + "\" expected \"" + tokID.String() + "\""
		parser.addError(msg)
		return false
	}
	return true
}

func (parser *Parser) expectLiteral(literal string) bool {
	next := parser.nextToken()
	if !next.IsLiteral(literal) {
		msg := "Unexpected literal \"" + next.Literal + "\" expected \"" + literal + "\""
		parser.addError(msg)
		return false
	}
	return true
}

// compound_stmt: if_stmt | while_stmt | for_stmt | try_stmt | with_stmt | funcdef | classdef | decorated | async_stmt
func (parser *Parser) parseCompoundStatement() *ast.CompoundStatement {
	compoundStmt := ast.NewCompoundStatement()
	return compoundStmt
}

// atom: ('(' [yield_expr|testlist_comp] ')' |
//        '[' [testlist_comp] ']' |
//        '{' [dictorsetmaker] '}' |
//        NAME | NUMBER | STRING+ | '...' | 'None' | 'True' | 'False')
func (parser *Parser) parseAtom() *ast.Atom {
	atom := ast.NewAtom()
	next := parser.nextToken()
	switch next.ID {
	case token.NAME, token.NUMBER, token.ELLIPSIS:
		atom.Append(ast.NewTokenNode(next))
	case token.STRING:
		atom.Append(ast.NewTokenNode(next))
		for {
			next := parser.nextToken()
			if next.ID != token.STRING {
				parser.unreadToken(next)
				break
			}
			atom.Append(ast.NewTokenNode(next))
		}
	}
	return atom
}

// trailer: '(' [arglist] ')' | '[' subscriptlist ']' | '.' NAME
func (parser *Parser) parseTrailer() *ast.Trailer {
	trailer := ast.NewTrailer()
	next := parser.nextToken()
	switch next.ID {
	case token.LPAR:
		next2 := parser.nextToken()
		if next2.ID != token.RPAR {
			return nil
		}
		trailer.Append(ast.NewTokenNode(next))
		trailer.Append(ast.NewTokenNode(next2))
	case token.LBRACE:
		next2 := parser.nextToken()
		if next2.ID != token.RBRACE {
			return nil
		}
		trailer.Append(ast.NewTokenNode(next))
		trailer.Append(ast.NewTokenNode(next2))
	case token.DOT:
		next2 := parser.nextToken()
		if next2.ID == token.NAME {
			trailer.Append(ast.NewTokenNode(next))
			trailer.Append(ast.NewTokenNode(next2))
		} else {
			parser.addError("Expected \"NAME\" instead found \"" + next.ID.String() + "\"")
			return nil
		}
	default:
		parser.unreadToken(next)
		return nil
	}
	return trailer
}

// atom_expr: [AWAIT] atom trailer*
func (parser *Parser) parseAtomExpression() *ast.AtomExpression {
	expr := ast.NewAtomExpression()
	next := parser.nextToken()
	if next.ID == token.AWAIT {
		expr.Append(ast.NewTokenNode(next))
	} else {
		parser.unreadToken(next)
	}

	atom := parser.parseAtom()
	if atom == nil {
		return nil
	}
	expr.Append(atom)
	for {
		trailer := parser.parseTrailer()
		if trailer == nil {
			break
		}
		expr.Append(trailer)
	}

	return expr
}

// power: atom_expr ['**' factor]
func (parser *Parser) parsePower() *ast.Power {
	power := ast.NewPower()
	atomExpr := parser.parseAtomExpression()
	if atomExpr == nil {
		return nil
	}
	power.Append(atomExpr)

	next := parser.nextToken()
	if next.ID == token.DOUBLESTAR {
		factor := parser.parseFactor()
		if factor == nil {
			return nil
		} else {
			power.Append(factor)
		}
	} else {
		parser.unreadToken(next)
	}
	return power
}

// factor: ('+'|'-'|'~') factor | power
func (parser *Parser) parseFactor() *ast.Factor {
	factor := ast.NewFactor()
	next := parser.nextToken()
	switch next.ID {
	case token.PLUS, token.MINUS, token.TILDE:
		node := parser.parseFactor()
		if node == nil {
			return nil
		}
		factor.Append(ast.NewTokenNode(next))
		factor.Append(node)
	default:
		parser.unreadToken(next)
		power := parser.parsePower()
		if power == nil {
			return nil
		}
		factor.Append(power)
	}

	return factor
}

// term: factor (('*'|'@'|'/'|'%'|'//') factor)*
func (parser *Parser) parseTerm() *ast.Term {
	term := ast.NewTerm()
	factor := parser.parseFactor()
	if factor == nil {
		return nil
	}
	term.Append(factor)
	for {
		next := parser.nextToken()
		if next.ID != token.STAR && next.ID != token.AMPER && next.ID != token.SLASH && next.ID != token.PERCENT && next.ID != token.DOUBLESLASH {
			parser.unreadToken(next)
			break
		}
		factor := parser.parseFactor()
		if factor == nil {
			return nil
		}
		term.Append(factor)
	}
	return term
}

// arith_expr: term (('+'|'-') term)*
func (parser *Parser) parseArithmetricExpression() *ast.ArithmeticExpression {
	expr := ast.NewArithmeticExpression()
	term := parser.parseTerm()
	if term == nil {
		return nil
	}
	expr.Append(term)
	for {
		next := parser.nextToken()
		if next.ID != token.PLUS || next.ID != token.MINUS {
			parser.unreadToken(next)
			break
		}
		term := parser.parseTerm()
		if term == nil {
			return nil
		}
		expr.Append(term)
	}
	return expr
}

// shift_expr: arith_expr (('<<'|'>>') arith_expr)*
func (parser *Parser) parseShiftExpression() *ast.ShiftExpression {
	expr := ast.NewShiftExpression()
	arithExpr := parser.parseArithmetricExpression()
	if arithExpr == nil {
		return nil
	}
	expr.Append(arithExpr)
	for {
		next := parser.nextToken()
		if next.ID != token.LEFTSHIFT && next.ID != token.RIGHTSHIFT {
			parser.unreadToken(next)
			break
		}
		expr.Append(ast.NewTokenNode(next))
		arithExpr := parser.parseArithmetricExpression()
		if arithExpr == nil {
			return nil
		}
		expr.Append(arithExpr)
	}
	return expr
}

// and_expr: shift_expr ('&' shift_expr)*
func (parser *Parser) parseAndExpression() *ast.AndExpression {
	expr := ast.NewAndExpression()
	shiftExpr := parser.parseShiftExpression()
	if shiftExpr == nil {
		return nil
	}
	expr.Append(shiftExpr)
	for {
		next := parser.nextToken()
		if next.ID != token.AMPER {
			parser.unreadToken(next)
			break
		}
		shiftExpr := parser.parseShiftExpression()
		if shiftExpr == nil {
			return nil
		}
		expr.Append(shiftExpr)
	}
	return expr
}

// xor_expr: and_expr ('^' and_expr)*
func (parser *Parser) parseXorExpression() *ast.XorExpression {
	expr := ast.NewXorExpression()
	andExpr := parser.parseAndExpression()
	if andExpr == nil {
		return nil
	}
	expr.Append(andExpr)
	for {
		next := parser.nextToken()
		if next.ID != token.CIRCUMFLEX {
			parser.unreadToken(next)
			break
		}
		andExpr := parser.parseAndExpression()
		if andExpr == nil {
			return nil
		}
		expr.Append(andExpr)
	}
	return expr
}

// expr: xor_expr ('|' xor_expr)*
func (parser *Parser) parseExpression() *ast.Expression {
	expr := ast.NewExpression()
	xorExpr := parser.parseXorExpression()
	if xorExpr == nil {
		return nil
	}
	expr.Append(xorExpr)
	for {
		next := parser.nextToken()
		if next.ID != token.VBAR {
			parser.unreadToken(next)
			break
		}
		xorExpr := parser.parseXorExpression()
		if xorExpr == nil {
			return nil
		}
		expr.Append(xorExpr)
	}
	return expr
}

// comparison: expr (comp_op expr)*
func (parser *Parser) parseComparison() *ast.Comparison {
	comparison := ast.NewComparison()
	expr := parser.parseExpression()
	if expr == nil {
		return nil
	}
	comparison.Append(expr)

	for {
		// comp_op: '<'|'>'|'=='|'>='|'<='|'<>'|'!='|'in'|'not' 'in'|'is'|'is' 'not'
		compOp := true
		next := parser.nextToken()
		switch next.Literal {
		case "<", ">", "==", ">=", "<=", "<>", "!=", "in":
			comparison.Append(ast.NewTokenNode(next))
		case "is":
			comparison.Append(ast.NewTokenNode(next))
			next2 := parser.nextToken()
			if next2.Literal == "not" {
				comparison.Append(ast.NewTokenNode(next2))
			} else {
				parser.unreadToken(next2)
			}
		case "not":
			next2 := parser.nextToken()
			if next2.Literal == "in" {
				comparison.Append(ast.NewTokenNode(next))
				comparison.Append(ast.NewTokenNode(next2))
			} else {
				parser.unreadToken(next2)
				parser.unreadToken(next)
				compOp = false
			}
		default:
			parser.unreadToken(next)
			compOp = false
		}
		if compOp == false {
			break
		}
		expr := parser.parseExpression()
		if expr == nil {
			return nil
		}
		comparison.Append(expr)
	}

	return comparison
}

// not_test: 'not' not_test | comparison
func (parser *Parser) parseNotTest() *ast.NotTest {
	notTest := ast.NewNotTest()
	next := parser.nextToken()
	if next.IsLiteral("not") {
		test := parser.parseNotTest()
		if test == nil {
			return nil
		}
		notTest.SetChild(test)
	} else {
		parser.unreadToken(next)
		comparison := parser.parseComparison()
		if comparison == nil {
			return nil
		}
		notTest.SetChild(comparison)
	}
	return notTest
}

// and_test: not_test ('and' not_test)*
func (parser *Parser) parseAndTest() *ast.AndTest {
	andTest := ast.NewAndTest()
	notTest := parser.parseNotTest()
	if notTest == nil {
		return nil
	}
	andTest.Append(notTest)
	for {
		next := parser.nextToken()
		if !next.IsLiteral("and") {
			parser.unreadToken(next)
			break
		}
		notTest = parser.parseNotTest()
		if notTest == nil {
			return nil
		}
		andTest.Append(notTest)
	}

	return andTest
}

// or_test: and_test ('or' and_test)*
func (parser *Parser) parseOrTest() *ast.OrTest {
	orTest := ast.NewOrTest()
	andTest := parser.parseAndTest()
	if andTest == nil {
		return nil
	}
	orTest.Append(andTest)
	for {
		next := parser.nextToken()
		if !next.IsLiteral("and") {
			parser.unreadToken(next)
			break
		}
		andTest = parser.parseAndTest()
		if andTest == nil {
			return nil
		}
		orTest.Append(andTest)
	}
	return orTest
}

// test: or_test ['if' or_test 'else' test] | lambdef
func (parser *Parser) parseTest() *ast.Test {
	test := ast.NewTest()

	orTest := parser.parseOrTest()
	if orTest != nil {
		test.Append(orTest)
		next := parser.nextToken()
		// Do not use `parser.expectLiteral`, this next part is optional
		if next.IsLiteral("if") {
			orTest = parser.parseOrTest()
			if orTest != nil {
				test.Append(orTest)
				if parser.expectLiteral("else") {
					elseTest := parser.parseTest()
					if elseTest == nil {
						return nil
					}
					test.Append(test)
				}
			}
		} else {
			parser.unreadToken(next)
		}
	} else {
		// TODO: parser.parseLambDef()
	}
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
	expr := parser.parseTestlistStarExpression()
	if expr == nil {
		return nil
	}
	exprStmt.Append(expr)

	if false {
	} else {
		for {
			next := parser.nextToken()
			if next.ID != token.EQUAL {
				parser.unreadToken(next)
				break
			}
			exprStmt.Append(ast.NewTokenNode(next))
			expr := parser.parseTestlistStarExpression()
			if expr == nil {
				return nil
			}
			exprStmt.Append(expr)
		}
	}

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
