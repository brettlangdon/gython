package ast

import (
	"fmt"
	"strconv"

	"github.com/brettlangdon/gython/grammar"
	"github.com/brettlangdon/gython/symbol"
	"github.com/brettlangdon/gython/token"
)

func ASTFromGrammar(root *grammar.FileInput) (Mod, error) {
	mod := NewModule()

	for _, child := range root.Children() {
		if child.ID() == symbol.STMT {
			stmt := astForStatement(child.(*grammar.Statement))
			mod.Append(stmt)
		}
	}

	return mod, nil
}

func isToken(node grammar.Node, tokId token.TokenID) bool {
	if n, isTokenNode := node.(*grammar.TokenNode); isTokenNode {
		return n.Token.ID == tokId
	}
	return false
}

func astForStatement(root *grammar.Statement) Statement {
	stmt := root.Child()

	if stmt.ID() == symbol.SIMPLE_STMT {
		stmt = stmt.(*grammar.SimpleStatement).Children()[0]
	}

	switch stmt := stmt.(type) {
	case *grammar.SmallStatement:
		switch child := stmt.Child().(type) {
		case *grammar.ExpressionStatement:
			return astForExpressionStatement(child)
		}

	case *grammar.CompoundStatement:
		fmt.Println(stmt)
	}
	return nil
}

func astForExpressionStatement(root *grammar.ExpressionStatement) Statement {
	children := root.Children()
	if len(children) == 1 {

	} else if "todo" == "augassign" {
	} else {
		if !isToken(children[1], token.EQUAL) {
			return nil
		}
		length := len(children)
		var value Expression
		switch child := children[length-1].(type) {
		case *grammar.TestlistStarExpression:
			value = astForTestList(child)
		default:
			value = nil
		}

		assign := NewAssign(value)
		for i := 0; i < length-2; i++ {
			target := astForTestList(children[i].(grammar.ExpressionStatementChild))
			switch target := target.(type) {
			case *Name:
				target.Context = NewStore()
			}
			assign.Append(target)
		}
		return assign
	}
	return nil
}

func astForTestList(root grammar.ExpressionStatementChild) Expression {
	switch root := root.(type) {
	case *grammar.TestlistStarExpression:
		if root.Length() == 1 {
			return astForExpression(root.Children()[0])
		}
	}
	return nil
}

func astForExpression(root grammar.Node) Expression {
	switch root := root.(type) {
	case *grammar.Test:
		if root.Length() == 1 {
			return astForExpression(root.Children()[0])
		}
	case *grammar.OrTest:
		if root.Length() == 1 {
			return astForExpression(root.Children()[0])
		}
	case *grammar.AndTest:
		if root.Length() == 1 {
			return astForExpression(root.Children()[0])
		}
	case *grammar.NotTest:
		if root.Length() == 1 {
			return astForExpression(root.Children()[0])
		}
	case *grammar.Comparison:
		if root.Length() == 1 {
			return astForExpression(root.Children()[0])
		}
	case *grammar.Expression:
		if root.Length() == 1 {
			return astForExpression(root.Children()[0])
		}
	case *grammar.XorExpression:
		if root.Length() == 1 {
			return astForExpression(root.Children()[0])
		}
	case *grammar.AndExpression:
		if root.Length() == 1 {
			return astForExpression(root.Children()[0])
		}
	case *grammar.ShiftExpression:
		if root.Length() == 1 {
			return astForExpression(root.Children()[0])
		}
	case *grammar.ArithmeticExpression:
		if root.Length() == 1 {
			return astForExpression(root.Children()[0])
		}
	case *grammar.Term:
		if root.Length() == 1 {
			return astForExpression(root.Children()[0])
		}
	case *grammar.Factor:
		if root.Length() == 1 {
			return astForExpression(root.Children()[0])
		}
	case *grammar.Power:
		return astForPower(root)
	default:
		fmt.Println(symbol.SymbolNames[root.ID()])
	}
	return nil
}

func astForPower(root *grammar.Power) Expression {
	children := root.Children()
	var expr Expression
	if child, isAtomExpr := children[0].(*grammar.AtomExpression); isAtomExpr {
		expr = astForAtomExpression(child)
	} else {
		return nil
	}

	if len(children) == 1 {
		return expr
	}

	return nil
}

func astForAtomExpression(root *grammar.AtomExpression) Expression {
	children := root.Children()
	switch child := children[0].(type) {
	case *grammar.Atom:
		return astForAtom(child)
	}
	return nil
}

func astForAtom(root *grammar.Atom) Expression {
	children := root.Children()
	if len(children) == 1 {
		switch child := children[0].(type) {
		case *grammar.TokenNode:
			switch child.Token.ID {
			case token.NAME:
				// TODO: Check for "None", "True", and "False"
				return NewName(child.Token.Literal, NewLoad())
			case token.NUMBER:
				value, err := strconv.ParseInt(child.Token.Literal, 10, 64)
				if err != nil {
					return nil
				}
				return NewNum(value)
			}
		}
	}
	return nil
}
