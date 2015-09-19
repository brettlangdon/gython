package ast

import "github.com/brettlangdon/gython/token"

type ListNode struct {
	BaseNode
	children []interface{}
}

func (node *ListNode) initListNode() {
	node.children = make([]interface{}, 0)
}

func (node *ListNode) AppendToken(t *token.Token) {
	node.children = append(node.children, t)
}

func (node *ListNode) AppendNode(n Node) {
	node.children = append(node.children, n)
}

func (node *ListNode) Children() []interface{} {
	return node.children
}

func (node *ListNode) Length() int {
	return len(node.children)
}

func (node *ListNode) Repr() []interface{} {
	out := node.BaseNode.Repr()
	for _, child := range node.Children() {
		switch child.(type) {
		case Node:
			out = append(out, child.(Node).Repr())
		default:
			out = append(out, child)
		}
	}
	return out
}
