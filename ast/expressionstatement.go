package ast

type ExpressionStatement struct {
	BaseNode
	Expression *TestlistStartExpr
}

func NewExpressionStatement() *ExpressionStatement {
	node := &ExpressionStatement{}
	node.initBaseNode(EXPR_STMT)
	return node
}

func (node *ExpressionStatement) SmallStatementNode() {}

func (node *ExpressionStatement) Repr() []interface{} {
	out := node.BaseNode.Repr()
	out = append(out, node.Expression.Repr())
	return out
}
