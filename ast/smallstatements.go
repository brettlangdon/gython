package ast

type SmallStatementNode interface {
	Node
	smallStmtNode()
}

type SmallStatement struct {
	ParentNode
}

func NewSmallStatement() *SmallStatement {
	node := &SmallStatement{}
	node.initBaseNode(SMALL_STMT)
	return node
}

func (node *SmallStatement) SetChild(n SmallStatementNode) { node.ParentNode.SetChild(n) }

type ExpressionStatement struct {
	ParentNode
	Expression *TestlistStarExpression
}

func NewExpressionStatement() *ExpressionStatement {
	node := &ExpressionStatement{}
	node.initBaseNode(EXPR_STMT)
	return node
}

func (node *ExpressionStatement) smallStmtNode()                     {}
func (node *ExpressionStatement) SetChild(n *TestlistStarExpression) { node.ParentNode.SetChild(n) }
