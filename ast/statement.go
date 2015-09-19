package ast

type StatementNode interface {
	Node
	StatementNode()
}

type Statement struct {
	BaseNode
	Statement StatementNode
}

func NewStatement() *Statement {
	node := &Statement{}
	node.initBaseNode(STMT)
	return node
}

func (node *Statement) StatementNode() {}

func (node *Statement) Repr() []interface{} {
	out := node.BaseNode.Repr()
	out = append(out, node.Statement.Repr())
	return out
}
