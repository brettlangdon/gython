package compiler

type Scope struct {
	Instructions []*Instruction
}

func NewScope() *Scope {
	return &Scope{
		Instructions: make([]*Instruction, 0),
	}
}

func (scope *Scope) AddInstruction(instr *Instruction) {
	scope.Instructions = append(scope.Instructions, instr)
}
