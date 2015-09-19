package ast

type BaseNode struct {
	ID NodeID
}

func (node *BaseNode) initBaseNode(id NodeID) {
	node.ID = id
}

func (node *BaseNode) Name() string {
	return NodeNames[node.ID]
}
