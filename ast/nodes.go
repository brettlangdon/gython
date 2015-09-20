package ast

import "github.com/brettlangdon/gython/token"

type Node interface {
	Name() string
	Repr() []interface{}
}

type TokenNode struct {
	Token *token.Token
}

func NewTokenNode(tok *token.Token) *TokenNode {
	return &TokenNode{
		Token: tok,
	}
}
func (node *TokenNode) atomChild()                {}
func (node *TokenNode) atomExpressionChild()      {}
func (node *TokenNode) comparisonChild()          {}
func (node *TokenNode) expressionStatementChild() {}
func (node *TokenNode) factorChild()              {}
func (node *TokenNode) fileInputChild()           {}
func (node *TokenNode) shiftExpressionChild()     {}
func (node *TokenNode) simpleStatementChild()     {}
func (node *TokenNode) trailerChild()             {}
func (node *TokenNode) Name() string              { return token.TokenNames[node.Token.ID] }
func (node *TokenNode) Repr() []interface{} {
	parts := make([]interface{}, 0)
	parts = append(parts, node.Name())
	return append(parts, node.Token.Literal)
}

type BaseNode struct {
	ID    NodeID
	child Node
}

func (node *BaseNode) initBaseNode(id NodeID)      { node.ID = id }
func (node *BaseNode) Name() string                { return NodeNames[node.ID] }
func (node *BaseNode) Repr() (parts []interface{}) { return append(parts, node.Name()) }

type ParentNode struct {
	BaseNode
	child Node
}

func (node *ParentNode) SetChild(n Node) { node.child = n }
func (node *ParentNode) Child() Node     { return node.child }
func (node *ParentNode) Repr() (parts []interface{}) {
	parts = node.BaseNode.Repr()
	child := node.Child()
	if child != nil {
		parts = append(parts, child.Repr())
	}
	return parts
}

type ListNode struct {
	BaseNode
	children []Node
}

func (node *ListNode) initListNode()    { node.children = make([]Node, 0) }
func (node *ListNode) Length() int      { return len(node.children) }
func (node *ListNode) Children() []Node { return node.children }
func (node *ListNode) Append(n Node)    { node.children = append(node.children, n) }
func (node *ListNode) Repr() (parts []interface{}) {
	parts = node.BaseNode.Repr()
	children := node.Children()
	for _, child := range children {
		parts = append(parts, child.Repr())
	}
	return parts
}
