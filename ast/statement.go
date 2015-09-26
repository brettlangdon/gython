package ast

import (
	"fmt"
	"strings"
)

type Statement interface {
	Node
	stmt()
}

type Assign struct {
	Targets []Expression
	Value   Expression
}

func NewAssign(value Expression) *Assign {
	return &Assign{
		Targets: make([]Expression, 0),
		Value:   value,
	}
}

func (assign *Assign) node() {}
func (assign *Assign) stmt() {}
func (assign *Assign) Append(target Expression) {
	assign.Targets = append(assign.Targets, target)
}
func (assign *Assign) String() string {
	exprs := make([]string, 0)
	for _, expr := range assign.Targets {
		exprs = append(exprs, expr.String())
	}

	return fmt.Sprintf("Assign(targets=[%s], value=%s)", strings.Join(exprs, ", "), assign.Value.String())
}
