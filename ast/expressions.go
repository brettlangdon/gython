package ast

type ExpressionNode interface {
	Node
	exprNode()
}

type TestlistStarExpression struct {
	ParentNode
}

func NewTestListStarExpression() *TestlistStarExpression {
	node := &TestlistStarExpression{}
	node.initBaseNode(TESTLIST_STAR_EXPR)
	return node
}
