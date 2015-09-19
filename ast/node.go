package ast

type Node interface {
	ID() NodeID
	Name() string
}
