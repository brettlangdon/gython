package gython

type CodeObject struct {
	ArgCount            int64
	KeywordOnlyArgCount int64
	LocalsCount         int64
	StackSize           int64
	Flags               int64

	Code              *Bytes
	Constants         *Tuple
	Names             *Tuple
	VariableNames     *Tuple
	FreeVariableNames *Tuple
	CellVariableNames *Tuple

	Filename        *Unicode
	Name            *Unicode
	FirstLineNumber int64
	LineNumberTable *Bytes
}

func NewCodeObject(filename []byte, name []byte, firstLineNumber int64) *CodeObject {
	return &CodeObject{}
}

func (codeobject *CodeObject) object() {}
