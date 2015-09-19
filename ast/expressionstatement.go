package ast

type ExpressionStatement struct {
	BaseNode
}

func NewExpressionStatement() *ExpressionStatement {
	node := &ExpressionStatement{}
	node.initBaseNode(EXPR_STMT)
	return node
}

func (node *ExpressionStatement) SmallStatementNode() {}
