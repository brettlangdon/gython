package ast

import "github.com/brettlangdon/gython/token"

type ListNode struct {
	children []interface{}
}

func (node *ListNode) initChildren() {
	node.children = make([]interface{}, 0)
}

func (node *ListNode) AppendToken(t *token.Token) {
	node.children = append(node.children, t)
}

func (node *ListNode) AppendNode(n Node) {
	node.children = append(node.children, n)
}
