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

func (instruction *Instruction) Size() int {
	if instruction.Hasarg == false {
		// 1 byte for the opcode
		return 1
	} else if instruction.Oparg.Value > 0xffff {
		// 1 (opcode) + 1 (EXTENDED_ARG opcode) + 2 (oparg) + 2(oparg extended)
		return 6
	}
	// 1 (opcode) + 2 (oparg)
	return 3
}
