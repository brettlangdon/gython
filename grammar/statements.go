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
	rule := &Statement{}
	rule.initBaseNode(symbol.STMT)
	return rule
}

func (rule *Statement) fileInputChild()           {}
func (rule *Statement) SetChild(n StatementChild) { rule.ParentNode.SetChild(n) }

type SimpleStatementChild interface {
	Node
	simpleStatementChild()
}

type SimpleStatement struct {
	ListNode
}

func NewSimpleStatement() *SimpleStatement {
	rule := &SimpleStatement{}
	rule.initBaseNode(symbol.SIMPLE_STMT)
	rule.initListNode()
	return rule
}

func (rule *SimpleStatement) stmtChild()                    {}
func (rule *SimpleStatement) Append(n SimpleStatementChild) { rule.ListNode.Append(n) }

type CompoundStatement struct {
	BaseNode
}

func NewCompoundStatement() *CompoundStatement {
	rule := &CompoundStatement{}
	rule.initBaseNode(symbol.COMPOUND_STMT)
	return rule
}
func (rule *CompoundStatement) stmtChild() {}

type SmallStatementChild interface {
	Node
	smallStmtChild()
}

type SmallStatement struct {
	ParentNode
}

func NewSmallStatement() *SmallStatement {
	rule := &SmallStatement{}
	rule.initBaseNode(symbol.SMALL_STMT)
	return rule
}

func (rule *SmallStatement) simpleStatementChild()          {}
func (rule *SmallStatement) SetChild(n SmallStatementChild) { rule.ParentNode.SetChild(n) }

type ExpressionStatementChild interface {
	Node
	expressionStatementChild()
}

type ExpressionStatement struct {
	ListNode
	Expression *TestlistStarExpression
}

func NewExpressionStatement() *ExpressionStatement {
	rule := &ExpressionStatement{}
	rule.initBaseNode(symbol.EXPR_STMT)
	rule.initListNode()
	return rule
}

func (rule *ExpressionStatement) smallStmtChild()                   {}
func (rule *ExpressionStatement) Append(n ExpressionStatementChild) { rule.ListNode.Append(n) }
