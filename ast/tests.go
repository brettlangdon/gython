package ast

type TestNode interface {
	Node
	test()
}

type Test struct {
	ParentNode
}

func NewTest() *Test {
	node := &Test{}
	node.initBaseNode(TEST)
	return node
}

func (node *Test) SetChild(n TestNode) { node.ParentNode.SetChild(n) }
