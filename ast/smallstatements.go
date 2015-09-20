package ast

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
