package ast

type StatementNode interface {
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
