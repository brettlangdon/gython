package ast

type TestlistStartExpr struct {
	BaseNode
}

func NewTestListStartExpr() *TestlistStartExpr {
	node := &TestlistStartExpr{}
	node.initBaseNode(TESTLIST_STAR_EXPR)
	return node
}
