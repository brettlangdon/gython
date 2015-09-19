package ast

type Node interface {
	Name() string
	Repr() []interface{}
}
