package ast

type SmallStatementChildNode interface {
	Node
	smallStmtChildNode()
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

type ExpressionStatement struct {
	ParentNode
	Expression *TestlistStarExpression
}

func NewExpressionStatement() *ExpressionStatement {
	node := &ExpressionStatement{}
	node.initBaseNode(EXPR_STMT)
	return node
}

func (node *ExpressionStatement) smallStmtChildNode()                {}
func (node *ExpressionStatement) SetChild(n *TestlistStarExpression) { node.ParentNode.SetChild(n) }
