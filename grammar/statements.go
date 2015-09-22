package grammar

import "github.com/brettlangdon/gython/symbol"

type StatementChild interface {
	Node
	stmtChild()
}

type Statement struct {
	ParentNode
}

func NewStatement() *Statement {
	node := &Statement{}
	node.initBaseNode(symbol.STMT)
	return node
}

func (node *Statement) fileInputChild()           {}
func (node *Statement) SetChild(n StatementChild) { node.ParentNode.SetChild(n) }

type SimpleStatementChild interface {
	Node
	simpleStatementChild()
}

type SimpleStatement struct {
	ListNode
}

func NewSimpleStatement() *SimpleStatement {
	node := &SimpleStatement{}
	node.initBaseNode(symbol.SIMPLE_STMT)
	node.initListNode()
	return node
}

func (node *SimpleStatement) stmtChild()                    {}
func (node *SimpleStatement) Append(n SimpleStatementChild) { node.ListNode.Append(n) }

type CompoundStatement struct {
	BaseNode
}

func NewCompoundStatement() *CompoundStatement {
	node := &CompoundStatement{}
	node.initBaseNode(symbol.COMPOUND_STMT)
	return node
}
func (node *CompoundStatement) stmtChild() {}

type SmallStatementChild interface {
	Node
	smallStmtChild()
}

type SmallStatement struct {
	ParentNode
}

func NewSmallStatement() *SmallStatement {
	node := &SmallStatement{}
	node.initBaseNode(symbol.SMALL_STMT)
	return node
}

func (node *SmallStatement) simpleStatementChild()          {}
func (node *SmallStatement) SetChild(n SmallStatementChild) { node.ParentNode.SetChild(n) }

type ExpressionStatementChild interface {
	Node
	expressionStatementChild()
}

type ExpressionStatement struct {
	ListNode
	Expression *TestlistStarExpression
}

func NewExpressionStatement() *ExpressionStatement {
	node := &ExpressionStatement{}
	node.initBaseNode(symbol.EXPR_STMT)
	node.initListNode()
	return node
}

func (node *ExpressionStatement) smallStmtChild()                   {}
func (node *ExpressionStatement) Append(n ExpressionStatementChild) { node.ListNode.Append(n) }
