package ast

import "fmt"

type Expression interface {
	Node
	expr()
}

type Name struct {
	Identifier string
	Context    ExpressionContext
}

func NewName(id string, ctx ExpressionContext) *Name {
	return &Name{
		Identifier: id,
		Context:    ctx,
	}
}

func (name *Name) node() {}
func (name *Name) expr() {}
func (name *Name) String() string {
	return fmt.Sprintf("Name(id=%#v, ctx=%s)", name.Identifier, name.Context.String())
}

type Num struct {
	Value int64
}

func NewNum(i int64) *Num {
	return &Num{
		Value: i,
	}
}

func (num *Num) node() {}
func (num *Num) expr() {}
func (num *Num) String() string {
	return fmt.Sprintf("Num(n=%d)", num.Value)
}
