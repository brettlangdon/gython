package ast

import (
	"fmt"

	"github.com/brettlangdon/gython/gython"
)

type Expression interface {
	Node
	expr()
}

type Name struct {
	Identifier *gython.Unicode
	Context    ExpressionContext
}

func NewName(id string, ctx ExpressionContext) *Name {
	return &Name{
		Identifier: gython.NewUnicode([]byte(id)),
		Context:    ctx,
	}
}

func (name *Name) node() {}
func (name *Name) expr() {}
func (name *Name) String() string {
	return fmt.Sprintf("Name(id=%#v, ctx=%s)", name.Identifier, name.Context.String())
}

type Num struct {
	Value *gython.Float
}

func NewNum(i int64) *Num {
	return &Num{
		Value: gython.NewFloat(float64(i)),
	}
}

func (num *Num) node() {}
func (num *Num) expr() {}
func (num *Num) String() string {
	return fmt.Sprintf("Num(n=%d)", num.Value)
}
