package ast

type StatementChildNode interface {
	Node
	stmtChildNode()
}

type Statement struct {
	ParentNode
}

func NewStatement() *Statement {
	node := &Statement{}
	node.initBaseNode(STMT)
	return node
}

func (node *Statement) fileInputChild()               {}
func (node *Statement) SetChild(n StatementChildNode) { node.ParentNode.SetChild(n) }

type SimpleStatementChildNode interface {
	Node
	simpleStatementChild()
}

type SimpleStatement struct {
	ListNode
}

func NewSimpleStatement() *SimpleStatement {
	node := &SimpleStatement{}
	node.initBaseNode(SIMPLE_STMT)
	node.initListNode()
	return node
}

func (node *SimpleStatement) stmtChildNode()                    {}
func (node *SimpleStatement) Append(n SimpleStatementChildNode) { node.ListNode.Append(n) }

type CompoundStatement struct {
	BaseNode
}

func NewCompoundStatement() *CompoundStatement {
	node := &CompoundStatement{}
	node.initBaseNode(COMPOUND_STMT)
	return node
}
func (node *CompoundStatement) stmtChildNode() {}

type SmallStatementChildNode interface {
	Node
	smallStmtChild()
}

type SmallStatement struct {
	ParentNode
}

func NewSmallStatement() *SmallStatement {
	node := &SmallStatement{}
	node.initBaseNode(SMALL_STMT)
	return node
}

func (node *SmallStatement) simpleStatementChild()              {}
func (node *SmallStatement) SetChild(n SmallStatementChildNode) { node.ParentNode.SetChild(n) }

type ExpressionStatementChildNode interface {
	Node
	expressionStatementChild()
}

type ExpressionStatement struct {
	ListNode
	Expression *TestlistStarExpression
}

func NewExpressionStatement() *ExpressionStatement {
	node := &ExpressionStatement{}
	node.initBaseNode(EXPR_STMT)
	node.initListNode()
	return node
}

func (node *ExpressionStatement) smallStmtChild()                       {}
func (node *ExpressionStatement) Append(n ExpressionStatementChildNode) { node.ListNode.Append(n) }
