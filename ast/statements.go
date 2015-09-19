package ast

type StatementChildNode interface {
	Node
	stmtChildNode()
}

type Statement struct {
	ParentNode
}

func NewStatement() *Statement {
	node := &Statement{}
	node.initBaseNode(STMT)
	return node
}

func (node *Statement) SetChild(n StatementChildNode) { node.ParentNode.SetChild(n) }

type SimpleStatement struct {
	ListNode
}

func NewSimpleStatement() *SimpleStatement {
	node := &SimpleStatement{}
	node.initBaseNode(SIMPLE_STMT)
	node.initListNode()
	return node
}

func (node *SimpleStatement) stmtChildNode()           {}
func (node *SimpleStatement) Append(n *SmallStatement) { node.ListNode.Append(n) }

type CompoundStatement struct {
	BaseNode
}

func NewCompoundStatement() *CompoundStatement {
	node := &CompoundStatement{}
	node.initBaseNode(COMPOUND_STMT)
	return node
}
func (node *CompoundStatement) stmtChildNode() {}
