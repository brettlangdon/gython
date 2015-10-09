package compiler

import (
	"github.com/brettlangdon/gython/bytecode"
	"github.com/brettlangdon/gython/gython"
)

type Instruction struct {
	Opcode bytecode.Opcode
	Oparg  *gython.Float
	Hasarg bool
	Line   int
}

func NewInstruction(opcode bytecode.Opcode, oparg *gython.Float, hasarg bool) *Instruction {
	return &Instruction{
		Opcode: opcode,
		Oparg:  oparg,
		Hasarg: hasarg,
		Line:   0,
	}
}
