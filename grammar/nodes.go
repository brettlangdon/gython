package grammar

import (
	"fmt"

	"github.com/brettlangdon/gython/symbol"
	"github.com/brettlangdon/gython/token"
)

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
func (rule *TokenNode) atomChild()                {}
func (rule *TokenNode) atomExpressionChild()      {}
func (rule *TokenNode) comparisonChild()          {}
func (rule *TokenNode) expressionStatementChild() {}
func (rule *TokenNode) factorChild()              {}
func (rule *TokenNode) fileInputChild()           {}
func (rule *TokenNode) shiftExpressionChild()     {}
func (rule *TokenNode) simpleStatementChild()     {}
func (rule *TokenNode) trailerChild()             {}
func (rule *TokenNode) Name() string              { return token.TokenNames[rule.Token.ID] }
func (rule *TokenNode) Repr() []interface{} {
	parts := make([]interface{}, 0)
	parts = append(parts, rule.Name())
	literal := fmt.Sprintf("%#v", rule.Token.Literal)
	return append(parts, literal)
}

type BaseNode struct {
	ID    symbol.SymbolID
	child Node
}

func (rule *BaseNode) initBaseNode(id symbol.SymbolID) { rule.ID = id }
func (rule *BaseNode) Name() string                    { return symbol.SymbolNames[rule.ID] }
func (rule *BaseNode) Repr() (parts []interface{})     { return append(parts, rule.Name()) }

type ParentNode struct {
	BaseNode
	child Node
}

func (rule *ParentNode) SetChild(n Node) { rule.child = n }
func (rule *ParentNode) Child() Node     { return rule.child }
func (rule *ParentNode) Repr() (parts []interface{}) {
	parts = rule.BaseNode.Repr()
	child := rule.Child()
	if child != nil {
		parts = append(parts, child.Repr())
	}
	return parts
}

type ListNode struct {
	BaseNode
	children []Node
}

func (rule *ListNode) initListNode()    { rule.children = make([]Node, 0) }
func (rule *ListNode) Length() int      { return len(rule.children) }
func (rule *ListNode) Children() []Node { return rule.children }
func (rule *ListNode) Append(n Node)    { rule.children = append(rule.children, n) }
func (rule *ListNode) Repr() (parts []interface{}) {
	parts = rule.BaseNode.Repr()
	children := rule.Children()
	for _, child := range children {
		parts = append(parts, child.Repr())
	}
	return parts
}
