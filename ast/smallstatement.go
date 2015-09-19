package ast

type SmallStatementNode interface {
	Node
	SmallStatementNode()
}

type SmallStatement struct {
	BaseNode
	Statement SmallStatementNode
}

func NewSmallStatement() *SmallStatement {
	node := &SmallStatement{}
	node.initBaseNode(SMALL_STMT)
	return node
}

func (node *SmallStatement) Repr() []interface{} {
	out := node.BaseNode.Repr()
	out = append(out, node.Statement.Repr())
	return out
}
