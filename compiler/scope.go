package compiler

import "github.com/brettlangdon/gython/gython"

type Scope struct {
	Instructions []*Instruction

	Constants         *gython.Dict
	Names             *gython.Dict
	VariableNames     *gython.Dict
	FreeVariableNames *gython.Dict
	CellVariableNames *gython.Dict
}

func NewScope() *Scope {
	return &Scope{
		Instructions:      make([]*Instruction, 0),
		Constants:         gython.NewDict(),
		Names:             gython.NewDict(),
		VariableNames:     gython.NewDict(),
		FreeVariableNames: gython.NewDict(),
		CellVariableNames: gython.NewDict(),
	}
}

func (scope *Scope) AddInstruction(instr *Instruction) {
	scope.Instructions = append(scope.Instructions, instr)
}
