package ast

type SmallStatementNode interface {
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
