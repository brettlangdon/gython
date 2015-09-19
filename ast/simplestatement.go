package ast

type SimpleStatement struct {
	ListNode
}

func NewSimpleStatement() *SimpleStatement {
	node := &SimpleStatement{}
	node.initBaseNode(SIMPLE_STMT)
	return node
}

func (node *SimpleStatement) StatementNode() {}

func (node *SimpleStatement) AppendNode(n *SmallStatement) {
	node.ListNode.AppendNode(n)
}
