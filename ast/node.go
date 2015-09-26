package ast

type Node interface {
	node()
	String() string
}
