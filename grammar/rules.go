package grammar

import (
	"fmt"

	"github.com/brettlangdon/gython/symbol"
	"github.com/brettlangdon/gython/token"
)

type Rule interface {
	Name() string
	Repr() []interface{}
}

type TokenRule struct {
	Token *token.Token
}

func NewTokenRule(tok *token.Token) *TokenRule {
	return &TokenRule{
		Token: tok,
	}
}
func (rule *TokenRule) atomChild()                {}
func (rule *TokenRule) atomExpressionChild()      {}
func (rule *TokenRule) comparisonChild()          {}
func (rule *TokenRule) expressionStatementChild() {}
func (rule *TokenRule) factorChild()              {}
func (rule *TokenRule) fileInputChild()           {}
func (rule *TokenRule) shiftExpressionChild()     {}
func (rule *TokenRule) simpleStatementChild()     {}
func (rule *TokenRule) trailerChild()             {}
func (rule *TokenRule) Name() string              { return token.TokenNames[rule.Token.ID] }
func (rule *TokenRule) Repr() []interface{} {
	parts := make([]interface{}, 0)
	parts = append(parts, rule.Name())
	literal := fmt.Sprintf("%#v", rule.Token.Literal)
	return append(parts, literal)
}

type BaseRule struct {
	ID    symbol.SymbolID
	child Rule
}

func (rule *BaseRule) initBaseRule(id symbol.SymbolID) { rule.ID = id }
func (rule *BaseRule) Name() string                    { return symbol.SymbolNames[rule.ID] }
func (rule *BaseRule) Repr() (parts []interface{})     { return append(parts, rule.Name()) }

type ParentRule struct {
	BaseRule
	child Rule
}

func (rule *ParentRule) SetChild(n Rule) { rule.child = n }
func (rule *ParentRule) Child() Rule     { return rule.child }
func (rule *ParentRule) Repr() (parts []interface{}) {
	parts = rule.BaseRule.Repr()
	child := rule.Child()
	if child != nil {
		parts = append(parts, child.Repr())
	}
	return parts
}

type ListRule struct {
	BaseRule
	children []Rule
}

func (rule *ListRule) initListRule()    { rule.children = make([]Rule, 0) }
func (rule *ListRule) Length() int      { return len(rule.children) }
func (rule *ListRule) Children() []Rule { return rule.children }
func (rule *ListRule) Append(n Rule)    { rule.children = append(rule.children, n) }
func (rule *ListRule) Repr() (parts []interface{}) {
	parts = rule.BaseRule.Repr()
	children := rule.Children()
	for _, child := range children {
		parts = append(parts, child.Repr())
	}
	return parts
}
