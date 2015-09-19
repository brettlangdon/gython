package ast

type TestlistStarExpressionChildNode interface {
	Node
	testlistStarExpressionChild()
}

type TestlistStarExpression struct {
	ParentNode
}

func NewTestListStarExpression() *TestlistStarExpression {
	node := &TestlistStarExpression{}
	node.initBaseNode(TESTLIST_STAR_EXPR)
	return node
}

func (node *TestlistStarExpression) SetChild(n TestlistStarExpressionChildNode) {
	node.ParentNode.SetChild(n)
}
