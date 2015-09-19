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

func (node *BaseNode) Repr() []interface{} {
	out := make([]interface{}, 0)
	out = append(out, node.Name())
	return out
}
