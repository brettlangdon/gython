package ast

type CompoundStatement struct {
	Statement
}

func NewCompoundStatement() *CompoundStatement {
	node := &CompoundStatement{}
	node.initBaseNode(SIMPLE_STMT)
	return node
}

func (node *CompoundStatement) StatementNode() {}
