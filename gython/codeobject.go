package gython

import "github.com/brettlangdon/gython/bytecode"

type CodeObject struct {
	ArgCount            int
	KeywordOnlyArgCount int
	LocalsCount         int
	StackSize           int
	Flags               int

	Code              *Bytes
	Constants         *Tuple
	Names             *Tuple
	VariableNames     *Tuple
	FreeVariableNames *Tuple
	CellVariableNames *Tuple

	Filename        *Unicode
	Name            *Unicode
	FirstLineNumber int
	LineNumberTable *Bytes
}

func NewCodeObject(filename []byte, name []byte, firstLineNumber int) *CodeObject {
	return &CodeObject{
		ArgCount:            0,
		KeywordOnlyArgCount: 0,
		LocalsCount:         0,
		StackSize:           0,
		Flags:               0,

		Code:              NewBytes(),
		Constants:         NewTuple(),
		Names:             NewTuple(),
		VariableNames:     NewTuple(),
		FreeVariableNames: NewTuple(),
		CellVariableNames: NewTuple(),

		Filename:        NewUnicode(filename),
		Name:            NewUnicode(name),
		FirstLineNumber: firstLineNumber,
		LineNumberTable: NewBytes(),
	}
}

func (codeobject *CodeObject) object() {}
func (codeobject *CodeObject) AppendOpcode(op bytecode.Opcode) {
	codeobject.Code.Append(byte(op))
}
func (codeobject *CodeObject) AppendInt(i int) {
	codeobject.Code.Append(byte(i))
}
