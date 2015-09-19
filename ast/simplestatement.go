package ast

type SimpleStatement struct {
	BaseNode
	Statements []*SmallStatement
}

func NewSimpleStatement() *SimpleStatement {
	node := &SimpleStatement{}
	node.initBaseNode(SIMPLE_STMT)
	return node
}

func (node *SimpleStatement) StatementNode() {}

func (node *SimpleStatement) AppendSmallStatement(n *SmallStatement) {
	node.Statements = append(node.Statements, n)
}
